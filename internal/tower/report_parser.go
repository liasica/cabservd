// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-07
// Based on cabservd by liasica, magicrolan@qq.com.

package tower

import (
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/cabinet"
    "github.com/auroraride/cabservd/internal/mem"
    "github.com/liasica/go-helpers/silk"
    "github.com/liasica/go-helpers/tools"
    "strconv"
)

// var cacheM = make(map[string]string)
//
// func saveCacheM() {
//     buf := adapter.NewBuffer()
//     defer adapter.ReleaseBuffer(buf)
//
//     for k := range cacheM {
//         buf.WriteString(k)
//         buf.WriteByte('\t')
//         buf.WriteString(cacheM[k])
//         buf.WriteByte('\n')
//     }
//
//     w, _ := os.OpenFile("./runtime/cache.log", os.O_CREATE|os.O_WRONLY, fs.ModePerm)
//
//     fmt.Println(buf.WriteTo(w))
// }

// GetSerial 获取电柜编码
func (r *Request) GetSerial() (string, bool) {
    return r.DevID, true
}

// GetCabinet 获取电柜信息
func (r *Request) GetCabinet() (cab *ent.CabinetPointer, exists bool) {
    // defer saveCacheM()

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
        // cacheM[attr.ID.String()] = v

        if _, ok := cabinetSignalMap[attr.ID]; ok {
            exists = true
        }

        if f, ok := cabinetSignals[attr.ID]; ok {
            f(cab, attr, v)
        } else {
            switch attr.ID {
            case SignalCabinetStatus:
                m := map[string]cabinet.Status{
                    "0": cabinet.StatusNormal,
                    "1": cabinet.StatusNormal,
                    "2": cabinet.StatusAbnormal,
                    "3": cabinet.StatusAbnormal,
                    "4": cabinet.StatusAbnormal,
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
    }
    return
}

// GetBins 获取仓位列表信息
func (r *Request) GetBins() (items ent.BinPointers) {
    // defer saveCacheM()

    m := make(map[int]*ent.BinPointer)

    for _, attr := range r.AttrList {
        // 原始字符串值
        v := attr.ValueString()
        // cacheM[attr.ID.String()] = v

        // 获取仓位序号
        ordinal, exists := attr.GetOrdinal()
        // 如果没有仓门信息, 直接跳过
        if !exists {
            continue
        }

        // 查询是否存在仓位信息
        b, ok := m[ordinal]
        if !ok {
            b = &ent.BinPointer{
                Serial:  silk.String(r.DevID),
                Ordinal: silk.Int(ordinal),
                Name:    silk.String(strconv.Itoa(ordinal) + "号仓"),
            }
            m[ordinal] = b
        }

        if f, extra := binSignals[attr.ID]; extra {
            f(b, attr, v)
        } else if attr.ID.Contains(SignalBatteryMonVoltage) {
            // 更新单芯电压
            // TODO 拓邦更新后删除
            mv := tools.StrToFloat64(v) / 1000.0
            index, _ := strconv.Atoi(attr.ID.String()[6:])
            mem.VoltageMonUpdate(r.DevID, ordinal, index, mv)
        } else {
            switch attr.ID {
            case SignalBinStatus:
                b.Health = silk.Bool(v != "5")
            case SignalBinDoorStatus:
                b.Open = silk.Bool(v == "1")
            case SignalBinEnable:
                b.Enable = silk.Bool(v == "1")
            case SignalBatterySN:
                b.BatterySn = silk.String(v)
                // 删除单芯电压
                // TODO 拓邦更新后删除
                if v == "" {
                    mem.VoltageClear(r.DevID, ordinal)
                }
            case SignalBatteryVoltage:
                vf := tools.StrToFloat64(v) / 100.0
                b.Voltage = silk.Float64(vf)
            case SignalBatteryCurrent:
                vf := tools.StrToFloat64(v) / 100.0
                b.Current = silk.Float64(vf)
            case SignalSOC:
                vf := tools.StrToFloat64(v)
                b.Soc = silk.Float64(vf)
            case SignalSOH:
                vf := tools.StrToFloat64(v)
                b.Soh = silk.Float64(vf)
            }
        }
    }

    for _, p := range m {
        p.Voltage = silk.Pointer(mem.VoltageGet(r.DevID, *p.Ordinal))
        items = append(items, p)
    }
    return
}
