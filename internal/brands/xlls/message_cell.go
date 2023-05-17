// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-16
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

type CellAttrs []*CellAttr

// CellAttr 格挡相关属性
type CellAttr struct {
	CellNo               *int     `json:"cellNo,omitempty"`               // 格挡的编号，默认从1开始。
	Temp                 *float64 `json:"temp,omitempty"`                 // 温度
	DoorStatus           *int     `json:"doorStatus,omitempty"`           // 门的状态 0--关闭 1--打开
	IndicatorLightStatus *int     `json:"indicatorLightStatus,omitempty"` // 指示灯的状态 0--关闭 1--红色 2--绿色 3--黄色
	CellLightStatus      *int     `json:"cellLightStatus,omitempty"`      // 仓内灯的状态 0--关闭 1--打开
	FanStatus            *int     `json:"fanStatus,omitempty"`            // 风扇的状态 0--关闭 1--打开
	HeatStatus           *int     `json:"heatStatus,omitempty"`           // 加热的状态 0--关闭 1--打开
	ChargeStatus         *int     `json:"chargeStatus,omitempty"`         // 充电器的状态：0--关闭 1--开机中 2--充电中 3--充满电 4--限制充电 -128--过压充电 64--过流充电 32--短路 16--温度过高. 10--超压 11--电池反接 12--NTC故障停机 13--输出短路停机
	ChargeV              *float64 `json:"chargeV,omitempty"`              // 充电器的电压
	ChargeA              *float64 `json:"chargeA,omitempty"`              // 充电器的电流
	Version              *string  `json:"version,omitempty"`              // 硬件版本号
	BatterySn            *string  `json:"batterySn,omitempty"`            // 电池编号，如果仓内有电池，那么该字段就会有值
	ForbidStatus         *int     `json:"forbidStatus,omitempty"`         // 格挡的禁用状态 0--解锁 1--加锁
	ForbidReason         *string  `json:"forbidReason,omitempty"`         // 禁用原因
	ForbidType           *int     `json:"forbidType,omitempty"`           // 禁用的类型 0--系统禁用 1--人为禁用
	AerosolStatus        *int     `json:"aerosolStatus,omitempty"`        // 气溶胶的打开状态 0--关闭 1--加锁
}

type CellOpenIndicatorLightData struct {
	Color *int `json:"color"` // 颜色 1--红色 2--绿色 3--黄色
}

type CellOpenChargeData struct {
	V *float64 `json:"v,omitempty"` // 开机电压
	A *int     `json:"a,omitempty"` // 开机电流
}
