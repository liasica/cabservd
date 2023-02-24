// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-15
// Based on yundong by liasica, magicrolan@qq.com.

package parser

import (
    "bytes"
    "encoding/binary"
    "github.com/auroraride/adapter"
    "strconv"
)

var binConfig = []Field{
    {Length: 1, Name: "Index"},
    {Length: 1, Name: "Status"},
    {Length: 1, Name: "Mode"},
    {Length: 24, Name: "BatterySN"},
    {Length: 4, Name: "Alarms"},
    {Length: 1, Name: "ChargeTemp"},
    {Length: 4, Name: "Chargei"},
    {Length: 2, Name: "SOP"},
    {Length: 1, Name: "SOC"},
    {Length: 1, Name: "SOH"},
    {Length: 2, Name: "ChargeTimes"},
    {Length: 2, Name: "DischargeTimes"},
    {Length: 16, Name: "BatteryTemp"},
    {Length: 4, Name: "Voltage"},
    {Length: 48, Name: "MonVoltage"},
    {Length: 4, Name: "Balance"},
    {Length: 2, Name: "Remaining"},
    {Length: 1, Name: "BatteryStatus"},
    {Length: 1, Name: "MosStatus"},
    {Length: 1, Name: "ChargerStatus"},
    {Length: 4, Name: "CurrentSet"},
    {Length: 1, Name: "MicroSwitchStatus"},
    {Length: 1, Name: "AmbientTemp"},
    {Length: 1, Name: "FanStatus"},
    {Length: 1, Name: "AerosolsStatus"},
}

// const (
//     BinStatusClose       = iota // 关
//     BinStatusOpen               // 开
//     BinStatusGrip               // 夹手
//     BinStatusFault              // 故障
//     BinStatusInsert             // 置入
//     BinStatusCanCharge          // 充电使能
//     BinStatusCanExchange        // 换电使能
// )

// type BinStatus []int
//
// func (arr BinStatus) String() string {
//     buf := adapter.NewBuffer()
//     defer adapter.ReleaseBuffer(buf)
//
//     for i, v := range arr {
//         if v == 1 {
//             if buf.Len() > 0 {
//                 buf.WriteString("|")
//             }
//             switch i {
//             case 0:
//                 buf.WriteString("关")
//             case 1:
//                 buf.WriteString("开")
//             case 2:
//                 buf.WriteString("夹手")
//             case 3:
//                 buf.WriteString("故障")
//             case 4:
//                 buf.WriteString("置入")
//             case 5:
//                 buf.WriteString("可充电")
//             case 6:
//                 buf.WriteString("可换电")
//             }
//         }
//     }
//
//     return buf.String()
// }

const (
    BinAlarmUndervoltage                = iota // 总电欠压
    BinAlarmBatteryUndervoltage                // 单节电池欠压
    BinAlarmOvervoltage                        // 总电压过压
    BinAlarmBatteryOvervoltage                 // 单节电池过压
    BinAlarmChargingShortCircuit               // 充电短路
    BinAlarmDischargeShortCircuit              // 放电短路
    BinAlarmChargingOvercurrent                // 充电过流
    BinAlarmDischargeOvercurrent               // 放电过流
    BinAlarmChargeMosHighTemperature           // 充电MOS高温
    BinAlarmDisChargeMosHighTemperature        // 放电MOS高温
    BinAlarmPos1HighTemperature                // 位置1高温
    BinAlarmPos2HighTemperature                // 位置2高温
    BinAlarmPos3HighTemperature                // 位置3高温
    BinAlarmPos4HighTemperature                // 位置4高温
    BinAlarmPoweroff                           // 断电
    BinAlarmVoltage                            // 压差
)

type BinAlarms []int

