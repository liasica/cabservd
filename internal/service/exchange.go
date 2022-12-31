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
    "github.com/liasica/go-helpers/silk"
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
        return adapter.ErrorCabinetNotFound
    }

    if cab.Status != cabinet.StatusIdle {
        return adapter.ErrorCabinetBusy
    }

    if !cab.Online {
        return adapter.ErrorCabinetOffline
    }

    if len(cab.Edges.Bins) < 2 {
        return adapter.ErrorCabinetNoEmpty
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
        app.Panic(http.StatusBadRequest, adapter.ErrorCabinetBusy)
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
            app.Panic(http.StatusBadRequest, adapter.ErrorCabinetDoorOpened)
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
        app.Panic(http.StatusBadRequest, adapter.ErrorCabinetNoFully)
    }

    // 如果无空仓
    if empty == nil {
        app.Panic(http.StatusBadRequest, adapter.ErrorCabinetNoEmpty)
    }

    // 拷贝属性
    _ = copier.Copy(res.Cabinet, cab)
    _ = copier.Copy(res.Fully, fully)
    _ = copier.Copy(res.Empty, empty)

    // 存储扫码记录
    sm := NewScan(s.User).Create(req.Serial, cab, res)
    res.UUID = sm.UUID.String()

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
        res.Success = result.Step == adapter.ExchangeStepFourth && result.Success

        // 取出的电池
        if result.Step == adapter.ExchangeStepThird && result.Success {
            res.PutoutBattery = sc.Data.Fully.BatterySn
        }

        // 放入的电池
        if result.Step <= adapter.ExchangeStepSecond && result.After.BatterySN != "" {
            res.PutinBattery = result.After.BatterySN
        }
    }

    if err != nil {
        res.Error = err.Error()
    }

    return
}

func (s *exchangeService) start(req *adapter.ExchangeRequest, sc *ent.Scan) (results []*adapter.ExchangeStepMessage, err error) {
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
        var result *adapter.ExchangeStepMessage
        result, err = s.step(req, sc, conf)
        if err != nil {
            return
        }
        results = append(results, result)
    }

    return
}

func (s *exchangeService) step(req *adapter.ExchangeRequest, sc *ent.Scan, conf adapter.ExchangeStepConfigure) (result *adapter.ExchangeStepMessage, err error) {
    var ordinal int

    switch conf.Step {
    case adapter.ExchangeStepFirst, adapter.ExchangeStepSecond:
        ordinal = sc.Data.Empty.Ordinal
    case adapter.ExchangeStepThird, adapter.ExchangeStepFourth:
        ordinal = sc.Data.Fully.Ordinal
    }

    var ec *ent.Console
    ec, err = NewBin(s.User).Operate(&adapter.OperateRequest{
        UUID:               sc.UUID,
        Serial:             sc.Serial,
        Ordinal:            silk.Pointer(ordinal),
        Operate:            conf.Operate,
        Step:               silk.Pointer(conf.Step),
        Timeout:            req.TimeOut,
        VerifyPutinBattery: req.Battery,
    })

    if err != nil {
        return
    }

    result = ec.StepResult()

    // 异步发送结果
    go notice.Aurservd.SendData(result)

    return
}
