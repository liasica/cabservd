// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-15
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

import (
	"strconv"

	"github.com/liasica/go-helpers/silk"

	"github.com/auroraride/cabservd/internal/ent"
	"github.com/auroraride/cabservd/internal/ent/cabinet"
)

type CabAttrs []*CabAttr

// CabAttr 柜机完整属性 (包含业务属性和物理属性)
type CabAttr struct {
	Name         *string  `json:"name,omitempty"`
	Sn           *string  `json:"sn,omitempty"`
	Online       *int     `json:"online,omitempty"`       // 柜机是否在线 0:在线 1:不在线
	OnlineTime   *int64   `json:"onlineTime,omitempty"`   // 柜机最近一次的上线
	ModelType    *string  `json:"modelType,omitempty"`    // 型号类型名称
	CellNums     *int     `json:"cellNums,omitempty"`     // 格子数量
	CustomPhone  *string  `json:"customPhone,omitempty"`  // 客服电话（用于柜机显示）
	CellAttrList BinAttrs `json:"cellAttrList,omitempty"` // 仓位信息
	// 业务属性↑↑↑ 物理属性↓↓↓
	Version *string  `json:"version,omitempty"` // 柜机app的版本号
	Voltage *float64 `json:"voltage,omitempty"` // 柜机电压
	Current *float64 `json:"current,omitempty"` // 柜机电流
	Temp    *float64 `json:"temp,omitempty"`    // 柜机温度
	Iccid   *string  `json:"iccid,omitempty"`   // 柜机sim卡号
	// 需要更新柜子属性↑↑↑
	PowerConsumption        *float64 `json:"powerConsumption,omitempty"`        // 耗电量（从柜机第一次插电开始使用到目前的电量）
	FanStatus               *int     `json:"fanStatus,omitempty"`               // 核心风扇状态，0:关闭，1:开启
	LightStatus             *int     `json:"lightStatus,omitempty"`             // 核心灯的状态，0:关闭，1:开启
	Power                   *float64 `json:"power,omitempty"`                   // 功率
	PowerFactor             *float64 `json:"powerFactor,omitempty"`             // 功率因数
	ActiveElectricityEnergy *float64 `json:"activeElectricityEnergy,omitempty"` // 有功电能
	WaterPumpStatus         *int     `json:"waterPumpStatus,omitempty"`         // 水泵开关 0:关闭 1:开启
	WaterLevelWarningStatus *int     `json:"waterLevelWarningStatus,omitempty"` // 水箱水位告警 0:关闭 1:打开
	WaterLeachingWarning    *int     `json:"waterLeachingWarning,omitempty"`    // 柜内水浸状态 0:关闭 1:打开
	Humidity                *float64 `json:"humidity,omitempty"`                // 湿度
	DoorStatus              *int     `json:"doorStatus,omitempty"`              // 柜机运维门是否打开 0:关闭 1:打开
	SmokeSensorStatus       *int     `json:"smokeSensorStatus,omitempty"`       // 烟感传感器是否打开 0:关闭， 1:打开
	ExtinguisherStatus      *int     `json:"extinguisherStatus,omitempty"`      // 灭火器是否打开 0:关闭 1:打开
	CVersion                *string  `json:"cVersion,omitempty"`                // 核心硬件控制版的版本号
}

func (attr *CabAttr) GetSerial() (string, bool) {
	if attr.Sn == nil || *attr.Sn == "" {
		return "", false
	}
	return *attr.Sn, true
}

func (attr *CabAttr) GetCabinet() (cab *ent.CabinetPointer, exist bool) {
	if attr.Sn == nil {
		return nil, false
	}

	exist = attr.Online != nil || attr.Voltage != nil || attr.Current != nil || attr.Iccid != nil || attr.Temp != nil || attr.PowerConsumption != nil
	if !exist {
		return
	}

	cab = &ent.CabinetPointer{
		Serial: silk.String(*attr.Sn),
		Status: silk.Pointer(cabinet.StatusNormal),
		Enable: silk.Bool(true),
	}

	// 是否在线
	if attr.Online != nil {
		cab.Online = silk.Bool(*attr.Online == 0)
	}

	// 电压
	if attr.Voltage != nil {
		cab.Voltage = attr.Voltage
	}

	// 电流
	if attr.Current != nil {
		cab.Current = attr.Current
	}

	// 温度
	if attr.Temp != nil {
		cab.Temperature = attr.Temp
	}

	// SIM卡
	if attr.Iccid != nil {
		cab.Sim = attr.Iccid
	}

	// 耗电量
	if attr.PowerConsumption != nil {
		cab.Electricity = attr.PowerConsumption
	}

	return cab, true
}

