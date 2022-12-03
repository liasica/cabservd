// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-03
// Based on cabservd by liasica, magicrolan@qq.com.

package types

type BinInfo struct {
    Ordinal   int     `json:"ordinal,omitempty"`   // 仓位序号
    BatterySN string  `json:"batterySn,omitempty"` // 电池编码
    Voltage   float64 `json:"voltage,omitempty"`   // 电压
    Current   float64 `json:"current,omitempty"`   // 电流
    Soc       float64 `json:"soc,omitempty"`       // 电量
    Soh       float64 `json:"soh,omitempty"`       // 健康
}
