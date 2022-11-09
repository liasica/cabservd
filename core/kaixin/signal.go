// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-05
// Based on cabservd by liasica, magicrolan@qq.com.

package kaixin

type Signal string

// SignalData 信号量结构体
type SignalData struct {
    ID    Signal `json:"id,omitempty"`    // 信号量ID
    Value any    `json:"value,omitempty"` // 参数值
}

const (
    SignalLng                Signal = "02111001" // 柜子经度
    SignalLat                Signal = "02112001" // 柜子纬度
    SignalStatus             Signal = "02102001" // 换电柜状态 0:上电初始化 1:无换电、放电、取电动作 2:换电中 3:在归还电池中 4:在取出电池中 5:换电柜异常
    SignalDoorStatus         Signal = "02103001" // 柜门状态 0:关 1:开
    SignalBinStatus          Signal = "02104001" // 柜体状态 (仓位状态) 0:无电池 1:电池正在充电 2:电池充满 5:异常
    SignalBinID              Signal = "01309001" // 柜门ID 1~16:对应柜门ID
    SignalBatteryID          Signal = "01310001" // 换电柜设备ID
    SignalGSM                Signal = "02105001" // GSM 信号强度
    SignalBatterySN          Signal = "02106001" // 柜内电池SN
    SignalBatteryVoltage     Signal = "01111001" // 电池总电压 (V)
    SignalBatteryCurrent     Signal = "01112001" // 电池总电流 (A)
    SignalCabinetVoltage     Signal = "02107001" // 换电柜总电压 (V)
    SignalCabinetCurrent     Signal = "02108001" // 换电柜总电流 (A)
    SignalSOC                Signal = "02109001"
    SignalSOH                Signal = "02110001"
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