func (attr *CabAttr) GetBins() (items ent.BinPointers) {
	if attr.Sn == nil {
		return
	}

	items = make(ent.BinPointers, len(attr.CellAttrList))
	for i, c := range attr.CellAttrList {
		// 获取健康状态
		var health *bool
		if c.ChargeStatus != nil {
			cs := *c.ChargeStatus
			health = silk.Bool(cs <= 4)
		}

		var (
			v   *float64
			a   *float64
			soc *float64
		)

		if c.Soc != nil {
			soc = silk.Float64(float64(*c.Soc))
		}

		// 充电器电压和电流对比电池会有0.1-0.2左右的区别, 所以仓位电流和电压可以随意使用二者值
		if c.ChargeV != nil {
			v = c.ChargeV
		}
		if c.ChargeA != nil {
			a = c.ChargeA
		}

		if c.BatteryV != nil {
			v = c.BatteryV
		}
		if c.BatteryA != nil {
			a = c.BatteryA
		}

		var batExists *bool

		// silk.Bool((c.BatterySn != nil && *c.BatterySn != "") || (c.ExistsBattery != nil && *c.ExistsBattery == 1))
		if c.BatterySn != nil {
			batExists = silk.Bool(*c.BatterySn != "")
		}

		// 暂时不使用ExistsBattery
		// if c.ExistsBattery != nil {
		// 	batExists = silk.Bool(*c.ExistsBattery == 1)
		// }

		p := &ent.BinPointer{
			Serial:        attr.Sn,
			Name:          silk.String(strconv.Itoa(*c.CellNo) + "号仓"),
			Ordinal:       c.CellNo,
			Open:          silk.PointerConditionBool(c.DoorStatus, 1),
			Enable:        silk.PointerConditionBool(c.ForbidStatus, 0),
			Health:        health,
			BatteryExists: batExists,
			BatterySn:     c.BatterySn,
			Voltage:       v,
			Current:       a,
			Soc:           soc,
			Remark:        c.ForbidReason,
		}

		// 清除备注
		if p.Enable != nil && *p.Enable {
			p.Remark = silk.String("")
		}

		items[i] = p
	}
	return items
}

type AppAttrs []*AppAttr

// AppAttr 柜机APP设置属性
type AppAttr struct {
	Sn                      *string            `json:"sn,omitempty"`
	SystemBarStatus         *int               `json:"systemBarStatus,omitempty"`         // 系统导航栏0:关闭 1:开启
	LockStatus              *int               `json:"lockStatus,omitempty"`              // 0:正常 1:反锁，此值不需要关心，西六楼会在出厂柜机前设置好。
	ChargingNum             *int               `json:"chargingNum,omitempty"`             // 可以同时充电的个数
	NormalObj               *AppAttrNormalObj  `json:"normalObj,omitempty"`               // 单型号的相关属性
	MultiVObjs              []AppAttrMultiVObj `json:"multiVObjs,omitempty"`              // 多型号相关属性
	Mode                    *string            `json:"mode,omitempty"`                    // 模式：单型号:NORMAL 多型号:MULTV，注意如果把柜机型号设成MULTV，那么充电策略就会变成多型号的充电策略,并且离线换电也会按照多型号进行。
	ShowRealPower           *int               `json:"showRealPower,omitempty"`           // 是否显示电池的真实电量，0:关闭，1:显示，默认关闭，在多型号模式下，电量只有0和100，当达到可以换电的电压时，系统就会把电池电量设置为100，没有达到设置0.如果开启显示，系统将不会对电量做特殊处理，会显示电池的真实电量。此参数只有在多型号模式下有效
	QrAddress               *string            `json:"qrAddress,omitempty"`               // 二维码内容
	OnlineMode              *int               `json:"onlineMode,omitempty"`              // 离线模式是否开启 0:关闭 1:开启
	LogStatus               *int               `json:"logStatus,omitempty"`               // 底层日志是否开启
	AutoTemp                *int               `json:"autoTemp,omitempty"`                // 自动温控 0:关闭 1:开启,开启后，温度和风扇参数设置才有效
	HeatCondition           *int               `json:"heatCondition,omitempty"`           // 低于多少度开启加热
	FanCondition            *int               `json:"fanCondition,omitempty"`            // 高于多少度开启风扇(包括核心版的风扇)
	ExistsBatteryLockStatus *int               `json:"existsBatteryLockStatus,omitempty"` // 电池异常消失是否需要锁仓 0:关闭 1:开启
	BmsHealthLockStatus     *int               `json:"bmsHealthLockStatus,omitempty"`     // BMS出现异常是否需要自动锁仓 0:关闭 1:开启
}

