// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-18
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

type CellOpenIndicatorLightData struct {
	Color *int `json:"color"` // 颜色 1:红色 2:绿色 3:黄色
}

type CellOpenChargeData struct {
	V *float64 `json:"v,omitempty"` // 开机电压
	A *int     `json:"a,omitempty"` // 开机电流
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

type BatterySnData struct {
	BatterySn string `json:"batterySn"`
}
