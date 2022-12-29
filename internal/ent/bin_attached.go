// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-28
// Based on cabservd by liasica, magicrolan@qq.com.

package ent

import "github.com/auroraride/adapter/model"

// ResetBattery 无电池的时候清除电池信息
// TODO: 是否有必要?
func (u *BinUpsert) ResetBattery() *BinUpsert {
    u.SetCurrent(0).SetVoltage(0).SetSoc(0).SetSoh(0)
    return u
}

// IsLooseHasBattery 是否有电池
// 宽松检测, 任意一项满足即判定为有电池, 常用在换电第一步判定仓位是否为空
// fakevoltage 指定电压
// fakecurrent 指定电流
// 在位 或 编码不为空 或 电压大于指定值 或 电流大于指定值 或 容量大于0
func (b *Bin) IsLooseHasBattery(fakevoltage, fakecurrent float64) bool {
    return b.BatteryExists || // 在位
        b.BatterySn != "" || // 编码不为空
        b.Voltage > fakevoltage || // 电压大于指定值
        b.Current > fakecurrent || // 电流大于指定值
        b.Soc > 0 // 容量大于0
}

// IsStrictHasBattery 是否有电池
// 严格检测, 所有选项都满足才判定为有电池, 常用在换电第三步检测电池是否放入
// 在位 并且 电池编码不为空 并且 电压大于指定电压 并且 容量大于0
func (b *Bin) IsStrictHasBattery(fakevoltage float64) (has bool) {
    has = b.BatteryExists &&
        b.BatterySn != "" &&
        b.Voltage > fakevoltage &&
        b.Soc > 0
    return
}

// IsUsable 检查仓位是否可用
func (b *Bin) IsUsable() bool {
    return b.Health && b.Enable
}

// ExchangePossible 检查仓位是否可操作换电
func (b *Bin) ExchangePossible(isFull bool, fakevoltage, fakecurrent, minsoc float64) bool {
    if !b.IsUsable() || b.Open {
        return false
    }
    if isFull {
        // 满仓严格检查是否有电池并且电量高于指定电量
        return b.IsStrictHasBattery(fakevoltage) && b.Soc >= minsoc
    } else {
        // 空仓宽松检查是否有电池
        return !b.IsLooseHasBattery(fakevoltage, fakecurrent)
    }
}

func (b *Bin) Info() *model.BinInfo {
    return &model.BinInfo{
        Ordinal:       b.Ordinal,
        BatterySN:     b.BatterySn,
        Voltage:       b.Voltage,
        Current:       b.Current,
        Soc:           b.Soc,
        Soh:           b.Soh,
        Health:        b.Health,
        Enable:        b.Enable,
        Open:          b.Open,
        BatteryExists: b.BatteryExists,
    }
}
