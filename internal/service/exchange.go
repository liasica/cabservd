// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on cabservd by liasica, magicrolan@qq.com.

package service

import (
    "github.com/auroraride/adapter"
    "github.com/auroraride/cabservd/internal/app"
    "github.com/auroraride/cabservd/internal/core"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/cabinet"
    "github.com/auroraride/cabservd/internal/ent/scan"
    "github.com/auroraride/cabservd/internal/notice"
    "github.com/jinzhu/copier"
    "net/http"
    "time"
)

type exchangeService struct {
    *BaseService
}

func NewExchange(params ...any) *exchangeService {
    return &exchangeService{
        BaseService: newService(params...),
    }
}

func (s *exchangeService) DetectCabinet(cab *ent.Cabinet) error {
    if cab == nil {
        return adapter.CabinetNotFound
    }

    if cab.Status != cabinet.StatusIdle {
        return adapter.CabinetBusy
    }

    if !cab.Online {
        return adapter.CabinetOffline
    }

    if len(cab.Edges.Bins) < 2 {
        return adapter.CabinetNoEmpty
    }

    return nil
}

// Usable 获取电柜待换电信息
func (s *exchangeService) Usable(req *adapter.ExchangeUsableRequest) (res *adapter.ExchangeUsableResponse) {
    res = &adapter.ExchangeUsableResponse{
        Cabinet: new(adapter.Cabinet),
        Fully:   new(adapter.Bin),
        Empty:   new(adapter.Bin),
    }

    // defer func() {
    //     b, _ := json.Marshal(res)
    //     log.Infof("[SCAN] %s", b)
    // }()

    // 获取电柜状态
    cab, _ := NewCabinet(s.User).QuerySerialWithBin(req.Serial)

    // 检查电柜是否可换电
    err := s.DetectCabinet(cab)
    if err != nil {
        app.Panic(http.StatusBadRequest, err)
    }

    // 查询限定时间内其他扫码用户
    exists, _ := ent.Database.Scan.Query().Where(
        scan.CabinetID(cab.ID),
        scan.UserIDNEQ(s.User.ID),
        scan.CreatedAtGT(time.Now().Add(-time.Duration(req.Lock)*time.Second)),
    ).Exist(s.ctx)
    if exists {
        app.Panic(http.StatusBadRequest, adapter.CabinetBusy)
    }

    // TODO 查询是否有正在执行的任务?

    // 获取仓位
    var (
        fully, empty *ent.Bin
    )

    fakevoltage, fakecurrent := core.Hub.Bean.GetEmptyDeviation()
    for _, item := range cab.Edges.Bins {
        // 如果仓位未启用或仓位不健康直接跳过
        if !item.Enable || !item.Health {
            continue
        } else if item.Open {
            // 有正常未关闭的仓门直接报错
            app.Panic(http.StatusBadRequest, adapter.CabinetDoorOpened)
        }
        // 宽松判定是否有电池
        if item.IsLooseHasBattery(fakevoltage, fakecurrent) {
            // 若有电池
            // 获取满电仓位
            if fully == nil || fully.Soc < item.Soc {
                // 该仓位电量小于最小电量
                if item.Soc < req.Minsoc {
                    continue
                }
                // 标定满仓
                fully = item
            }
        } else {
            // 若无电池
            if empty == nil {
                empty = item
            }
        }
    }

    // 如果无满电
    if fully == nil {
        app.Panic(http.StatusBadRequest, adapter.CabinetNoFully)
    }

    // 如果无空仓
    if empty == nil {
        app.Panic(http.StatusBadRequest, adapter.CabinetNoEmpty)
    }

    // 拷贝属性
    _ = copier.Copy(res.Cabinet, cab)
    _ = copier.Copy(res.Fully, fully)
    _ = copier.Copy(res.Empty, empty)

    // 存储扫码记录
    sm := NewScan(s.User).Create(req.Serial, cab, res)
    res.UUID = sm.ID.String()

    return
}

