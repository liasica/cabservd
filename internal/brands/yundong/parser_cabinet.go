// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-15
// Based on yundong by liasica, magicrolan@qq.com.

package yundong

import (
    "encoding/binary"
    "github.com/auroraride/adapter"
    "github.com/auroraride/cabservd/internal/core"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/cabinet"
    "github.com/liasica/go-helpers/silk"
    "strconv"
    "time"
)

var cabinetConfig = []Field{
    {Length: 4, Name: "Time"},
    {Length: 1, Name: "Strength4G"},
    {Length: 1, Name: "StrengthWifi"},
    {Length: 1, Name: "Powermode"},
    {Length: 1, Name: "BatteryNum"},
    {Length: 4, Name: "Current"},
    {Length: 1, Name: "Temperature"},
}

const (
    CabinetAlarmHighTemperature = iota // 柜体高温告警
    CabinetAlarmSmoke                  // 柜体烟雾告警
    CabinetAlarmFlooded                // 柜体进水告警
)

type CabinetAlarms []int

func (arr CabinetAlarms) String() string {
    buf := adapter.NewBuffer()
    defer adapter.ReleaseBuffer(buf)

    for i, v := range arr {
        if v == 1 {
            if buf.Len() > 0 {
                buf.WriteString("|")
            }
            switch i {
            case 0:
                buf.WriteString("高温")
            case 1:
                buf.WriteString("烟雾")
            case 2:
                buf.WriteString("进水")
            }
        }
    }

    return buf.String()
}

type CabinetData struct {
    Serial       string        `json:"serial"`
    Time         time.Time     `json:"time"`         // 时间戳(s)
    Strength4G   byte          `json:"strength4g"`   // 4G信号强度 (0-31)
    StrengthWifi byte          `json:"strengthWifi"` // Wifi信号强度 (0-100)
    Powermode    byte          `json:"powermode"`    // 供电模式, 1:电源供电 2:电池供电
    BatteryNum   byte          `json:"batteryNum"`   // 在位电池数量
    Current      uint32        `json:"current"`      // 总电流(mA)
    Temperature  int8          `json:"temperature"`  // 环境温度
    Bins         Bins          `json:"bins"`         // 非空电池仓
    Key          string        `json:"key"`          // AES-128-ECB密钥
    Alarms       CabinetAlarms `json:"alarms"`       // 柜体告警
}

func (d *CabinetData) String() string {
    buf := adapter.NewBuffer()
    defer adapter.ReleaseBuffer(buf)

    buf.WriteString("Serial=")
    buf.WriteString(d.Serial)

    buf.WriteString(", 4G信号强度=")
    buf.WriteString(strconv.Itoa(int(d.Strength4G)))

    buf.WriteString(", Wifi信号强度=")
    buf.WriteString(strconv.Itoa(int(d.StrengthWifi)))

    buf.WriteString(", 供电模式=")
    switch d.Powermode {
    case 1:
        buf.WriteString("电源供电")
    case 2:
        buf.WriteString("电池供电")
    }

    buf.WriteString(", 在位电池数量=")
    buf.WriteString(strconv.Itoa(int(d.BatteryNum)))

    buf.WriteString(", 总电流(mA)=")
    buf.WriteString(strconv.Itoa(int(d.Current)))

    buf.WriteString(", 环境温度(℃)=")
    buf.WriteString(strconv.Itoa(int(d.Temperature)))

    // TODO bins
    for _, b := range d.Bins {
        buf.WriteString(", ")
        buf.WriteString(b.String())
    }

    buf.WriteString(", key=")
    buf.WriteString(d.Key)

    buf.WriteString(", 柜体告警=")
    buf.WriteString(d.Alarms.String())

    return buf.String()
}

func (d *CabinetData) SetField(b []byte, name string) {
    switch name {
    case "Time":
        d.Time = time.Unix(int64(binary.BigEndian.Uint32(b)), 0).In(time.Local)
    case "Strength4G":
        d.Strength4G = b[0]
    case "StrengthWifi":
        d.StrengthWifi = b[0]
    case "Powermode":
        d.Powermode = b[0]
    case "BatteryNum":
        d.BatteryNum = b[0]
    case "Current":
        d.Current = binary.BigEndian.Uint32(b)
    case "Temperature":
        d.Temperature = int8(b[0])
    case "Key":
        d.Key = adapter.ConvertBytes2String(b)
    case "Alarms":
        d.Alarms = adapter.GetTrueBits(b[0], 2)
    }
}

func (d *CabinetData) GetSerial() (string, bool) {
    return d.Serial, d.Serial != ""
}

func (d *CabinetData) GetCabinet() (cp *ent.CabinetPointer, exists bool) {
    if d.Serial == "" {
        return
    }

    status := cabinet.StatusNormal
    if len(d.Alarms) > 0 {
        status = cabinet.StatusAbnormal
    }

    return &ent.CabinetPointer{
        Online: silk.Bool(true),
        Power:  silk.Bool(d.Powermode == 1),
        Serial: silk.String(d.Serial),
        Status: &status,
        Enable: silk.Bool(true),
    }, true
}

func (d *CabinetData) GetBins() (ebs ent.BinPointers) {
    sp := silk.Pointer(d.Serial)
    ebs = make(ent.BinPointers, len(d.Bins))
    for i, b := range d.Bins {
        isOpen := b.Status == BinStatusOpen || b.Status == BinStatusOpenAndDisabled
        isEnable := b.Status != BinStatusOpenAndDisabled && b.Status != BinStatusCloseAndDisabled
        // 启用 并且 无警告 并且 充电开
        isHealth := isEnable && len(b.Alarms) == 0
        ordinal := int(b.Index) + 1
        ebs[i] = &ent.BinPointer{
            Serial:        sp,
            Ordinal:       silk.Int(ordinal),
            Name:          silk.String(strconv.Itoa(ordinal) + "号仓"),
            Open:          silk.Bool(isOpen),
            Enable:        silk.Bool(isEnable),
            Health:        silk.Bool(isHealth),
            BatteryExists: silk.Bool(b.BatterySN != ""),
            BatterySn:     silk.String(b.BatterySN),
            Voltage:       silk.Float64(float64(b.Voltage) / 1000.0),
            Current:       silk.Float64(float64(b.Chargei) / 1000.0),
            Soc:           silk.Float64(float64(b.SOC)),
            Soh:           silk.Float64(float64(b.SOH)),
        }
    }
    return
}

func (p *Parser) CabinetData(h *Handler, serial string, b []byte) string {
    data := new(CabinetData)
    index := Parse(data, b, cabinetConfig)
    binBytes := b[index : len(b)-17]
    data.Bins = p.Bins(binBytes)

    index += len(binBytes)
    data.SetField(b[index:index+16], "Key")
    data.SetField(b[len(b)-1:], "Alarms")

    if serial != "" {
        data.Serial = serial
        core.UpdateCabinet(h, data)
    }

    return data.String()
}
