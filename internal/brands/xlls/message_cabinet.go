// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-15
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

type BusinessAttrs []*BusinessAttr

// BusinessAttr 柜机业务属性
type BusinessAttr struct {
	Name        *string `json:"name,omitempty"`
	Sn          *string `json:"sn,omitempty"`
	Online      *int    `json:"online,omitempty"`      // 柜机是否在线 0--在线 1--不在线
	OnlineTime  *int64  `json:"onlineTime,omitempty"`  // 柜机最近一次的上线
	ModelType   *string `json:"modelType,omitempty"`   // 型号类型名称
	CellNums    *int    `json:"cellNums,omitempty"`    // 格子数量
	CustomPhone *string `json:"customPhone,omitempty"` // 客服电话（用于柜机显示）
}

type PhysicsAttrs []*PhysicsAttr

// PhysicsAttr 柜机物理属性
type PhysicsAttr struct {
	Sn                      *string  `json:"sn,omitempty"`
	Version                 *string  `json:"version,omitempty"`                 // 柜机app的版本号
	Voltage                 *float64 `json:"voltage,omitempty"`                 // 柜机电压
	Current                 *float64 `json:"current,omitempty"`                 // 柜机电流
	Temp                    *float64 `json:"temp,omitempty"`                    // 柜机温度
	Iccid                   *string  `json:"iccid,omitempty"`                   // 柜机sim卡号
	PowerConsumption        *float64 `json:"powerConsumption,omitempty"`        // 耗电量（从柜机第一次插电开始使用到目前的电量）
	FanStatus               *int     `json:"fanStatus,omitempty"`               // 核心风扇状态，0--关闭，1--开启
	LightStatus             *int     `json:"lightStatus,omitempty"`             // 核心灯的状态，0--关闭，1--开启
	Power                   *float64 `json:"power,omitempty"`                   // 功率
	PowerFactor             *float64 `json:"powerFactor,omitempty"`             // 功率因数
	ActiveElectricityEnergy *float64 `json:"activeElectricityEnergy,omitempty"` // 有功电能
	WaterPumpStatus         *int     `json:"waterPumpStatus,omitempty"`         // 水泵开关 0--关闭 1--开启
	WaterLevelWarningStatus *int     `json:"waterLevelWarningStatus,omitempty"` // 水箱水位告警 0--关闭 1--打开
	WaterLeachingWarning    *int     `json:"waterLeachingWarning,omitempty"`    // 柜内水浸状态 0--关闭 1--打开
	Humidity                *float64 `json:"humidity,omitempty"`                // 湿度
	DoorStatus              *int     `json:"doorStatus,omitempty"`              // 柜机运维门是否打开 0--关闭 1--打开
	SmokeSensorStatus       *int     `json:"smokeSensorStatus,omitempty"`       // 烟感传感器是否打开 0--关闭， 1--打开
	ExtinguisherStatus      *int     `json:"extinguisherStatus,omitempty"`      // 灭火器是否打开 0--关闭 1--打开
	CVersion                *string  `json:"cVersion,omitempty"`                // 核心硬件控制版的版本号
	ReportTime              *int64   `json:"reportTime,omitempty"`              // 柜机上报的时间
}

// AppAttrMultiVObj 多型号相关属性
type AppAttrMultiVObj struct {
	BatteryModel            *string  `json:"batteryModel,omitempty"`            // 电池型号简称
	ChargeV                 *float64 `json:"chargeV,omitempty"`                 // 充电电压
	ChargeA                 *float64 `json:"chargeA,omitempty"`                 // 充电电流
	ChargeVUpperLimit       *float64 `json:"chargeVUpperLimit,omitempty"`       // 充电的上限,按照电压设置
	ChargeVLowerLimit       *float64 `json:"chargeVLowerLimit,omitempty"`       // 充电的下限,按照电压设置
	ExchangePowerVCondition *float64 `json:"exchangePowerVCondition,omitempty"` // 可以进行换电的最低电压，不能高于充电的上限
}

