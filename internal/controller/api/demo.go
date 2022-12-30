// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-11
// Based on cabservd by liasica, magicrolan@qq.com.

package api

import (
    "context"
    "fmt"
    errs "github.com/auroraride/adapter/errors"
    "github.com/auroraride/adapter/model"
    "github.com/auroraride/cabservd/internal/core"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/bin"
    "github.com/auroraride/cabservd/internal/ent/cabinet"
    "github.com/gin-gonic/gin"
    log "github.com/sirupsen/logrus"
    "net/http"
    "sync"
    "time"
)

type demo struct {
}

var tasks sync.Map

type taskStep struct {
    *ent.Bin
    status  uint8 // 0:进行中 1:成功 2:失败
    message string
}

type task struct {
    serial  string
    empty   *ent.Bin
    fully   *ent.Bin
    step    int
    steps   []*taskStep
    running bool
}

var Demo = new(demo)

func (*demo) Control(c *gin.Context) {
    var req struct {
        Serial  string         `form:"serial" json:"serial"`
        Type    model.Operator `json:"type" form:"type"`
        Ordinal int            `json:"ordinal" form:"ordinal"`
    }
    err := c.Bind(&req)
    if err == nil {
        err = core.Hub.Bean.SendControl(req.Serial, req.Type, req.Ordinal)
    }

    c.JSON(http.StatusOK, gin.H{
        "error": fmt.Sprintf("%v", err),
    })
}

func (*demo) Exchange(c *gin.Context) {
    var req struct {
        SN string `form:"sn" json:"sn"`
    }

    err := c.Bind(&req)
    if err != nil {
        c.JSON(http.StatusOK, gin.H{
            "error": err,
        })
        return
    }

    // 获取仓位状态
    var items ent.Bins
    items, err = ent.Database.Bin.Query().
        Where(bin.Serial(req.SN)).
        Order(ent.Asc(bin.FieldOrdinal)).
        All(context.Background())

    c.HTML(http.StatusOK, "demo/exchange.go.html", gin.H{
        "items": items,
        "sn":    req.SN,
    })
}

func (*demo) isBusy(sn string) (busy bool) {
    v, ok := tasks.Load(sn)
    if !ok {
        return
    }
    if t, isTask := v.(*task); isTask {
        return t.running
    }
    return
}

func (d *demo) Start(c *gin.Context) {
    var req struct {
        SN string `form:"sn" json:"sn"`
    }
    err := c.Bind(&req)
    if err != nil {
        c.JSON(http.StatusOK, gin.H{"error": errs.ParamValidateFailed.Error()})
        return
    }

    // TODO 判断电柜状态
    // 查询电柜信息
    cab, _ := ent.Database.Cabinet.Query().Where(cabinet.Serial(req.SN)).First(context.Background())
    if cab == nil {
        c.JSON(http.StatusOK, gin.H{"error": errs.CabinetNotFound.Error()})
        return
    }
    if cab.Status == cabinet.StatusInitializing {
        c.JSON(http.StatusOK, gin.H{"error": errs.CabinetInitializing.Error()})
        return
    }
    if cab.Status == cabinet.StatusAbnormal {
        c.JSON(http.StatusOK, gin.H{"error": errs.CabinetAbnormal.Error()})
        return
    }
    if !cab.Online {
        c.JSON(http.StatusOK, gin.H{"error": errs.CabinetOffline.Error()})
        return
    }

    // 是否有正在执行的任务
    if d.isBusy(req.SN) {
        c.JSON(http.StatusOK, gin.H{"error": errs.CabinetBusy.Error()})
        return
    }

    // 获取仓位状态
    var items ent.Bins
    items, err = ent.Database.Bin.Query().
        Where(bin.Serial(req.SN), bin.Enable(true), bin.Open(false), bin.Health(true)).
        Order(ent.Desc(bin.FieldSoc)).
        All(context.Background())

    var (
        fully  *ent.Bin
        empty  *ent.Bin
        minsoc float64 = 1
    )

    // 获取仓位
    fakevoltage, fakecurrent := core.Hub.Bean.GetEmptyDeviation()
    for _, item := range items {
        // 获取满电仓位
        if fully == nil {
            // 若该仓位无电池, 继续循环
            if !item.IsStrictHasBattery(fakevoltage) {
                // TODO 该仓位是否出错
                continue
            }
            // 该仓位电量小于最小电量
            if item.Soc < minsoc {
                continue
            }
            fully = item
        }
        if empty == nil {
            // 若该仓位无电池, 标记为空仓
            if !item.IsLooseHasBattery(fakevoltage, fakecurrent) {
                empty = item
            }
        }
    }

    // 如果无满电
    if fully == nil {
        c.JSON(http.StatusOK, gin.H{"error": errs.CabinetNoFully.Error()})
        return
    }

    // 如果无空仓
    if empty == nil {
        c.JSON(http.StatusOK, gin.H{"error": errs.CabinetNoEmpty.Error()})
        return
    }

    // 开始执行换电任务
    t := &task{
        serial: req.SN,
        empty:  empty,
        fully:  fully,
        step:   0,
        steps: []*taskStep{
            {Bin: empty},
            {Bin: empty},
            {Bin: fully},
            {Bin: fully},
        },
        running: true,
    }

    tasks.Store(req.SN, t)

    go t.run()

    // 返回
    c.JSON(http.StatusOK, gin.H{
        "fully": fully,
        "empty": empty,
    })
}

