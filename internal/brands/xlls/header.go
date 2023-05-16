// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-16
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

const allowMethod = "POST" // POST

const (
	pathHardwareOperation = "/hop" // 硬件操作结果通知
	pathBusinesss         = "/biz" // 业务结果通知
	pathOfflineExchange   = "/oex" // 离线换电结果通知
	pathCellChange        = "/bin" // 格挡状态变化通知
	pathBatteryChange     = "/bat" // 电池盒充电器状态变化通知
	pathCabinetChange     = "/cab" // 柜机状态变化通知
	pathHardwareFault     = "/hfl" // 硬件故障通知
	pathSelfServiceOpen   = "/sso" // 自助开仓回调通知
)

var (
	headerXRealIP       = []byte("X-Real-IP")
	headerXForwardedFor = []byte("X-Forwarded-For")
)