// AppAttrNormalObj 单型号相关属性
type AppAttrNormalObj struct {
	ChargeV                *float64 `json:"chargeV,omitempty"`                // 充电电压
	ChargeA                *float64 `json:"chargeA,omitempty"`                // 充电电流
	ChargeUpperLimit       *int     `json:"chargeUpperLimit,omitempty"`       // 充电的上限，值必须是0-100之间，按照电量设置
	ChargeLowerLimit       *int     `json:"chargeLowerLimit,omitempty"`       // 充电的下限，值必须是0-100之间，按照电量设置
	ExchangePowerCondition *int     `json:"exchangePowerCondition,omitempty"` // 可以进行换电的最低电量，值必须是0-100之间并且不能高于充电的上限
}

type AppAttrs []*AppAttr

// AppAttr 柜机APP设置属性
type AppAttr struct {
	Sn                      *string            `json:"sn,omitempty"`
	SystemBarStatus         *int               `json:"systemBarStatus,omitempty"`         // 系统导航栏0--关闭 1--开启
	LockStatus              *int               `json:"lockStatus,omitempty"`              // 0--正常 1--反锁，此值不需要关心，西六楼会在出厂柜机前设置好。
	ChargingNum             *int               `json:"chargingNum,omitempty"`             // 可以同时充电的个数
	NormalObj               *AppAttrNormalObj  `json:"normalObj,omitempty"`               // 单型号的相关属性
	MultiVObjs              []AppAttrMultiVObj `json:"multiVObjs,omitempty"`              // 多型号相关属性
	Mode                    *string            `json:"mode,omitempty"`                    // 模式：单型号--NORMAL 多型号--MULTV，注意如果把柜机型号设成MULTV，那么充电策略就会变成多型号的充电策略,并且离线换电也会按照多型号进行。
	ShowRealPower           *int               `json:"showRealPower,omitempty"`           // 是否显示电池的真实电量，0--关闭，1--显示，默认关闭，在多型号模式下，电量只有0和100，当达到可以换电的电压时，系统就会把电池电量设置为100，没有达到设置0.如果开启显示，系统将不会对电量做特殊处理，会显示电池的真实电量。此参数只有在多型号模式下有效
	QrAddress               *string            `json:"qrAddress,omitempty"`               // 二维码内容
	OnlineMode              *int               `json:"onlineMode,omitempty"`              // 离线模式是否开启 0--关闭 1--开启
	LogStatus               *int               `json:"logStatus,omitempty"`               // 底层日志是否开启
	AutoTemp                *int               `json:"autoTemp,omitempty"`                // 自动温控 0--关闭 1--开启,开启后，温度和风扇参数设置才有效
	HeatCondition           *int               `json:"heatCondition,omitempty"`           // 低于多少度开启加热
	FanCondition            *int               `json:"fanCondition,omitempty"`            // 高于多少度开启风扇(包括核心版的风扇)
	ExistsBatteryLockStatus *int               `json:"existsBatteryLockStatus,omitempty"` // 电池异常消失是否需要锁仓 0--关闭 1--开启
	BmsHealthLockStatus     *int               `json:"bmsHealthLockStatus,omitempty"`     // BMS出现异常是否需要自动锁仓 0--关闭 1--开启
}

// CabinetMeta 柜机基本属性
type CabinetMeta struct {
	Sn          *string `json:"sn,omitempty"`
	Name        *string `json:"name,omitempty"`        // 柜机名称
	ModelType   *string `json:"modelType,omitempty"`   // 柜机型号
	CustomPhone *string `json:"customPhone,omitempty"` // 客服电话
}

// CabinetModel 柜机型号
type CabinetModel struct {
	CellNums  *int    `json:"cellNums,omitempty"`  // 格挡数量
	ModelType *string `json:"modelType,omitempty"` // 型号名称
}

// CabinetPassword 柜机密码
type CabinetPassword struct {
	PWD *string `json:"pwd"`
}