func (*demo) Status(c *gin.Context) {
    var res struct {
        Message string `json:"message"`
        Step    int    `json:"step"`
        Status  uint8  `json:"status"` // 0:进行中 1:成功 2:失败
    }

    var err error

    defer func() {
        if err != nil {
            c.JSON(http.StatusOK, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, res)
    }()

    var req struct {
        Step int    `json:"step" form:"step"`
        SN   string `json:"sn" form:"sn"`
    }

    err = c.Bind(&req)
    if err != nil {
        return
    }

    res.Step = req.Step
    t, ok := tasks.Load(req.SN)
    if !ok {
        err = errs.ExchangeTaskNotExist
        return
    }
    s := t.(*task).steps[req.Step]

    startAt := time.Now()

    for {
        if time.Now().Sub(startAt).Seconds() > 30 {
            return
        }
        res.Status = s.status
        res.Message = s.message
        if s.status != 0 {
            return
        }
    }
}

// TODO 取走电池未关门超时逻辑
func (t *task) run() {
    var (
        err error
    )

    defer func() {
        t.running = false
        msg := t.steps[t.step].message
        if err != nil {
            t.steps[t.step].message = fmt.Sprintf("%s, 失败: %v", msg, err)
            t.steps[t.step].status = 2
            return
        }
    }()

    // 第一步, 开启空电仓门并检查仓门是否开启
    t.steps[t.step].message = fmt.Sprintf("第①步, 开启空电仓门[%s]", t.empty.Name)
    err = t.doorOpen(t.empty)
    if err != nil {
        return
    }
    t.steps[t.step].status = 1
    log.Infof("%s, 成功", t.steps[t.step].message)

    // 第二步, 识别仓门是否关闭
    t.step += 1
    t.steps[t.step].message = fmt.Sprintf("第②步, 监控电池放入空电仓[%s]并关闭", t.empty.Name)
    // 识别仓门是否关闭且电池是否放入
    err = t.doorOpenStatus(t.empty, false, 1)
    if err != nil {
        return
    }
    t.steps[t.step].status = 1
    log.Infof("%s, 成功", t.steps[t.step].message)

    // 第三步, 开启满电仓门
    t.step += 1
    t.steps[t.step].message = fmt.Sprintf("第③步, 开启满电仓门[%s]", t.fully.Name)
    err = t.doorOpen(t.fully)
    if err != nil {
        return
    }
    t.steps[t.step].status = 1
    log.Infof("%s, 成功", t.steps[t.step].message)

    // 第四步, 识别电池取出并关闭仓门
    t.step += 1
    t.steps[t.step].message = fmt.Sprintf("第④步, 监控电池取走[%s]并关闭", t.fully.Name)
    // 识别仓门是否关闭且电池是否取走
    err = t.doorOpenStatus(t.fully, false, 2)
    if err != nil {
        return
    }
    t.steps[t.step].status = 1
    log.Infof("%s, 成功", t.steps[t.step].message)
}

func (t *task) doorOpen(target *ent.Bin) (err error) {
    err = core.Hub.Bean.SendControl(t.serial, model.OperatorBinOpen, target.Ordinal)
    if err != nil {
        return
    }

    return t.doorOpenStatus(target, true, 0)
}

// 死循环查询仓门状态
// target: 检查的仓位
// status: 待检查的状态 true:开门 false:关门
// battery: 是否检查电池放入状态 0不检查 1放入检查 2取出检查
func (t *task) doorOpenStatus(target *ent.Bin, status bool, battery uint) (err error) {
    var (
        // 仓位信息
        item *ent.Bin

        // 步骤最长时间
        maxtime float64 = 120 // TODO 测试使用120s超时

        // 开始时间
        startAt = time.Now()

        // 检测电池是否放入或取出时间
        batteryCheckMaxtime float64 = 15

        // 仓门状态匹配时间
        statusTime time.Time
    )

    fakevoltage, _ := core.Hub.Bean.GetEmptyDeviation()

    for {
        // 10ms查询一次
        time.Sleep(10 * time.Millisecond)

        // 超时
        if time.Now().Sub(startAt).Seconds() > maxtime {
            err = errs.ExchangeTimeOut
            return
        }

        // TODO: 使用缓存
        item, err = ent.Database.Bin.Query().Where(bin.UUID(target.UUID)).First(context.Background())
        if err != nil {
            return
        }

        // 检查成功
        if item.Open == status {
            if statusTime.IsZero() {
                statusTime = time.Now()
            }

            // 若仓门关闭并且检查电池
            switch battery {
            case 1:
                // 检查电池是否放入
                if item.IsStrictHasBattery(fakevoltage) {
                    // 检测到电池, 返回成功
                    return
                }

                // TODO: 重复弹开
                // 未检测到电池, 继续轮询
                // 超时
                if time.Now().Sub(statusTime).Seconds() > batteryCheckMaxtime {
                    err = errs.ExchangeBatteryLost
                    // 返回错误
                    return
                }
                continue
            case 2:
                // 检查电池是否取出
                // TODO: 是否取走, 重复弹开
                if !item.IsStrictHasBattery(fakevoltage) {
                    return errs.ExchangeBatteryExist
                }
            }
            return
        }
    }
}