func (arr BinAlarms) String() string {
    buf := adapter.NewBuffer()
    defer adapter.ReleaseBuffer(buf)

    for i, v := range arr {
        if v == 1 {
            if buf.Len() > 0 {
                buf.WriteString("|")
            }
            switch i {
            case 0:
                buf.WriteString("总电欠压")
            case 1:
                buf.WriteString("单节电池欠压")
            case 2:
                buf.WriteString("总电压过压")
            case 3:
                buf.WriteString("单节电池过压")
            case 4:
                buf.WriteString("充电短路")
            case 5:
                buf.WriteString("放电短路")
            case 6:
                buf.WriteString("充电过流")
            case 7:
                buf.WriteString("放电过流")
            case 8:
                buf.WriteString("充电MOS高温")
            case 9:
                buf.WriteString("放电MOS高温")
            case 10:
                buf.WriteString("位置1高温")
            case 11:
                buf.WriteString("位置2高温")
            case 12:
                buf.WriteString("位置3高温")
            case 13:
                buf.WriteString("位置4高温")
            case 14:
                buf.WriteString("断电")
            case 15:
                buf.WriteString("压差")
            }
        }
    }

    return buf.String()
}

type MosStatus [2]int

func (arr MosStatus) String() string {
    buf := adapter.NewBuffer()
    defer adapter.ReleaseBuffer(buf)

    buf.WriteString("充电")
    if arr[0] == 1 {
        buf.WriteString("开")
    } else {
        buf.WriteString("关")
    }
    buf.WriteString(" / 放电")
    if arr[1] == 1 {
        buf.WriteString("开")
    } else {
        buf.WriteString("关")
    }

    return buf.String()
}

type Bins []*BinData

type BinStatus byte

const (
    BinStatusClose            BinStatus = 0 // 仓门关
    BinStatusOpen             BinStatus = 1 // 仓门开
    BinStatusCloseAndDisabled BinStatus = 2 // 仓门关&禁用
    BinStatusOpenAndDisabled  BinStatus = 3 // 仓门开&禁用
)

func (s BinStatus) String() string {
    switch s {
    default:
        return "-"
    case 0:
        return "仓门关"
    case 1:
        return "仓门开"
    case 2:
        return "仓门关&禁用"
    case 3:
        return "仓门开&禁用"
    }
}

type BinData struct {
    Index             byte      `json:"index"`               // 仓位序号, 从0开始
    Status            BinStatus `json:"status"`              // 仓位状态, 0:仓门关 1:仓门开 2:仓门关&禁用 3:仓门开&禁用
    Mode              byte      `json:"mode"`                // 仓位休眠状态, 1:正常 2:休眠 3:深度休眠
    BatterySN         string    `json:"batterySn,omitempty"` // 电池编号
    Alarms            BinAlarms `json:"alarms,omitempty"`    // 仓位告警
    ChargeTemp        byte      `json:"chargeTemp"`          // 充电机温度(摄氏度)
    Chargei           uint32    `json:"chargei"`             // 充电电流(mA)
    SOP               uint16    `json:"sop"`                 // 电池电量(mAH)
    SOC               byte      `json:"soc"`                 // 剩余容量(%)
    SOH               byte      `json:"soh"`                 // 电池健康状态(%)
    ChargeTimes       uint16    `json:"chargeTimes"`         // 充电次数
    DischargeTimes    uint16    `json:"dischargeTimes"`      // 放电次数
    BatteryTemp       []int     `json:"batteryTemp"`         // 电池温度采样点温度,单位为摄氏度，目前为 3 个采样点
    Voltage           uint32    `json:"voltage"`             // 电压(mV)
    MonVoltage        []uint16  `json:"monVoltage"`          // 电池各串电压(mV), 目前为13 串
    Balance           []int     `json:"balance"`             // 电池各串均衡状态, 从bit0到bit15表示地1串到第16串 0:无均衡 1:均衡
    Remaining         uint16    `json:"remaining"`           // 充满电剩余时长(m)
    BatteryStatus     byte      `json:"batteryStatus"`       // 电池状态, 1:充电 2:放电 3:静止
    MosStatus         MosStatus `json:"mosStatus"`           // Mos状态
    ChargerStatus     byte      `json:"chargerStatus"`       // 充电机状态, 0:关机 1:开机 2:故障
    CurrentSet        uint32    `json:"currentSet"`          // 当前设置的电流大小(mA)
    MicroSwitchStatus byte      `json:"microSwitchStatus"`   // 微动开关状态, 0:未触发 1:已触发 2:故障
    AmbientTemp       byte      `json:"ambientTemp"`         // 电池仓温度
    FanStatus         byte      `json:"fanStatus"`           // 风扇信号反馈, 0:未启动 1:已启动
    AerosolsStatus    byte      `json:"aerosolsStatus"`      // 气溶胶状态, 0:未触发 1:已触发
}

