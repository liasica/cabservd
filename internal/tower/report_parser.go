// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-07
// Based on cabservd by liasica, magicrolan@qq.com.

package tower

import (
	"strconv"

	"github.com/auroraride/adapter"
	"github.com/auroraride/cabservd/internal/ent"
	"github.com/auroraride/cabservd/internal/ent/cabinet"
	"github.com/auroraride/cabservd/internal/g"
	"github.com/auroraride/cabservd/internal/mem"
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
		if attr == nil {
			continue
		}

		v := attr.ValueString()

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
	m := make(map[int]*ent.BinPointer)

	for _, attr := range r.AttrList {
		if attr == nil {
			continue
		}

		// 原始字符串值
		v := attr.ValueString()

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

		// 如果信号量是特殊信号量
		if f, extra := binSignals[attr.ID]; extra {
			f(b, attr, v)
			continue
		}

		// 计算电芯单体电压
		if attr.ID.Contains(SignalBatteryMonVoltage) {
			// 更新单芯电压
			// TODO 拓邦更新后删除
			mv := tools.StrToFloat64(v) / 1000.0
			index, _ := strconv.Atoi(attr.ID.String()[6:])
			mem.VoltageMonUpdate(r.DevID, ordinal, index, mv)
			continue
		}

		switch attr.ID {
		case SignalBinStatus:
			b.Health = silk.Bool(v != "5")
		case SignalBinDoorStatus:
			b.Open = silk.Bool(v == "1")
		case SignalBinEnable:
			b.Enable = silk.Bool(v == "1")
		case SignalBatterySN:
			// 如果不需要和电柜通讯, 直接保存电池编码
			vaild := g.Config.NonBms || v == ""
			// 如果需要和电柜通讯, 需要解析电池编码
			if !vaild {
				// 解析电池编码
				// 接收到电池编码后, 尝试格式化电池编码, 若可以正常被格式化, 则是有效电池 ---- by: 曹博文, 2023-03-25 23:08
				bat, _ := adapter.ParseBatterySN(v)
				if bat.Brand != "" {
					vaild = true
				}
			}
			// // 如果不需要和电柜通讯, 直接保存电池编码
			// vaild := g.Config.NonBms
			// // 如果需要和电柜通讯, 需要解析电池编码
			// // TODO 优化 if else
			// if !vaild {
			// 	// 解析电池编码
			// 	if v != "" {
			// 		// 接收到电池编码后, 尝试格式化电池编码, 若可以正常被格式化, 则是有效电池 ---- by: 曹博文, 2023-03-25 23:08
			// 		bat, _ := adapter.ParseBatterySN(v)
			// 		if bat.Brand != "" {
			// 			vaild = true
			// 		}
			// 	} else {
			// 		vaild = true
			// 	}
			// }

			// 需要保存电池编码
			if vaild {
				b.BatterySn = silk.String(v)
			}

			// 删除单芯电压
			// TODO 拓邦更新后删除
			if v == "" && g.CalculateMonVoltage {
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

	for _, p := range m {
		if g.CalculateMonVoltage {
			p.Voltage = silk.Pointer(mem.VoltageGet(r.DevID, *p.Ordinal))
		}

		items = append(items, p)
	}
	return
}
