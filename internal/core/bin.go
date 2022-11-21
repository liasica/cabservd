// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-21
// Based on cabservd by liasica, magicrolan@qq.com.

package core

import (
    "context"
    "fmt"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/cabinetbin"
    "github.com/auroraride/cabservd/internal/errs"
    "github.com/liasica/go-helpers/tools"
)

type Bin interface {
    GetOpen() (v bool, exists bool)
    GetEnable() (v bool, exists bool)
    GetDoorIndex() (v int, exists bool)
    GetBatterySN() (v string, exists bool)
    GetVoltage() (v float64, exists bool)
    GetCurrent() (v float64, exists bool)
}

func SaveBin(brand, sn string, bin Bin) (*ent.CabinetBin, error) {
    ctx := context.Background()
    return SaveBinWithContext(brand, sn, bin, ctx)
}

func SaveBinWithContext(brand, sn string, bin Bin, ctx context.Context) (cb *ent.CabinetBin, err error) {
    index, exists := bin.GetDoorIndex()
    if !exists {
        err = errs.CabinetBinIndexRequired
        return
    }

    uuid := tools.Md5String(fmt.Sprintf("%s_%s_%d", brand, sn, index))

    q := ent.Database.CabinetBin.Create().
        SetUUID(uuid).
        SetBrand(brand).
        SetBrand(brand).
        SetSn(sn).
        SetName(fmt.Sprintf("%d号仓", index+1)).
        SetIndex(index).
        OnConflictColumns(cabinetbin.FieldUUID).
        UpdateNewValues()

    // 仓门状态
    if open, ok := bin.GetOpen(); ok {
        q.SetOpen(open)
    }

    // 仓位启用状态
    if enable, ok := bin.GetEnable(); ok {
        q.SetEnable(enable)
    }

    // 电池编号
    if bs, ok := bin.GetBatterySN(); ok {
        q.SetBatterySn(bs)
        if bs == "" {
            // 无电池的时候将电压和电流标记为0
            q.SetVoltage(0).SetCurrent(0)
        }
    }

    // 电压
    if v, ok := bin.GetVoltage(); ok {
        q.SetVoltage(v)
    }

    // 电流
    if v, ok := bin.GetCurrent(); ok {
        q.SetCurrent(v)
    }

    return q.Save(ctx)
}
