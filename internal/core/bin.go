// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-21
// Based on cabservd by liasica, magicrolan@qq.com.

package core

import (
    "context"
    "fmt"
    "github.com/auroraride/cabservd/internal/ent"
)

type Bin interface {
    GetBrand() (string, bool)
    GetSN() (string, bool)
    GetOpen() (bool, bool)
    GetDoorIndex() (int, bool)
}

// // 品牌
// Brand string `json:"brand,omitempty"`
// // 电柜设备序列号
// Sn string `json:"sn,omitempty"`
// // 仓门(N号仓)
// Door string `json:"door,omitempty"`
// // 仓门编号(从0开始)
// DoorIndex int `json:"door_index,omitempty"`
// // 仓门是否开启
// DoorOpen bool `json:"door_open,omitempty"`
// // 电池序列号
// BatterySn *string `json:"battery_sn,omitempty"`
// // 当前电压
// Voltage *float64 `json:"voltage,omitempty"`
// // 当前电流
// Current *float64 `json:"current,omitempty"`

func SaveBin(bin Bin) (*ent.CabinetBin, error) {
    ctx := context.Background()
    return SaveBinWithContext(bin, ctx)
}

func SaveBinWithContext(bin Bin, ctx context.Context) (*ent.CabinetBin, error) {
    q := ent.Database.CabinetBin.Create().
        OnConflictColumns().
        SetBrand(bin.GetBrand()).
        SetSn(bin.GetSN()).
        SetName(fmt.Sprintf("%d号仓", bin.Index+1)).
        SetIndex(bin.Index).
        SetOpen(bin.Open)

    // 电池序列号
    if bin.BatterySn == nil {
        q.ClearBatterySn()
    } else {
        q.SetBatterySn(*bin.BatterySn)
    }

    // 电压
    return q.Save(ctx)
}
