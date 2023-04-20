// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-15
// Based on yundong by liasica, magicrolan@qq.com.

package yundong

const (
	CodePassthrough       byte = 31 // 透传命令字, 该命令字用于接入服务器对电池柜的下行命令，电池柜通过响应消息满足服务器的服务请求和数据上报
	CodeAlarm             byte = 5  // 告警, 服务器进行回应
	CodeHeartbeat         byte = 39 // 心跳维持
	CodeUpgradeNotify     byte = 42 // http升级结果通知, 电池柜APK升级完成后通过该指令通知接入服务器升级结果
	CodeLogin             byte = 59 // 登录, 电池柜上电时通过该指令登录到接入服务器
	CodePeriodMsg         byte = 60 // 电池柜周期上报数据
	CodeDoorEvent         byte = 61 // 电池柜事件状态上报
	CodeStartMaintaining  byte = 62 // 电池柜开始维护, 维护工程师进入电池柜维护后台
	CodeStopMaintaining   byte = 63 // 电池柜结束维护, 维护工程师退出电池柜维护后台
	CodeRequestExchange   byte = 64 // 电池柜预约码请求换电, 预约换电请求
	CodeGenerateOrder     byte = 65 // 生成消费订单, 用户放入待换电电池后生产订单
	CodePayOrder          byte = 66 // 支付消费订单, 支付用户前一次换电消费订单
	CodeExchangeEnd       byte = 67 // 完成换电
	CodeExchangeCancel    byte = 68 // 取消换电, 由于异常原因取消换电
	CodeExchangeException byte = 69 // 电池柜上报换电异, 换电过程中上报换电异常
	CodeGetQrcode         byte = 70 // 获取换电二维码, 每次换电后电池柜重新获取换电二维码

	CodeGetAllowReserCount byte = 72  // 获取可换电池及预约电池数
	CodePeriodMsgV2        byte = 73  // 电池柜周期上报数据
	CodePeriodMsgBlankBox  byte = 74  // 电池柜周期上报数据
	CodeUnknown            byte = 75  // 未知的指令???
	CodePeriodMsgV3        byte = 90  // 电池柜周期上报数据
	CodePeriodMsgV4        byte = 92  // 电池柜周期上报数据
	CodeRequestExchangeV2  byte = 93  // 电池柜预约码请求换电
	CodeGenerateOrderV2    byte = 94  // 生成消费订单
	CodeDoorEventV2        byte = 95  // 电池柜事件状态上报
	CodeBinDisable         byte = 101 // 电柜仓门禁用事件上报
)
