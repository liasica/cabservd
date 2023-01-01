// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-28
// Based on cabservd by liasica, magicrolan@qq.com.

package ent

import (
    "fmt"
    "github.com/auroraride/adapter"
)

// ResetBattery 无电池的时候清除电池信息
// TODO: 是否有必要?
func (u *BinUpsert) ResetBattery() *BinUpsert {
    u.SetCurrent(0).SetVoltage(0).SetSoc(0).SetSoh(0)
    return u
}

// IsLooseHasBattery 宽松检测有电池
// 任意一项满足即判定为有电池
// 常用于: ?
func (b *Bin) IsLooseHasBattery(fakevoltage, fakecurrent float64) bool {
    return b.BatteryExists || // 在位
        b.BatterySn != "" || // 编码不为空
        b.Voltage > fakevoltage || // 电压大于指定值
        b.Current > fakecurrent || // 电流大于指定值
        b.Soc > 0 // 容量大于0
}

// IsStrictNoBattery 严格检测无电池
// 所有项都满足才视为空仓, 防止误开仓导致电池丢失
// 常用于: 检测仓位是否为无电池仓位
func (b *Bin) IsStrictNoBattery(fakevoltage, fakecurrent float64) bool {
    return !b.BatteryExists && // 不在位
        b.BatterySn == "" && // 编码为空
        b.Voltage <= fakevoltage && // 电压小于等于指定值
        b.Current <= fakecurrent && // 电流小于等于指定值
        b.Soc <= 0 // 容量小于等于0
}

// IsStrictHasBattery 严格检测有电池
// 所有项都满足才判定为有电池
// 常用于: 检测用户电池是否放入
func (b *Bin) IsStrictHasBattery(fakevoltage float64) (has bool) {
    has = b.BatteryExists && // 在位
        b.BatterySn != "" && // 电池编码不为空
        b.Voltage > fakevoltage && // 电压大于指定值
        b.Soc > 0 // 容量不为0
    return
}

// IsLooseNoBattery 宽松检测无电池
// 任意一项满足即视为无电池
// 常用于: 检测电池是否取走
func (b *Bin) IsLooseNoBattery(fakevoltage float64) bool {
    return !b.BatteryExists || // 不在位
        b.BatterySn == "" || // 电池编号为空
        b.Voltage <= fakevoltage || // 电压小于等于指定值
        b.Soc <= 0 // 容量小于等于0
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
        // 严格检查是否为空仓
        return b.IsStrictNoBattery(fakevoltage, fakecurrent)
    }
}

func (b *Bin) Info() *adapter.BinInfo {
    return &adapter.BinInfo{
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

func (b *Bin) GetID() uint64 {
    return b.ID
}

func (b *Bin) GetSerial() string {
    return b.Serial
}

func (b *Bin) GetListenerKey() string {
    return fmt.Sprintf("%s-%d", b.GetTableName(), b.ID)
}
