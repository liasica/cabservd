// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-21
// Based on cabservd by liasica, magicrolan@qq.com.

package core

import (
    "context"
    "fmt"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/bin"
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
    GetSoC() (v float64, exists bool)
    GetSoH() (v float64, exists bool)
}

func SaveBin(brand, sn string, bin Bin) error {
    ctx := context.Background()
    return SaveBinWithContext(brand, sn, bin, ctx)
}

func SaveBinWithContext(brand, sn string, item Bin, ctx context.Context) (err error) {
    index, exists := item.GetDoorIndex()
    if !exists {
        err = errs.CabinetBinIndexRequired
        return
    }

    uuid := tools.Md5String(fmt.Sprintf("%s_%s_%d", brand, sn, index))

    return ent.Database.Bin.Create().
        SetUUID(uuid).
        SetBrand(brand).
        SetSn(sn).
        SetName(fmt.Sprintf("%d号仓", index+1)).
        SetIndex(index).
        OnConflictColumns(bin.FieldUUID).
        Update(func(u *ent.BinUpsert) {
            // 仓门状态
            if open, ok := item.GetOpen(); ok {
                fmt.Printf("%d open:->%v\n", index, open)
                u.SetOpen(open)
            }

            // 仓位启用状态
            if enable, ok := item.GetEnable(); ok {
                u.SetEnable(enable)
            }

            // 电池编号
            if bs, ok := item.GetBatterySN(); ok {
                fmt.Printf("%d battery:->%v\n", index, bs)
                u.SetBatterySn(bs)
                if bs == "" {
                    // 无电池的时候清除电池信息
                    // TODO: 是否有必要?
                    u.SetCurrent(0).SetVoltage(0).SetSoc(0).SetSoh(0)
                }
            }

            // 电压
            if v, ok := item.GetVoltage(); ok {
                u.SetVoltage(v)
            }

            // 电流
            if v, ok := item.GetCurrent(); ok {
                u.SetCurrent(v)
            }

            // 电量
            if v, ok := item.GetSoC(); ok {
                u.SetSoc(v)
            }

            // 健康
            if v, ok := item.GetSoH(); ok {
                u.SetSoh(v)
            }
        }).
        UpdateUUID().
        Exec(ctx)
}

// ResetBins 重置电柜仓位信息
func ResetBins(sn string) error {
    return ent.Database.Bin.Update().
        Where(bin.Sn(sn)).
        SetBatterySn("").
        SetSoc(0).
        SetSoh(0).
        SetVoltage(0).
        SetCurrent(0).
        // SetEnable(true). // TODO 是否单独设置LOCK
        SetOpen(false).
        Exec(context.Background())
}
