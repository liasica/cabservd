// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on cabservd by liasica, magicrolan@qq.com.

package service

import (
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/console"
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

// func (s *consoleService) Create(step int, typ console.Type, req *adapter.OperateRequest, cabinetID uint64, b *ent.Bin) (ec *ent.Console, err error) {
//     // 查询最新仓位信息
//     var before *adapter.BinInfo
//     var bid *uint64
//     // if req.Ordinal != nil {
//     //     b, _ = NewBin(s.User).QuerySerialOrdinal(req.Serial, *req.Ordinal)
//     //     if b == nil {
//     //         err = adapter.ErrorCabinetBinNotFound
//     //         return
//     //     }
//     //     bid = silk.Pointer(b.ID)
//     //     before = b.Info()
//     // }
//     if b != nil {
//         bid = silk.Pointer(b.ID)
//         before = b.Info()
//     }
//
//     creator := s.orm.Create().
//         SetOperate(req.Operate).
//         SetCabinetID(cabinetID).
//         SetNillableBinID(bid).
//         SetSerial(req.Serial).
//         SetUserID(s.User.ID).
//         SetUserType(s.User.Type).
//         SetStatus(console.StatusRunning).
//         SetStartAt(time.Now()).
//         SetBeforeBin(before).
//         SetStep(step).
//         SetType(typ)
//
//     if req.UUID == nil {
//         creator.SetUUID(uuid.New())
//     } else {
//         creator.SetUUID(*req.UUID)
//     }
//
//     ec, err = creator.Save(s.ctx)
//
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
