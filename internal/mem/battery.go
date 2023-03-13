// Copyright (C) liasica. 2023-present.
//
// Created at 2023-03-11
// Based on cabservd by liasica, magicrolan@qq.com.

package mem

import (
    "math"
    "sync"
)

// 单芯电压记录更新
// TODO 拓邦更新后删除
var (
    // 电池单芯电压
    // string => *Voltage
    batteryVoltage sync.Map
)

type MonVoltage struct {
    Total    float64
    Monomers map[int]float64
}

func VoltageClear(serial string, ordinal int) {
    batteryVoltage.Delete(binKey(serial, ordinal))
}

func VoltageMonUpdate(serial string, ordinal int, index int, voltage float64) {
    key := binKey(serial, ordinal)
    v, ok := batteryVoltage.Load(key)
    if !ok {
        v = &MonVoltage{
            Total:    0,
            Monomers: make(map[int]float64),
        }
    }

    t := v.(*MonVoltage)
    t.Monomers[index] = voltage
    t.Total = 0

    for k := range t.Monomers {
        t.Total += t.Monomers[k]
    }

    batteryVoltage.Store(key, v)
}

func VoltageGet(serial string, ordinal int) float64 {
    v, ok := batteryVoltage.Load(binKey(serial, ordinal))
    if !ok {
        return 0
    }
    return math.Round(v.(*MonVoltage).Total*100.00) / 100.0
}