func (d *BinData) String() string {
    buf := adapter.NewBuffer()
    defer adapter.ReleaseBuffer(buf)

    buf.WriteString("仓位序号=")
    buf.WriteString(strconv.Itoa(int(d.Index)))

    buf.WriteString(", 仓位状态=")
    buf.WriteString(d.Status.String())

    buf.WriteString(", 仓位休眠状态=")
    switch d.Mode {
    default:
        buf.WriteString("-")
    case 1:
        buf.WriteString("正常")
    case 2:
        buf.WriteString("休眠")
    case 3:
        buf.WriteString("深度休眠")
    }

    buf.WriteString(", 电池编号=")
    buf.WriteString(d.BatterySN)

    buf.WriteString(", 仓位警告=")
    buf.WriteString(d.Alarms.String())

    buf.WriteString(", 充电机温度=")
    buf.WriteString(strconv.Itoa(int(d.ChargeTemp)))

    buf.WriteString(", 充电电流(mA)=")
    buf.WriteString(strconv.Itoa(int(d.Chargei)))

    buf.WriteString(", 电池电量(mAH)=")
    buf.WriteString(strconv.Itoa(int(d.SOP)))

    buf.WriteString(", 剩余容量(%)=")
    buf.WriteString(strconv.Itoa(int(d.SOC)))

    buf.WriteString(", 电池健康状态(%)=")
    buf.WriteString(strconv.Itoa(int(d.SOH)))

    buf.WriteString(", 充电次数=")
    buf.WriteString(strconv.Itoa(int(d.ChargeTimes)))

    buf.WriteString(", 放电次数=")
    buf.WriteString(strconv.Itoa(int(d.DischargeTimes)))

    buf.WriteString(", 电池温度采样点温度(℃)=")
    for i, t := range d.BatteryTemp {
        if i > 0 {
            buf.WriteString("|")
        }
        buf.WriteString(strconv.Itoa(t))
    }

    buf.WriteString(", 电压(mV)=")
    buf.WriteString(strconv.Itoa(int(d.Voltage)))

    buf.WriteString(", 电池各串电压(mV)=")
    for i, t := range d.MonVoltage {
        if i > 0 {
            buf.WriteString("|")
        }
        buf.WriteString(strconv.Itoa(int(t)))
    }

    buf.WriteString(", 电池各串均衡状态=")
    for i, t := range d.Balance {
        if i > 0 {
            buf.WriteString("|")
        }
        buf.WriteString(strconv.Itoa(t))
    }

    buf.WriteString(", 充满电剩余时长(m)=")
    buf.WriteString(strconv.Itoa(int(d.Remaining)))

    buf.WriteString(", 电池状态=")
    switch d.BatteryStatus {
    default:
        buf.WriteString("-")
    case 1:
        buf.WriteString("充电")
    case 2:
        buf.WriteString("放电")
    case 3:
        buf.WriteString("静止")
    }

    buf.WriteString(", Mos状态=")
    buf.WriteString(d.MosStatus.String())

    buf.WriteString(", 充电机状态=")
    switch d.ChargerStatus {
    default:
        buf.WriteString("-")
    case 0:
        buf.WriteString("关机")
    case 1:
        buf.WriteString("开机")
    case 2:
        buf.WriteString("故障")
    }

    buf.WriteString(", 当前设置的电流大小(mA)=")
    buf.WriteString(strconv.Itoa(int(d.CurrentSet)))

    buf.WriteString(", 微动开关状态=")
    switch d.MicroSwitchStatus {
    default:
        buf.WriteString("-")
    case 0:
        buf.WriteString("未触发")
    case 1:
        buf.WriteString("已触发")
    case 2:
        buf.WriteString("故障")
    }

    buf.WriteString(", 电池仓温度(℃)=")
    buf.WriteString(strconv.Itoa(int(d.AmbientTemp)))

    buf.WriteString(", 风扇信号反馈=")
    switch d.FanStatus {
    default:
        buf.WriteString("-")
    case 0:
        buf.WriteString("未启动")
    case 1:
        buf.WriteString("已启动")
    }

    buf.WriteString(", 气溶胶状态=")
    switch d.AerosolsStatus {
    default:
        buf.WriteString("-")
    case 0:
        buf.WriteString("未触发")
    case 1:
        buf.WriteString("已触发")
    }

    return buf.String()
}

