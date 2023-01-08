// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-07
// Based on cabservd by liasica, magicrolan@qq.com.

package kaixin

import (
    "fmt"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/cabinet"
    "github.com/liasica/go-helpers/silk"
    "github.com/liasica/go-helpers/tools"
)

// GetSerial 获取电柜编码
func (r *Request) GetSerial() (string, bool) {
    return r.DevID, true
}

// GetCabinet 获取电柜信息
func (r *Request) GetCabinet() (cab *ent.CabinetPointer, exists bool) {
    cab = &ent.CabinetPointer{
        Serial: silk.String(r.DevID),
    }

    // 如果是全量上报, 标记电柜在线
    if r.IsFull == ReportCateFull {
        cab.Online = silk.Bool(true)
    }

    // 解析详细属性
    for _, attr := range r.AttrList {
        v := attr.ValueString()

        if _, ok := CabinetSignal[attr.ID]; ok {
            exists = true
        }

        switch attr.ID {
        case SignalCabinetStatus:
            m := map[string]cabinet.Status{
                "0": cabinet.StatusIdle,
                "1": cabinet.StatusIdle,
                "2": cabinet.StatusBusy,
                "3": cabinet.StatusBusy,
                "4": cabinet.StatusBusy,
                "5": cabinet.StatusAbnormal,
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
        case SignalPower:
            cab.Power = silk.Bool(v == "0")
        }
    }
    return
}

// GetBins 获取仓位列表信息
func (r *Request) GetBins() (items ent.BinPointers) {
    m := make(map[string]*ent.BinPointer)

    for _, attr := range r.AttrList {
        // 原始字符串值
        v := attr.ValueString()

        // 获取仓位序号
        ordinal, exists := attr.GetOrdinal()
        // 如果没有仓门信息, 直接跳过
        if !exists {
            continue
        }

        // 查询是否存在仓位信息
        bin, ok := m[attr.DoorID]
        if !ok {
            bin = &ent.BinPointer{
                Serial:  silk.String(r.DevID),
                Ordinal: silk.Int(ordinal),
                Name:    silk.String(fmt.Sprintf("%d号仓", ordinal)),
            }
            m[attr.DoorID] = bin
        }

        // TODO 电池在位检测信号量
        switch attr.ID {
        case SignalBinStatus:
            bin.Health = silk.Bool(v != "5")
        case SignalBinDoorStatus:
            bin.Open = silk.Bool(v == "1")
        case SignalBinEnable:
            bin.Enable = silk.Bool(v == "1")
        case SignalBatteryExists:
            bin.BatteryExists = silk.Bool(v == "1")
        case SignalBatterySN:
            bin.BatterySn = silk.String(v)
        case SignalBatteryVoltage:
            vf := tools.StrToFloat64(v)
            bin.Voltage = silk.Float64(vf)
        case SignalBatteryCurrent:
            vf := tools.StrToFloat64(v)
            bin.Current = silk.Float64(vf)
        case SignalSOC:
            vf := tools.StrToFloat64(v)
            bin.Soc = silk.Float64(vf)
        case SignalSOH:
            vf := tools.StrToFloat64(v)
            bin.Soh = silk.Float64(vf)
        }
    }

    for _, p := range m {
        items = append(items, p)
    }
    return
}