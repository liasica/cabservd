// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-03
// Based on cabservd by liasica, magicrolan@qq.com.

package service

import (
    "fmt"
    errs "github.com/auroraride/adapter/errors"
    "github.com/auroraride/adapter/model"
    "github.com/auroraride/cabservd/internal/core"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/bin"
    "github.com/auroraride/cabservd/internal/types"
    log "github.com/sirupsen/logrus"
)

type binService struct {
    // ordinal int
    // bin     *ent.Bin
    // cabinet *ent.Cabinet
    *BaseService
    orm *ent.BinClient
}

func NewBin(params ...any) *binService {
    return &binService{
        BaseService: newService(params...),
        orm:         ent.Database.Bin,
    }
}

// func NewBin(cab *ent.Cabinet, ordinal int) *binService {
//     s := &binService{
//         cabinet: cab,
//         ordinal: ordinal,
//         ctx:     context.WithValue(context.Background(), "cabinet", cab),
//     }
//
//     ctx := context.Background()
//     cb, _ := ent.Database.Bin.Query().Where(bin.Serial(cab.Serial), bin.Brand(cab.Brand), bin.Ordinal(ordinal)).First(ctx)
//     ctx = context.WithValue(ctx, "bin", cb)
//
//     s.ctx = ctx
//     s.bin = cb
//     return s
// }
//
// // Enable 控制仓位启用/禁用
// func (s *binService) Enable(enable bool) error {
//     var t types.ControlType
//
//     switch enable {
//     case true:
//         t = types.ControlTypeBinEnable
//     case false:
//         t = types.ControlTypeBinDisable
//     default:
//         return errs.CabinetControlParamError
//     }
//
//     return core.Hub.Control(&types.ControlRequest{
//         Type:    t,
//         Serial:  s.bin.Serial,
//         Ordinal: silk.Int(s.bin.Ordinal),
//     })
// }
//
// // Open 打开仓门
// func (s *binService) Open() error {
//     return core.Hub.Control(&types.ControlRequest{
//         Type:    types.ControlTypeBinOpen,
//         Serial:  s.bin.Serial,
//         Ordinal: silk.Int(s.bin.Ordinal),
//     })
// }

func (s *binService) QueryAllBin() ent.Bins {
    items, _ := s.orm.Query().All(s.ctx)
    return items
}

func (s *binService) Query(id uint64) (*ent.Bin, error) {
    return s.orm.Query().Where(bin.ID(id)).First(s.ctx)
}

func (s *binService) OpenDoor(serial string, ordinal int) (err error) {
    err = core.Hub.Bean.SendControl(serial, types.ControlTypeBinOpen, ordinal)
    message := "成功"
    if err != nil {
        message = fmt.Sprintf("失败, %v", err)
    }
    log.Infof("[BIN] [%s - %d, %s] 开门请求: %s", serial, ordinal, s.User, message)

    return
}

// DetectUsable 检查仓位是否可用
func (s *binService) DetectUsable(b *ent.Bin) (ok model.Bool, message string) {
    ok = model.Bool(b.IsUsable())
    if ok {
        message = fmt.Sprintf("仓位可用")
    } else {
        message = fmt.Sprintf("仓位不可用(health: %v, enable: %v)", b.Health, b.Enable)
    }
    return
}

// DetectDoor 识别仓门开关状态
func (s *binService) DetectDoor(b *ent.Bin, d model.DetectDoor) (ok model.Bool, err error) {
    switch d {
    case model.DetectDoorOpen:
        ok = model.Bool(b.Open)
    case model.DetectDoorClose:
        ok = model.Bool(!b.Open)
    default:
        // 忽略
        ok = model.True
        return
    }

    usable, message := s.DetectUsable(b)

    if !usable {
        err = errs.CabinetBinNotUsable
    }

    log.Infof("[BIN] [%s - %d, %s] 仓门%s检测: %s, %s", b.Serial, b.Ordinal, s.User, d, ok, message)
    return
}

// DetectBattery 识别电池
func (s *binService) DetectBattery(b *ent.Bin, d model.DetectBattery) (ok model.Bool, err error) {
    fakevoltage, _ := core.Hub.Bean.GetEmptyDeviation()

    switch d {
    case model.DetectBatteryPutin:
        // 检测放入
        ok = model.Bool(b.IsStrictHasBattery(fakevoltage))
    case model.DetectBatteryPutout:
        // 检测取出
        ok = model.Bool(!b.IsStrictHasBattery(fakevoltage))
    default:
        // 忽略
        ok = model.True
        return
    }

    usable, message := s.DetectUsable(b)

    if !usable {
        err = errs.CabinetBinNotUsable
    }

    log.Infof("[BIN] [%s - %d, %s] 电池%s检测: %s, %s", b.Serial, b.Ordinal, s.User, d, ok, message)
    return
}
