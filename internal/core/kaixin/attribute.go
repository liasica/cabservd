// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-29
// Based on cabservd by liasica, magicrolan@qq.com.

package kaixin

import (
    "fmt"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/cabinet"
    "github.com/liasica/go-helpers/silk"
    "github.com/liasica/go-helpers/tools"
    log "github.com/sirupsen/logrus"
    "strconv"
)

type Attributes []*Attribute

// Attribute 属性信息
type Attribute struct {
    SignalData
    DoorID string `json:"doorId,omitempty"` // 柜门ID (可为空)
}

func (d SignalData) ValueString() (str string) {
    str = fmt.Sprintf("%v", d.Value)
    if str == "null" {
        str = ""
    }
    return
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
func (a *Attribute) BinStatus(bin *ent.BinPointer, v string) {
    // 是否异常
    bin.Health = silk.Bool(v != "5")
    // 如果无电池
    if v == "0" {
        bin.BatterySn = silk.String("")
    }
    return
}

func (attrs Attributes) Bins() (items ent.BinPointers) {
    m := make(map[string]*ent.BinPointer)

    for _, attr := range attrs {
        // 原始字符串值
        v := attr.ValueString()

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
            m[attr.DoorID] = bin
        }

        switch attr.ID {
        case SignalBinStatus:
            attr.BinStatus(bin, v)
        case SignalDoorStatus:
            bin.Open = silk.Bool(v == "1")
        case SignalBinUsingStatus:
            bin.Enable = silk.Bool(v == "1")
        case SignalBatterySN:
            bin.BatterySn = silk.String(v)
        case SignalBatteryVoltage:
            bin.Voltage = silk.Float64(tools.StrToFloat64(v))
        case SignalBatteryCurrent:
            c := tools.StrToFloat64(v)
            bin.Current = silk.Float64(c)
        case SignalSOC:
            bin.Soc = silk.Float64(tools.StrToFloat64(v))
        case SignalSOH:
            bin.Soh = silk.Float64(tools.StrToFloat64(v))
        }
    }

    for _, p := range m {
        items = append(items, p)
    }
    return
}

func (attrs Attributes) Cabinet() (cab ent.CabinetPointer, exists bool) {
    for _, attr := range attrs {
        v := attr.ValueString()

        if _, ok := CabinetSignal[attr.ID]; ok {
            exists = true
        }

        switch attr.ID {
        case SignalCabinetStatus:
            m := map[string]cabinet.Status{
                "0": cabinet.StatusPoweron,
                "1": cabinet.StatusIdle,
                "2": cabinet.StatusBusy,
                "3": cabinet.StatusAbnormal,
            }
            cab.Status = silk.Pointer(m[v])
        case SignalLng:
            cab.Lng = silk.Float64(tools.StrToFloat64(v))
        case SignalLat:
            cab.Lat = silk.Float64(tools.StrToFloat64(v))
        case SignalGSM:
            cab.Gsm = silk.Float64(tools.StrToFloat64(v))
        case SignalCabinetVoltage:
            cab.Voltage = silk.Float64(tools.StrToFloat64(v))
        case SignalCabinetCurrent:
            cab.Current = silk.Float64(tools.StrToFloat64(v))
        case SignalCabinetTemp:
            cab.Temperature = silk.Float64(tools.StrToFloat64(v))
        case SignalEnable:
            cab.Enable = silk.Bool(v == "1")
        case SignalElectricity:
            cab.Electricity = silk.Float64(tools.StrToFloat64(v))
        }
    }
    return
}
