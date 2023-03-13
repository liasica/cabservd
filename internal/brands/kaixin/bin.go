// Copyright (C) liasica. 2023-present.
//
// Created at 2023-03-09
// Based on cabservd by liasica, magicrolan@qq.com.

package kaixin

import (
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/tower"
    "github.com/liasica/go-helpers/silk"
    "github.com/liasica/go-helpers/tools"
)

const (
    SignalBatteryExists tower.Signal = "02140001" // 电池在位检测 0:无电池 1:有电池
)

var binSignals = map[tower.Signal]tower.BinSignalFunc{
    SignalBatteryExists: func(b *ent.BinPointer, attr *tower.Attribute, v string) {
        b.BatteryExists = silk.Bool(v == "1")
    },
    tower.SignalBatteryVoltage: func(b *ent.BinPointer, attr *tower.Attribute, v string) {
        vf := tools.StrToFloat64(v)
        b.Voltage = silk.Float64(vf)
    },
    tower.SignalBatteryCurrent: func(b *ent.BinPointer, attr *tower.Attribute, v string) {
        vf := tools.StrToFloat64(v)
        b.Current = silk.Float64(vf)
    },
}