var emptyBatterySN = bytes.Repeat([]byte{48}, 16)

func (d *BinData) SetField(b []byte, name string) {
    // bytes.Repeat()
    switch name {
    case "Index":
        d.Index = b[0]
    case "Status":
        d.Status = BinStatus(b[0])
    case "Mode":
        d.Mode = b[0]
    case "BatterySN":
        b = bytes.TrimRightFunc(b, func(r rune) bool {
            return r < 48
        })
        if !bytes.Equal(b, emptyBatterySN) {
            d.BatterySN = adapter.ConvertBytes2String(b)
        }
    case "Alarms":
        d.Alarms = adapter.GetTrueBits(binary.BigEndian.Uint32(b), 15)
    case "ChargeTemp":
        d.ChargeTemp = b[0]
    case "Chargei":
        d.Chargei = binary.BigEndian.Uint32(b)
    case "SOP":
        d.SOP = binary.BigEndian.Uint16(b)
    case "SOC":
        d.SOC = b[0]
    case "SOH":
        d.SOH = b[0]
    case "ChargeTimes":
        d.ChargeTimes = binary.BigEndian.Uint16(b)
    case "DischargeTimes":
        d.DischargeTimes = binary.BigEndian.Uint16(b)
    case "BatteryTemp":
        d.BatteryTemp = make([]int, 3)
        for i := 0; i < 2; i++ {
            d.BatteryTemp[i] = int(b[i])
        }
    case "Voltage":
        d.Voltage = binary.BigEndian.Uint32(b)
    case "MonVoltage":
        d.MonVoltage = make([]uint16, 13)
        for i := 0; i < 13; i++ {
            d.MonVoltage[i] = binary.BigEndian.Uint16(b[i*2 : i*2+2])
        }
    case "Balance":
        d.Balance = adapter.GetTrueBits(binary.BigEndian.Uint32(b), 15)
    case "Remaining":
        d.Remaining = binary.BigEndian.Uint16(b)
    case "BatteryStatus":
        d.BatteryStatus = b[0]
    case "MosStatus":
        for i := 0; i < 2; i++ {
            d.MosStatus[i] = int(b[0] >> i & 1)
        }
    case "ChargerStatus":
        d.ChargerStatus = b[0]
    case "CurrentSet":
        d.CurrentSet = binary.BigEndian.Uint32(b)
    case "MicroSwitchStatus":
        d.MicroSwitchStatus = b[0]
    case "AmbientTemp":
        d.AmbientTemp = b[0]
    case "FanStatus":
        d.FanStatus = b[0]
    case "AerosolsStatus":
        d.AerosolsStatus = b[0]
    }
}

func (p *Parser) Bins(b []byte) (bs Bins) {
    bl := 129
    loops := len(b) / bl
    bs = make(Bins, loops)
    for i := 0; i < loops; i++ {
        bd := new(BinData)
        Parse(bd, b[i*bl:(i+1)*bl], binConfig)
        bs[i] = bd
    }
    return
}
