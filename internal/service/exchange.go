// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on cabservd by liasica, magicrolan@qq.com.

package service

import (
    "github.com/auroraride/adapter"
    "github.com/auroraride/adapter/defs/cabdef"
    "github.com/auroraride/cabservd/internal/app"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/cabinet"
    "github.com/auroraride/cabservd/internal/ent/scan"
    "github.com/auroraride/cabservd/internal/notice"
    "github.com/auroraride/cabservd/internal/types"
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

// Usable 获取电柜待换电信息
func (s *exchangeService) Usable(req *cabdef.ExchangeUsableRequest) (res *cabdef.CabinetBinUsableResponse) {
    res = &cabdef.CabinetBinUsableResponse{
        Cabinet: new(cabdef.Cabinet),
        Fully:   new(cabdef.Bin),
        Empty:   new(cabdef.Bin),
    }

    cs := NewCabinet(s.User)

    // 获取电柜状态
    cab, _ := cs.QuerySerialWithBin(req.Serial)

    // 检查电柜是否可换电
    err := cs.DetectCabinet(cab)
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

    // 获取空仓和满电仓位
    var fully, empty *ent.Bin
    fully, empty, err = cs.BusinessInfo(req.Model, cab, req.Minsoc, 1, 1)

    if err != nil {
        app.Panic(http.StatusBadRequest, err)
    }

    // // 如果无满电
    // if fully == nil {
    //     app.Panic(http.StatusBadRequest, adapter.ErrorCabinetNoFully)
    // }
    //
    // // 如果无空仓
    // if empty == nil {
    //     app.Panic(http.StatusBadRequest, adapter.ErrorCabinetNoEmpty)
    // }

    // 拷贝属性
    _ = copier.Copy(res.Cabinet, cab)
    _ = copier.Copy(res.Fully, fully)
    _ = copier.Copy(res.Empty, empty)

    // 存储扫码记录
    sm := NewScan(s.User).Create(adapter.BusinessExchange, req.Serial, cab, res)
    res.UUID = sm.UUID.String()

    return
}

func (s *exchangeService) Do(req *cabdef.ExchangeRequest) (res *cabdef.ExchangeResponse) {
    // 查询扫码记录
    sc := NewScan(s.User).CensorX(req.UUID, req.Timeout, req.Minsoc)

    // 开始同步换电流程
    results, err := s.start(req, sc)

    res = &cabdef.ExchangeResponse{
        Results: results,
    }

    for _, result := range results {
        res.Success = result.Step == 4 && result.Success

        // 取出的电池 (第三步成功即视为取走电池)
        if result.Step == 3 && result.Success {
            res.PutoutBattery = sc.Data.Fully.BatterySn
        }

        // 放入的电池, 第二步成功且有电池视为放入电池
        if result.Step <= 2 && result.BatterySN != "" {
            res.PutinBattery = result.BatterySN
        }
    }

    if err != nil {
        res.Error = err.Error()
    }

    return
}

func (s *exchangeService) start(req *cabdef.ExchangeRequest, sc *ent.Scan) (res []*cabdef.ExchangeStepMessage, err error) {
    cab, _ := NewCabinet(s.User).QueryWithBin(sc.CabinetID)

    // 检查电柜是否可换电
    err = NewCabinet(s.User).DetectCabinet(cab)
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

    bins := []*cabdef.Bin{
        sc.Data.Empty,
        sc.Data.Fully,
    }

    cb := func(r *cabdef.BusinessStepResult) {
        data := silk.Pointer(cabdef.ExchangeStepMessage(*r))
        res = append(res, data)
        // 异步发送结果
        go notice.Aurservd.SendMessage(data)
    }

    for i, conf := range types.ExchangeConfigure {
        b := bins[i]

        err = NewBin(s.User).Operate(&types.Bin{
            Timeout:      req.Timeout,
            Serial:       sc.Serial,
            UUID:         req.UUID,
            Ordinal:      b.Ordinal,
            Business:     adapter.BusinessExchange,
            Steps:        conf,
            Battery:      req.Battery,
            StepCallback: cb,
        })

        if err != nil {
            return
        }

    }

    return
}
