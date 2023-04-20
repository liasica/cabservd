// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-15
// Based on yundong by liasica, magicrolan@qq.com.

package yundong

type Alarm byte

const (
	AlarmPoweroff              Alarm = 6  // 断电
	AlarmUndervoltage          Alarm = 26 // 总电压欠压
	AlarmBatteryUndervoltage   Alarm = 27 // 单节电池欠压
	AlarmOvervoltage           Alarm = 28 // 总电压过压
	AlarmBatteryOvervoltage    Alarm = 29 // 单节电池过压
	AlarmBatteryOverload       Alarm = 30 // 电池过载
	AlarmPos1HighTemperature   Alarm = 31 // 位置1高温
	AlarmPos2HighTemperature   Alarm = 32 // 位置2高温
	AlarmPos3HighTemperature   Alarm = 33 // 位置3高温
	AlarmMosHighTemperature    Alarm = 34 // 充电MOS高温
	AlarmChargingShortCircuit  Alarm = 35 // 充电短路
	AlarmDischargeShortCircuit Alarm = 36 // 放电短路
	AlarmChargingOvercurrent   Alarm = 37 // 充电过流
	AlarmDischargeOvercurrent  Alarm = 38 // 放电过流
)

func (d Alarm) String() string {
	switch d {
	default:
		return " - "
	case AlarmPoweroff:
		return "断电"
	case AlarmUndervoltage:
		return "总电压欠压"
	case AlarmBatteryUndervoltage:
		return "单节电池欠压"
	case AlarmOvervoltage:
		return "总电压过压"
	case AlarmBatteryOvervoltage:
		return "单节电池过压"
	case AlarmBatteryOverload:
		return "电池过载"
	case AlarmPos1HighTemperature:
		return "位置1高温"
	case AlarmPos2HighTemperature:
		return "位置2高温"
	case AlarmPos3HighTemperature:
		return "位置3高温"
	case AlarmMosHighTemperature:
		return "充电MOS高温"
	case AlarmChargingShortCircuit:
		return "充电短路"
	case AlarmDischargeShortCircuit:
		return "放电短路"
	case AlarmChargingOvercurrent:
		return "充电过流"
	case AlarmDischargeOvercurrent:
		return "放电过流"
	}
}

func (p *Parser) Alarm(b byte) string {
	// TODO Save
	data := Alarm(b)
	return "警告=" + data.String()
}
