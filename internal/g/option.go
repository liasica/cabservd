// Copyright (C) liasica. 2023-present.
//
// Created at 2023-03-18
// Based on cabservd by liasica, magicrolan@qq.com.

package g

var (
	BatteryReign              bool // 是否有电池在位检测
	AutoResetWithoutBatterySN bool // 当电池编号不存在的时候是否自动清除电池信息 (soc / soh / current / voltage)
	CalculateMonVoltage       bool // 是否需要计算当前电池电芯单体电压作为仓位电压 (拓邦电柜不上报电柜电压)
)
