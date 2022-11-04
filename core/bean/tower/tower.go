// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-03
// Based on cabservd by liasica, magicrolan@qq.com.

package tower

import "cabservd/pkg/silk"

// KXMessageType 消息类型
type KXMessageType int

const (
    KXMTLoginRequest    KXMessageType = 100 // 登录请求
    KXMTLoginResponse   KXMessageType = 101 // 登录响应
    KXMTReportRequest   KXMessageType = 300 // 属性上报请求
    KXMTReportResponse  KXMessageType = 301 // 属性上报响应
    KXMTNoticeRequest   KXMessageType = 400 // 告警上报
    KXMTNoticeResponse  KXMessageType = 401 // 告警上报响应
    KXMTControlResponse KXMessageType = 500 // 电柜控制请求
    KXMTControlRequest  KXMessageType = 501 // 电柜控制响应
)

type KXMessage struct {
    MessageType KXMessageType `json:"messageType"` // 消息类型
    TxnNo       int64         `json:"txnNo"`       // 流水号: 13位时间戳
    DevID       string        `json:"devId"`       // 设备ID
}

// KXResult 响应结果
type KXResult int

const (
    KXLoginResultFail KXResult = iota
    KXLoginResultSuccess
)

// KXResponse 响应
type KXResponse struct {
    KXMessage
    Result KXResult `json:"result"` // 结果
}

// KXLoginRequest 上报登录请求
type KXLoginRequest struct {
    KXMessage
    IMSI            string `json:"imsi"`            // Sim卡IMSI信息
    HardVersion     string `json:"hardVersion"`     // 硬件版本
    SoftVersion     string `json:"softVersion"`     // 软件版本
    ProtocolVersion string `json:"protocolVersion"` // 协议版本
}

// KXReportRequest 上报电柜属性
type KXReportRequest struct {
    KXMessage
    AttrList []KXAttr `json:"attrList"` // 属性列表
    IsFull   int      `json:"isFull"`   // 是否全量上报: 0:增量 1:全量
}

// KXAttr 属性信息
type KXAttr struct {
    KXSignalData
    DoorID string `json:"doorId,omitempty"` // 柜门ID
}

// KXNoticeRequest 上报告警信息
type KXNoticeRequest struct {
    KXMessage
    AlarmList []KXAlarm `json:"alarmList"` // 告警列表
}

// KXAlarm 告警属性
type KXAlarm struct {
    ID         string `json:"id"`                   // 信号量ID
    AlarmTime  int64  `json:"alarmTime"`            // 告警时间: 13位时间戳
    AlarmmDesc string `json:"alarmmDesc,omitempty"` // 描述
    AlarmFlag  int    `json:"alarmFlag"`            // 告警标识: 0:结束 1:开始
    DoorID     string `json:"doorId,omitempty"`     // 柜门ID (对应信号量类别属于电池、柜门，字段必填)
    BatteryID  string `json:"batteryId,omitempty"`  // 电池ID (对应信号量类别属于电池，字段必填)
    UserID     string `json:"userId,omitempty"`     // App用户ID (首放、换电、退电的解绑、绑定事件必填)
}

// KXControlRequest 下发控制请求
type KXControlRequest struct {
    KXMessage
    ParamList []KXControlParam `json:"paramList"` // 控制参数列表
}

type KXScanBattery *int

var (
    KXScanBatterySupport    KXScanBattery = silk.Int(0)
    KXScanBatteryNotSupport KXScanBattery = silk.Int(1)
)

// KXControlParam 控制参数
type KXControlParam struct {
    KXSignalData
    DoorID      string        `json:"doorId,omitempty"`      // 柜门ID (如果控制的是柜门，字段必填，否则可以省略该字段)
    BatteryID   string        `json:"batteryId,omitempty"`   // 电池设备ID (如果控制的是电池，字段必填，否则可以省略该字段)
    Voltage     string        `json:"voltage,omitempty"`     // 电池电压 (首放流程中这个字段必填)
    ScanBattery KXScanBattery `json:"scanBattery,omitempty"` // 电池是否支持通讯 (0:支持 1:不支持，扫换电柜这个字段必填)
    UserID      string        `json:"userId,omitempty"`      // App用户ID (首放、换电、退电流程必填)
}

type KXSignal string

const (
    KXSignalLng                KXSignal = "02111001" // 柜子经度
    KXSignalLat                KXSignal = "02112001" // 柜子纬度
    KXSignalStatus             KXSignal = "02102001" // 换电柜状态 0:上电初始化 1:无换电、放电、取电动作 2:换电中 3:在归还电池中 4:在取出电池中 5:换电柜异常
    KXSignalDoorStatus         KXSignal = "02103001" // 柜门状态 0:关 1:开
    KXSignalBinStatus          KXSignal = "02104001" // 柜体状态 (仓位状态) 0:无电池 1:电池正在充电 2:电池充满 5:异常
    KXSignalBinID              KXSignal = "01309001" // 柜门ID 1~16:对应柜门ID
    KXSignalBatteryID          KXSignal = "01310001" // 换电柜设备ID
    KXSignalGSM                KXSignal = "02105001" // GSM 信号强度
    KXSignalBatterySN          KXSignal = "02106001" // 柜内电池SN
    KXSignalBatteryVoltage     KXSignal = "01111001" // 电池总电压 (V)
    KXSignalBatteryCurrent     KXSignal = "01112001" // 电池总电流 (A)
    KXSignalCabinetVoltage     KXSignal = "02107001" // 换电柜总电压 (V)
    KXSignalCabinetCurrent     KXSignal = "02108001" // 换电柜总电流 (A)
    KXSignalSOC                KXSignal = "02109001"
    KXSignalSOH                KXSignal = "02110001"
    KXSignalBatteryCellID      KXSignal = "01116001" // XX单芯电压串数 (表示电池内电芯总数量值)
    KXSignalBatteryCellVoltage KXSignal = "01117001" // XX单芯电压值 (XX如果为01，对应的信号量ID结尾为01)
    KXSignalCabinetTemp        KXSignal = "02113001" // 柜体温度值 (换电柜温度)
    KXSignalPCBTemp            KXSignal = "01118001" // 功率温度值 (电池内部PCB板表面温度)
    KXSignalBatteryCellTemp    KXSignal = "01119001" // 电芯温度值 (电池内部多组电芯中间表面温度)
    KXSignalBatteryAmbientTemp KXSignal = "01120001" // 环境温度 (电池壳体内部整体温度)
    KXSignalBatteryStatus      KXSignal = "01121001" // 电池状态 0:移动 1:静止 2:存储 3:休眠
    KXSignalBatteryControl     KXSignal = "01122001" // 电池控制
    KXSignalDischarge          KXSignal = "02116001" // 总放电 (Ah)
    KXSignalCharge             KXSignal = "02117001" // 总充电 (Ah)
    KXSignalBatteryChargeTime  KXSignal = "02114001" // 电池充电时长
    KXSignalBinEnable          KXSignal = "02118001" // 柜门是否禁用 (0:禁用 1:启用)
    KXSignalEnable             KXSignal = "02119001" // 柜体是否禁用 (0:禁用 1:启用)
    KXSignalEnergy             KXSignal = "02120001" // 柜子总用电量 (kwh)
    KXSignalCabinetControl     KXSignal = "02301001" // 控制换电柜命令
)

// KXSignalData 信号量
type KXSignalData struct {
    ID    KXSignal `json:"id"`    // 信号量ID
    Value string   `json:"value"` // 参数值
}
