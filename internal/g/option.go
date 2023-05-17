// Copyright (C) liasica. 2023-present.
//
// Created at 2023-03-18
// Based on cabservd by liasica, magicrolan@qq.com.

package g

// 获取空仓忽略电压和电流 (空仓的时候有可能会有一定的电压和电流)
var (
	Fakevoltage float64 = 40 // 默认为40V虚拟电压
	Fakecurrent float64 = -1 // 值为-1的时候默认不检测
)

var (
	UseHttp                     bool // 是否使用http协议控制
	BatteryReign                bool // 是否有电池在位检测
	AutoResetWithoutBatterySN   bool // 当电池编号不存在的时候是否自动清除电池信息 (soc / soh / current / voltage)
	CalculateMonVoltage         bool // 是否需要计算当前电池电芯单体电压作为仓位电压 (拓邦电柜不上报电柜电压)
	LogBinary                   bool // 是否记录原始日志
	ExchangeFirstStepRetryTimes = 1  // 换电第一步是否重复发送开仓指令, 默认为1:不重试
	ExchangeThirdStepRetryTimes = 1  // 换电第三步是否重复发送开仓指令, 默认为1:不重试
)
