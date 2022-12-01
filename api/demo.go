// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-11
// Based on cabservd by liasica, magicrolan@qq.com.

package api

import (
    "context"
    "fmt"
    "github.com/auroraride/cabservd/internal/core/kaixin"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/bin"
    "github.com/auroraride/cabservd/internal/errs"
    "github.com/gin-gonic/gin"
    log "github.com/sirupsen/logrus"
    "net/http"
    "sync"
    "time"
)

type demo struct {
}

var tasks sync.Map

type cabinetBin struct {
    Index int     `json:"-"`
    Name  string  `json:"name"`
    Soc   float64 `json:"soc"`
}

type taskStep struct {
    *cabinetBin
    status  uint8 // 0:进行中 1:成功 2:失败
    message string
}

type task struct {
    sn      string
    empty   *cabinetBin
    fully   *cabinetBin
    step    int
    steps   []*taskStep
    running bool
}

var Demo = new(demo)

func (*demo) Control(c *gin.Context) {
    var req struct {
        SN     string                `form:"sn" json:"sn"`
        Params []kaixin.ControlParam `json:"params"`
    }
    err := c.Bind(&req)
    if err == nil {
        err = kaixin.SendControl(req.SN, kaixin.ControlRequest{ParamList: req.Params})
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
        Order(ent.Asc(bin.FieldIndex)).
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

    // 是否有正在执行的任务
    if d.isBusy(req.SN) {
        c.JSON(http.StatusOK, gin.H{"error": errs.CabinetBusy.Error()})
        return
    }

    // 获取仓位状态
    var items ent.Bins
    items, err = ent.Database.Bin.Query().
        Where(bin.Serial(req.SN), bin.Enable(true), bin.Open(false), bin.Lock(false)).
        Order(ent.Desc(bin.FieldSoc)).
        All(context.Background())

    if len(items) == 0 {
        c.JSON(http.StatusOK, gin.H{"error": errs.CabinetNoFully.Error()})
        return
    }

    // 获取满电仓位
    max := items[0]
    if max.BatterySn == "" || max.Soc == 0 || max.Soc < 50 {
        c.JSON(http.StatusOK, gin.H{"error": errs.CabinetNoFully.Error()})
        return
    }
    fully := &cabinetBin{
        Name:  max.Name,
        Soc:   max.Soc,
        Index: max.Index,
    }

    // 获取空电仓位
    min := items[len(items)-1]
    if min.BatterySn != "" {
        c.JSON(http.StatusOK, gin.H{"error": errs.CabinetNoEmpty.Error()})
        return
    }
    empty := &cabinetBin{
        Name:  min.Name,
        Soc:   min.Soc,
        Index: min.Index,
    }

    // 开始执行换电任务
    t := &task{
        sn:    req.SN,
        empty: empty,
        fully: fully,
        step:  0,
        steps: []*taskStep{
            {cabinetBin: empty},
            {cabinetBin: empty},
            {cabinetBin: fully},
            {cabinetBin: fully},
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

    start := time.Now()
    for {
        if time.Now().Sub(start).Seconds() > 30 {
            return
        }
        res.Status = s.status
        res.Message = s.message
        if s.status != 0 {
            return
        }
    }
}

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
    err = t.doorOpen(t.empty.Index)
    if err != nil {
        return
    }
    t.steps[t.step].status = 1
    log.Infof("%s, 成功", t.steps[t.step].message)

    // 第二步, 识别仓门是否关闭
    t.step += 1
    t.steps[t.step].message = fmt.Sprintf("第②步, 监控电池放入空电仓[%s]并关闭", t.empty.Name)
    // 识别仓门是否关闭且电池是否放入
    err = t.doorOpenStatus(t.empty.Index, false, 1)
    if err != nil {
        return
    }
    t.steps[t.step].status = 1
    log.Infof("%s, 成功", t.steps[t.step].message)

    // 第三步, 开启满电仓门
    t.step += 1
    t.steps[t.step].message = fmt.Sprintf("第③步, 开启满电仓门[%s]", t.fully.Name)
    err = t.doorOpen(t.fully.Index)
    if err != nil {
        return
    }
    t.steps[t.step].status = 1
    log.Infof("%s, 成功", t.steps[t.step].message)

    // 第四步, 识别电池取出并关闭仓门
    t.step += 1
    t.steps[t.step].message = fmt.Sprintf("第④步, 监控电池取走[%s]并关闭", t.fully.Name)
    // 识别仓门是否关闭且电池是否放入
    err = t.doorOpenStatus(t.fully.Index, false, 2)
    if err != nil {
        return
    }
    t.steps[t.step].status = 1
    log.Infof("%s, 成功", t.steps[t.step].message)
}

func (t *task) doorOpen(index int) (err error) {
    params := []kaixin.ControlParam{
        {SignalData: kaixin.SignalData{
            ID:    kaixin.SignalCabinetControl,
            Value: kaixin.ControlOpenDoor,
        }, DoorID: fmt.Sprintf("%d", index+1)},
    }
    err = kaixin.SendControl(t.sn, kaixin.ControlRequest{ParamList: params})
    if err != nil {
        return
    }

    return t.doorOpenStatus(index, true, 0)
}

// 死循环查询仓门状态
// index: 检查的仓门index
// status: 待检查的状态
// battery: 是否检查电池放入状态 0不检查 1放入检查 2取出检查
func (t *task) doorOpenStatus(index int, status bool, battery uint) (err error) {
    var item *ent.Bin
    start := time.Now()
    var maxtime float64 = 120

    // time.Sleep(3 * time.Second)
    // return

    for {
        // TODO: 缓存
        item, err = ent.Database.Bin.Query().Where(bin.Serial(t.sn), bin.Index(index), bin.Enable(true)).First(context.Background())
        if err != nil {
            return
        }
        // 检查成功
        if item.Open == status {
            // 若仓门关闭并且检查电池
            switch battery {
            case 1:
                // 检查电池是否放入
                if item.BatterySn == "" {
                    maxtime = 30
                    start = time.Now()

                    // TODO: 重复弹开
                    if time.Now().Sub(start).Seconds() > maxtime {
                        err = errs.ExchangeBatteryLost
                    }
                }
                // case 2:
                //     // 检查电池是否取出
                //     // TODO: 是否取走
                //     if item.BatterySn != "" {
                //         err = errs.ExchangeBatteryLost
                //     }
            }
            return
        }
        // 超时
        if time.Now().Sub(start).Seconds() > maxtime {
            err = errs.ExchangeTimeOut
            return
        }
        // 10ms查询一次
        time.Sleep(10 * time.Millisecond)
    }
}
