// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-05
// Based on cabservd by liasica, magicrolan@qq.com.

package kaixin

import (
	"fmt"
	"strconv"
)

type Signal string

// SignalData 信号量结构体
type SignalData struct {
	ID    Signal `json:"id,omitempty"`    // 信号量ID
	Value any    `json:"value,omitempty"` // 参数值
}

func (s *SignalData) StringValue() (v string) {
	v = fmt.Sprintf("%v", s.Value)
	if v == "null" {
		v = ""
	}
	return
}

func (s *SignalData) ValueFloat64() (v float64) {
	v, _ = strconv.ParseFloat(s.StringValue(), 64)
	return
}

func (s *SignalData) ValueInt64() (v int64) {
	v, _ = strconv.ParseInt(s.StringValue(), 10, 64)
	return
}

const (
	DoorStatusClose = "0" // 仓门 - 关
	DoorStatusOpen  = "1" // 仓门 - 开

	BinEnable  = "1" // 仓位 - 启用
	BinDisable = "0" // 仓位 - 禁用
)

const (
	SignalLng                Signal = "02111001" // 柜子经度
	SignalLat                Signal = "02112001" // 柜子纬度
	SignalBinStatus          Signal = "02102001" // 仓位状态 0:上电初始化 1:无换电、放电、取电动作 2:换电中 3:在归还电池中 4:在取出电池中 5:换电柜异常
	SignalBinChargeStatus    Signal = "02104001" // 仓位充电状态 0:无电池 1:电池正在充电 2:电池充满 5:异常
	SignalDoorStatus         Signal = "02103001" // 仓位柜门状态 0:关 1:开
	SignalBinID              Signal = "01309001" // 柜门ID 1~16:对应柜门ID
	SignalBatteryID          Signal = "01310001" // 换电柜设备ID
	SignalGSM                Signal = "02105001" // GSM 信号强度
	SignalBatterySN          Signal = "02106001" // 柜内电池SN
	SignalBatteryVoltage     Signal = "01111001" // 电池总电压 (V)
	SignalBatteryCurrent     Signal = "01112001" // 电池总电流 (A)
	SignalCabinetVoltage     Signal = "02107001" // 换电柜总电压 (V)
	SignalCabinetCurrent     Signal = "02108001" // 换电柜总电流 (A)
	SignalSOC                Signal = "02109001" // 电池电量
	SignalSOH                Signal = "02110001" // 电池健康
	SignalBatteryCellID      Signal = "01116001" // XX单芯电压串数 (表示电池内电芯总数量值)
	SignalBatteryCellVoltage Signal = "01117001" // XX单芯电压值 (XX如果为01，对应的信号量ID结尾为01)
	SignalCabinetTemp        Signal = "02113001" // 柜体温度值 (换电柜温度)
	SignalPCBTemp            Signal = "01118001" // 功率温度值 (电池内部PCB板表面温度)
	SignalBatteryCellTemp    Signal = "01119001" // 电芯温度值 (电池内部多组电芯中间表面温度)
	SignalBatteryAmbientTemp Signal = "01120001" // 环境温度 (电池壳体内部整体温度)
	SignalBatteryStatus      Signal = "01121001" // 电池状态 0:移动 1:静止 2:存储 3:休眠
	SignalBatteryControl     Signal = "01122001" // 电池控制
	SignalDischarge          Signal = "02116001" // 总放电 (Ah)
	SignalCharge             Signal = "02117001" // 总充电 (Ah)
	SignalBatteryChargeTime  Signal = "02114001" // 电池充电时长
	SignalBinEnable          Signal = "02118001" // 柜门是否禁用 (0:禁用 1:启用)
	SignalEnable             Signal = "02119001" // 柜体是否禁用 (0:禁用 1:启用)
	SignalEnergy             Signal = "02120001" // 柜子总用电量 (kwh)
	SignalCabinetControl     Signal = "02301001" // 控制换电柜命令
)

type ControlValue string

const (
	ControlCabinetDisable ControlValue = "00" // 设置换电柜不可用
	ControlExchange       ControlValue = "01" // 换电
	ControlPutIn          ControlValue = "02" // 放电
	ControlPutOut         ControlValue = "03" // 取电
	ControlOpenDoor       ControlValue = "04" // 开启柜门
	ControlBinDisable     ControlValue = "06" // 设置柜门不可用
	ControlBinEnable      ControlValue = "07" // 设置柜门可用
	ControlBatteryBind    ControlValue = "08" // 柜门绑定电池序列号
	ControlBatteryUnbind  ControlValue = "09" // 柜门解绑电池序列号
	ControlCabinetEnable  ControlValue = "10" // 设置换电柜可用
	ControlBatteryRent    ControlValue = "11" // 租用电池(首放)
	ControlBatteryTenancy ControlValue = "12" // 退还电池
)