func (s *exchangeService) Do(req *adapter.ExchangeRequest) (res *adapter.ExchangeResponse) {
    // 查询扫码记录
    sc := NewScan(s.User).CensorX(req)

    // 开始同步换电流程
    results, err := s.start(req, sc)

    res = &adapter.ExchangeResponse{
        Results: results,
    }

    for _, result := range results {
        if result.Step == adapter.ExchangeStepFourth && result.Success {
            res.Success = true
            res.AfterBattery = sc.Data.Fully.BatterySn
        }
        if result.Step == adapter.ExchangeStepSecond && result.Success {
            res.BeforeBattery = result.After.BatterySN
        }
    }

    if err != nil {
        res.Error = err.Error()
    }

    return
}

func (s *exchangeService) start(req *adapter.ExchangeRequest, sc *ent.Scan) (results []*adapter.ExchangeStepResult, err error) {
    cab, _ := NewCabinet(s.User).QueryWithBin(sc.CabinetID)

    // 检查电柜是否可换电
    err = s.DetectCabinet(cab)
    if err != nil {
        return
    }

    // 标记电柜为换电中
    _ = cab.Update().SetStatus(cabinet.StatusExchange).Exec(s.ctx)

    defer func() {
        // 标记电柜为空闲
        _ = cab.Update().SetStatus(cabinet.StatusIdle).Exec(s.ctx)

        // 标记扫码失效
        _ = sc.Update().SetEfficient(false).Exec(s.ctx)

        // TODO 任务标记???
    }()

    for _, conf := range adapter.ExchangeStepConfigures {
        var result *adapter.ExchangeStepResult
        result, err = s.step(req.TimeOut, sc, conf)
        if err != nil {
            return
        }
        results = append(results, result)
    }

    return
}

func (s *exchangeService) step(sec int64, sc *ent.Scan, conf adapter.ExchangeStepConfigure) (result *adapter.ExchangeStepResult, err error) {
    var (
        ec   *ent.Console
        eb   *ent.Bin
        bs   = NewBin(s.User)
        cs   = NewConsole(s.User)
        open = conf.Door == adapter.DetectDoorOpen
    )

    // 创建换电步骤
    ec, eb, err = cs.StartExchangeStep(sc, conf.Step, open)
    if err != nil {
        return
    }

    ch := make(chan notice.IDSerialGetter)
    notice.Postgres.SetListener(notice.PostgresChannelBin, eb.ID, ch)

    defer func() {
        // 删除监听
        notice.Postgres.DeleteListener(ch)

        // 更新记录
        ec, eb = cs.Update(ec, err)

        // 获取result
        result = ec.StepResult()

        // 发送result
        notice.Aurservd.SendData(&adapter.Data[adapter.ExchangeStepResult]{
            Type:  adapter.DataTypeExchangeStep,
            Value: result,
        })
    }()

    // 如果需要开仓
    if open {
        // 开仓
        err = bs.OpenDoor(sc.Serial, eb.Ordinal)
        if err != nil {
            return
        }
    }

    // TODO: 开仓失败后是否重复弹开逻辑???

    var batteryOk, doorOk adapter.Bool

    // 定义超时时间
    timeout := time.After(time.Duration(sec) * time.Second)

    for {
        select {
        case x := <-ch:
            b := x.(*ent.Bin)

            // 检查仓位是否满足条件
            doorOk, err = bs.DetectDoor(b, conf.Door)
            if err != nil {
                return
            }

            // 检查电池是否满足条件
            batteryOk, err = bs.DetectBattery(b, conf.Battery)
            if err != nil {
                return
            }

            if doorOk && batteryOk {
                // 更新仓位信息
                *eb = *b
                return
            }

        case <-timeout:
            // 超时
            err = adapter.ExchangeTimeOut
            return
        }
    }
}
