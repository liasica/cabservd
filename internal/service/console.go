// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on cabservd by liasica, magicrolan@qq.com.

package service

import (
    "github.com/auroraride/adapter"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/console"
    "github.com/google/uuid"
    "time"
)

type consoleService struct {
    *BaseService

    orm *ent.ConsoleClient
}

func NewConsole(params ...any) *consoleService {
    return &consoleService{
        BaseService: newService(params...),
        orm:         ent.Database.Console,
    }
}

// // StartExchangeStep 开始换电步骤
// func (s *consoleService) StartExchangeStep(sc *ent.Scan, conf adapter.ExchangeStepConfigure) (ec *ent.Console, b *ent.Bin, err error) {
//     var ordinal int
//
//     switch conf.Step {
//     case adapter.ExchangeStepFirst, adapter.ExchangeStepSecond:
//         ordinal = sc.Data.Empty.Ordinal
//     case adapter.ExchangeStepThird, adapter.ExchangeStepFourth:
//         ordinal = sc.Data.Fully.Ordinal
//     }
//
//     return s.Create(&adapter.OperateRequest{
//         UUID:    sc.UUID,
//         Serial:  sc.Serial,
//         Ordinal: ordinal,
//         Operate: conf.Operate,
//         User:    s.User,
//         Step:    silk.Pointer(conf.Step),
//     })
//
//     // // 查询仓位信息
//     // b, err = NewBin(s.User).Query(bid)
//     // if err != nil {
//     //     return
//     // }
//     //
//     // if conf.Door == adapter.DetectDoorOpen {
//     //     op = adapter.OperateBinOpen
//     // }
//     //
//     // if conf.Battery != adapter.DetectBatteryIgnore {
//     //     op = adapter.OperateBinDetect
//     // }
//     //
//     // ec, err = s.orm.Create().
//     //     SetOperate(op).
//     //     SetCabinetID(sc.CabinetID).
//     //     SetSerial(sc.Serial).
//     //     SetBinID(b.ID).
//     //     SetUUID(sc.UUID).
//     //     SetType(console.TypeExchange).
//     //     SetUserID(sc.UserID).
//     //     SetUserType(sc.UserType).
//     //     SetStep(conf.Step).
//     //     SetStatus(console.StatusRunning).
//     //     SetBeforeBin(b.Info()).
//     //     SetStartAt(time.Now()).
//     //     Save(s.ctx)
//     // return
// }

func (s *consoleService) Create(req *adapter.OperateRequest) (ec *ent.Console, b *ent.Bin, err error) {
    // 查询最新仓位信息
    b, _ = NewBin(s.User).QuerySerialOrdinal(req.Serial, *req.Ordinal)
    if b == nil {
        err = adapter.ErrorCabinetBinNotFound
        return
    }

    ec, err = s.orm.Create().
        SetOperate(req.Operate).
        SetCabinetID(b.CabinetID).
        SetBinID(b.ID).
        SetSerial(b.Serial).
        SetType(console.TypeOperate).
        SetUUID(uuid.New()).
        SetUserID(s.User.ID).
        SetUserType(s.User.Type).
        SetStatus(console.StatusRunning).
        SetStartAt(time.Now()).
        SetBeforeBin(b.Info()).
        SetNillableStep(req.Step).
        Save(s.ctx)

    return
}

// func (s *consoleService) Operate(req *adapter.OperateRequest) (ec *ent.Console, b *ent.Bin, err error) {
//     if req.Ordinal == nil {
//         app.Panic(adapter.ErrorCabinetBinOrdinalRequired)
//     }
//
//     // 查询仓位信息
//     b, err = NewBin(s.User).QuerySerialOrdinal(req.Serial, *req.Ordinal)
//     if err != nil {
//         err = adapter.ErrorCabinetBinNotFound
//         return
//     }
//
//     ec, err = s.orm.Create().
//         SetOperate(req.Type).
//         SetCabinetID(b.CabinetID).
//         SetBinID(b.ID).
//         SetSerial(b.Serial).
//         SetType(console.TypeOperate).
//         SetUUID(uuid.New()).
//         SetUserID(s.User.ID).
//         SetUserType(s.User.Type).
//         SetStatus(console.StatusRunning).
//         SetStartAt(time.Now()).
//         SetBeforeBin(b.Info()).Save(s.ctx)
//     return
// }

// Update 更新记录
func (s *consoleService) Update(ec *ent.Console, b *ent.Bin, err error) *ent.Console {
    now := time.Now()
    cr := ec.Update().SetStopAt(now)
    if ec.StartAt != nil {
        cr.SetDuration(now.Sub(*ec.StartAt).Seconds())
    }

    // 仓位信息
    if b != nil {
        cr.SetAfterBin(b.Info())
    }

    if err != nil {
        cr.SetStatus(console.StatusFailed).SetMessage(err.Error())
    } else {
        cr.SetStatus(console.StatusSuccess)
    }

    ec, _ = cr.Save(s.ctx)
    return ec
}
