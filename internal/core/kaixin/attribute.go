// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-29
// Based on cabservd by liasica, magicrolan@qq.com.

package kaixin

import (
    "fmt"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/liasica/go-helpers/silk"
    log "github.com/sirupsen/logrus"
    "strconv"
)

type Attributes []*Attribute

// Attribute 属性信息
type Attribute struct {
    SignalData
    DoorID string `json:"doorId,omitempty"` // 柜门ID (可为空)
}

func (a *Attribute) GetDoorIndex() (index int, exists bool) {
    if a.DoorID == "" {
        return
    }

    exists = true
    id, err := strconv.Atoi(a.DoorID)
    if err != nil {
        log.Errorf("仓位解析失败")
    }
    index = id - 1
    return
}

// BinStatus 获取仓位状态
// hasBattery 是否有电池
func (a *Attribute) BinStatus(bin *ent.BinPointer) {
    // 是否异常
    bin.Health = silk.Bool(a.RawValue != BinStatusException)
    // 如果无电池
    if a.RawValue == BinStatusNoBattery {
        bin.BatterySn = silk.String("")
    }
    return
}

func (attrs Attributes) Bins() (items ent.BinPointers) {
    m := make(map[string]*ent.BinPointer)

    for _, attr := range attrs {
        // 原始字符串值
        attr.RawValue = fmt.Sprintf("%v", attr.Value)

        // 获取仓位Index
        index, exists := attr.GetDoorIndex()
        // 如果没有仓门信息, 直接跳过
        if !exists {
            continue
        }

        // 查询是否存在仓位信息
        bin, ok := m[attr.DoorID]
        if !ok {
            bin = &ent.BinPointer{Index: silk.Int(index)}
        }

        switch attr.ID {
        case SignalBinStatus:
            attr.BinStatus(bin)
        case SignalDoorStatus:
            bin.Open = silk.Bool(attr.RawValue == DoorStatusOpen)
        case SignalBinUsingStatus:
            bin.Enable = silk.Bool(attr.RawValue == BinUsingEnable)
        case SignalBatterySN:
            bin.BatterySn = silk.String(attr.RawValue)
        case SignalBatteryVoltage:
            bin.Voltage = silk.Float64(attr.ValueFloat64())
        case SignalBatteryCurrent:
            bin.Current = silk.Float64(attr.ValueFloat64())
        case SignalSOC:
            bin.Soc = silk.Float64(attr.ValueFloat64())
        case SignalSOH:
            bin.Soh = silk.Float64(attr.ValueFloat64())
        }
    }

    for _, p := range m {
        items = append(items, p)
    }
    return
}