type BinAttrs []*BinAttr

// BinAttr 仓位完整属性 (包含格挡信息和充电器信息)
type BinAttr struct {
	CellNo               *int     `json:"cellNo,omitempty"`               // 格挡的编号，默认从1开始。
	Temp                 *float64 `json:"temp,omitempty"`                 // 温度
	DoorStatus           *int     `json:"doorStatus,omitempty"`           // 门的状态 0:关闭 1:打开
	IndicatorLightStatus *int     `json:"indicatorLightStatus,omitempty"` // 指示灯的状态 0:关闭 1:红色 2:绿色 3:黄色
	CellLightStatus      *int     `json:"cellLightStatus,omitempty"`      // 仓内灯的状态 0:关闭 1:打开
	FanStatus            *int     `json:"fanStatus,omitempty"`            // 风扇的状态 0:关闭 1:打开
	HeatStatus           *int     `json:"heatStatus,omitempty"`           // 加热的状态 0:关闭 1:打开
	ChargeStatus         *int     `json:"chargeStatus,omitempty"`         // 充电器的状态：0:关闭 1:开机中 2:充电中 3:充满电 4:限制充电 -128:过压充电 64:过流充电 32:短路 16:温度过高. 10:超压 11:电池反接 12:NTC故障停机 13:输出短路停机
	ChargeV              *float64 `json:"chargeV,omitempty"`              // 充电器的电压
	ChargeA              *float64 `json:"chargeA,omitempty"`              // 充电器的电流
	Version              *string  `json:"version,omitempty"`              // 硬件版本号
	BatterySn            *string  `json:"batterySn,omitempty"`            // 电池编号，如果仓内有电池，那么该字段就会有值
	ForbidStatus         *int     `json:"forbidStatus,omitempty"`         // 格挡的禁用状态 0:解锁 1:加锁
	ForbidReason         *string  `json:"forbidReason,omitempty"`         // 禁用原因
	ForbidType           *int     `json:"forbidType,omitempty"`           // 禁用的类型 0:系统禁用 1:人为禁用
	AerosolStatus        *int     `json:"aerosolStatus,omitempty"`        // 气溶胶的打开状态 0:关闭 1:加锁
	ExistsBattery        *int     `json:"existsBattery,omitempty"`        // 0:不存在 ,1:存在 是否真实存在电池,（保留）
	BatteryAttr
}

type BatteryAttrs []*BatteryAttr

// BatteryAttr 电池属性
type BatteryAttr struct {
	Soc           *int      `json:"soc,omitempty"`           // 电池的电量，0-100
	BatteryA      *float64  `json:"batteryA,omitempty"`      // 电池的电流
	BatteryV      *float64  `json:"batteryV,omitempty"`      // 电池的电压
	CoreNum       *int      `json:"coreNum,omitempty"`       // 电芯数量
	CoreV         []float64 `json:"coreV,omitempty"`         // 每个电芯的电压，根据电芯数量的不同，列表的数量也不同，默认从电芯1开始。
	BatteryHealth []string  `json:"batteryHealth,omitempty"` // 电池的健康状态信息，详见附录六
	Capacity      *float64  `json:"capacity,omitempty"`      // 电池容量
	EnvTemp       *float64  `json:"envTemp,omitempty"`       // 电池环境温度
	CoreTemp      *float64  `json:"coreTemp,omitempty"`      // 电池电芯温度
	BoardTemp     *float64  `json:"boardTemp,omitempty"`     // 电池板卡温度
	ModelType     *string   `json:"modelType,omitempty"`     // 电池型号名称，如果是多型号，就会返回
	Longitude     *string   `json:"longitude,omitempty"`     // 经度
	Latitude      *string   `json:"latitude,omitempty"`      // 纬度
}
