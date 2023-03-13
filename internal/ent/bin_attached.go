// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-28
// Based on cabservd by liasica, magicrolan@qq.com.

package ent

import (
    "fmt"
    "github.com/auroraride/adapter"
    "github.com/auroraride/adapter/defs/cabdef"
    "github.com/liasica/go-helpers/silk"
    "strings"
)

// ResetBattery 无电池的时候清除电池信息
func (u *BinMutation) ResetBattery() {
    u.SetCurrent(0)
    u.SetVoltage(0)
    u.SetSoc(0)
    u.SetSoh(0)
}

func (b *BinPointer) ResetBattery() {
    b.BatteryExists = silk.Bool(false)
    b.BatterySn = silk.String("")
    b.Voltage = silk.Float64(0)
    b.Current = silk.Float64(0)
    b.Soc = silk.Float64(0)
    b.Soh = silk.Float64(0)
}

// LostBattery 是否无电池
func (b *BinPointer) LostBattery() bool {
    if b.BatterySn != nil && *b.BatterySn == "" {
        return true
    }
    if b.BatteryExists != nil && !*b.BatteryExists {
        return true
    }

    return false
}

// IsLooseHasBattery 宽松检测有电池
// 任意一项满足即判定为有电池
// 常用于: ?
func (b *Bin) IsLooseHasBattery(fakevoltage, fakecurrent float64) bool {
    return b.BatteryExists || // 在位
        b.BatterySn != "" || // 编码不为空
        b.Voltage > fakevoltage // 电压大于指定值
}

// IsStrictNoBattery 严格检测无电池
// 所有项都满足才视为空仓, 防止误开仓导致电池丢失
// 常用于: 检测仓位是否为无电池仓位
func (b *Bin) IsStrictNoBattery(fakevoltage, fakecurrent float64) bool {
    return !b.BatteryExists && // 不在位
        b.BatterySn == "" && // 编码为空
        b.Voltage <= fakevoltage // 电压小于等于指定值
}

// IsStrictHasBattery 严格检测有电池
// 所有项都满足才判定为有电池
// 常用于: 检测用户电池是否放入或该仓位是否可办理业务
func (b *Bin) IsStrictHasBattery(fakevoltage float64) (has bool) {
    has = b.BatteryExists && // 在位
        b.BatterySn != "" && // 电池编码不为空
        b.Voltage > fakevoltage // 电压大于指定值
    return
}

// IsLooseNoBattery 宽松检测无电池
// 任意一项满足即视为无电池
// 常用于: 检测电池是否取走
func (b *Bin) IsLooseNoBattery(fakevoltage float64) bool {
    return !b.BatteryExists || // 不在位
        b.BatterySn == "" || // 电池编号为空
        b.Voltage <= fakevoltage // 电压小于等于指定值
}

// IsUsable 检查仓位是否可用
func (b *Bin) IsUsable() bool {
    return b.Health && b.Enable
}

// BusinessPossible 检查仓位是否可操作换电
func (b *Bin) BusinessPossible(isFull bool, fakevoltage, fakecurrent, minsoc float64) bool {
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

func (b *Bin) Info() *cabdef.BinInfo {
    return &cabdef.BinInfo{
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

func (b *BinPointer) String() string {
    var builder strings.Builder
    builder.WriteString("仓位[")
    builder.WriteString(*b.Serial)
    builder.WriteString(" - ")
    builder.WriteString(*b.Name)
    builder.WriteString("]变动 ->")

    if b.Open != nil {
        builder.WriteString(" 开门=")
        builder.WriteString(adapter.Bool(*b.Open).String())
    }

    if b.Enable != nil {
        builder.WriteString(" 启用=")
        builder.WriteString(adapter.Bool(*b.Enable).String())
    }

    if b.Health != nil {
        builder.WriteString(" 健康=")
        builder.WriteString(adapter.Bool(*b.Health).String())
    }

    if b.BatteryExists != nil {
        builder.WriteString(" 电池在位=")
        builder.WriteString(adapter.Bool(*b.BatteryExists).String())
    }

    if b.BatterySn != nil {
        builder.WriteString(" 电池=")
        builder.WriteString(*b.BatterySn)
    }

    if b.Voltage != nil {
        builder.WriteString(" 电压=")
        builder.WriteString(fmt.Sprintf("%.2f", *b.Voltage))
    }

    if b.Current != nil {
        builder.WriteString(" 电流=")
        builder.WriteString(fmt.Sprintf("%.2f", *b.Current))
    }

    if b.Soc != nil {
        builder.WriteString(" 容量=")
        builder.WriteString(fmt.Sprintf("%.2f", *b.Soc))
    }

    return builder.String()
}
