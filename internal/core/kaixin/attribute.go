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

func (a *Attribute) GetOrdinal() (ordinal int, exists bool) {
    if a.DoorID == "" {
        return
    }

    var err error
    ordinal, err = strconv.Atoi(a.DoorID)
    if err != nil {
        log.Errorf("仓位解析失败")
        return
    }

    exists = true
    return
}

// BinStatus 获取仓位状态
// hasBattery 是否有电池
func (a *Attribute) BinStatus(bin *ent.BinPointer, v string) {
    // 是否异常
    bin.Health = silk.Bool(v != "5")
    // // TODO 如果无电池
    // if v == "0" {
    //     bin.BatterySn = silk.String("")
    // }
    return
}

func (req ReportRequest) Bins() (items ent.BinPointers) {
    m := make(map[string]*ent.BinPointer)

    for _, attr := range req.AttrList {
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
            bin = &ent.BinPointer{Ordinal: silk.Int(ordinal)}
            m[attr.DoorID] = bin
        }

        // TODO 电池在位检测信号量
        switch attr.ID {
        case SignalBinStatus:
            attr.BinStatus(bin, v)
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
            fmt.Println(">>>>>>>>>> [V]", vf, "<<<<<<<<<<")
            bin.Voltage = silk.Float64(vf)
        case SignalBatteryCurrent:
            vf := tools.StrToFloat64(v)
            fmt.Println(">>>>>>>>>> [A]", vf, "<<<<<<<<<<")
            bin.Current = silk.Float64(vf)
        case SignalSOC:
            vf := tools.StrToFloat64(v)
            fmt.Println(">>>>>>>>>> [SOC]", vf, "<<<<<<<<<<")
            bin.Soc = silk.Float64(vf)
        case SignalSOH:
            vf := tools.StrToFloat64(v)
            fmt.Println(">>>>>>>>>> [SOH]", vf, "<<<<<<<<<<")
            bin.Soh = silk.Float64(vf)
        }
    }

    for _, p := range m {
        items = append(items, p)
    }
    return
}

func (req ReportRequest) Cabinet() (cab ent.CabinetPointer, exists bool) {
    // 如果是全量上报, 标记电柜在线
    if req.IsFull == ReportCateFull {
        cab.Online = silk.Bool(true)
    }

    // 解析详细属性
    for _, attr := range req.AttrList {
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
        }
    }
    return
}
