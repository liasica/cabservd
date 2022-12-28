// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on cabservd by liasica, magicrolan@qq.com.

package service

import (
    "github.com/auroraride/adapter/model"
    "github.com/auroraride/cabservd/internal/app"
    "github.com/auroraride/cabservd/internal/core"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/cabinet"
    "github.com/auroraride/cabservd/internal/errs"
    "github.com/jinzhu/copier"
    "net/http"
)

type exchangeService struct {
    *BaseService
}

func NewExchange(params ...any) *exchangeService {
    return &exchangeService{
        BaseService: newService(params...),
    }
}

func (s *exchangeService) Usable(req *model.ExchangeRequest) (res *model.ExchangeUsableResponse) {
    res = &model.ExchangeUsableResponse{}
    creator := ent.Database.Scan.Create().SetSerial(req.Serial).SetUserID(s.User.ID).SetUserType(s.User.Type)
    defer func() {
        scan, _ := creator.Save(s.ctx)
        res.UUID = scan.ID.String()
    }()

    // 获取电柜状态
    cab, _ := NewCabinet().QueryCabinetWithBin(req.Serial)

    if cab == nil {
        app.Panic(http.StatusBadRequest, errs.CabinetNotFound)
    }

    if cab.Status != cabinet.StatusIdle {
        app.Panic(http.StatusBadRequest, errs.CabinetBusy)
    }

    if !cab.Online {
        app.Panic(http.StatusBadRequest, errs.CabinetOffline)
    }

    if len(cab.Edges.Bins) < 2 {
        app.Panic(http.StatusBadRequest, errs.CabinetNoEmpty)
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
            app.Panic(http.StatusBadRequest, errs.CabinetDoorOpened)
        }
        // 获取满电仓位
        if fully == nil || fully.Soc < item.Soc {
            // 若该仓位无电池, 继续循环
            if !item.IsStrictHasBattery(fakevoltage) {
                // TODO 该仓位是否出错
                continue
            }
            // 该仓位电量小于最小电量
            if item.Soc < req.MinSoc {
                continue
            }
            // 标定满仓
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
        app.Panic(http.StatusBadRequest, errs.CabinetNoFully)
    }

    // 如果无空仓
    if empty == nil {
        app.Panic(http.StatusBadRequest, errs.CabinetNoEmpty)
    }

    // 存储扫码记录
    creator.SetData(res)

    // 拷贝属性
    _ = copier.Copy(&res.Cabinet, cab)
    _ = copier.Copy(&res.Fully, fully)

    return
}
