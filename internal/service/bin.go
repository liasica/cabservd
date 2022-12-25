// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-03
// Based on cabservd by liasica, magicrolan@qq.com.

package service

import (
    "context"
    "github.com/auroraride/cabservd/internal/core"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/bin"
    "github.com/auroraride/cabservd/internal/errs"
    "github.com/auroraride/cabservd/internal/types"
    "github.com/liasica/go-helpers/silk"
)

type binService struct {
    ordinal int
    bin     *ent.Bin
    cabinet *ent.Cabinet
    ctx     context.Context
}

func NewBin(cab *ent.Cabinet, ordinal int) *binService {
    s := &binService{
        cabinet: cab,
        ordinal: ordinal,
        ctx:     context.WithValue(context.Background(), "cabinet", cab),
    }

    ctx := context.Background()
    cb, _ := ent.Database.Bin.Query().Where(bin.Serial(cab.Serial), bin.Brand(cab.Brand), bin.Ordinal(ordinal)).First(ctx)
    ctx = context.WithValue(ctx, "bin", cb)

    s.ctx = ctx
    s.bin = cb
    return s
}

// Enable 控制仓位启用/禁用
func (s *binService) Enable(enable bool) error {
    var t types.ControlType

    switch enable {
    case true:
        t = types.ControlTypeBinEnable
    case false:
        t = types.ControlTypeBinDisable
    default:
        return errs.CabinetControlParamError
    }

    return core.Hub.Control(&types.ControlRequest{
        Type:    t,
        Serial:  s.bin.Serial,
        Ordinal: silk.Int(s.bin.Ordinal),
    })
}

// Open 打开仓门
func (s *binService) Open() error {
    return core.Hub.Control(&types.ControlRequest{
        Type:    types.ControlTypeBinOpen,
        Serial:  s.bin.Serial,
        Ordinal: silk.Int(s.bin.Ordinal),
    })
}
