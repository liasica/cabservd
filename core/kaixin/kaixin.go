// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-04
// Based on cabservd by liasica, magicrolan@qq.com.

package kaixin

import "cabservd/pkg/silk"

// MessageType 消息类型
type MessageType int

const (
    MessageTypeLoginRequest    MessageType = 100 // 登录请求
    MessageTypeLoginResponse   MessageType = 101 // 登录响应
    MessageTypeReportRequest   MessageType = 300 // 属性上报请求
    MessageTypeReportResponse  MessageType = 301 // 属性上报响应
    MessageTypeNoticeRequest   MessageType = 400 // 告警上报
    MessageTypeNoticeResponse  MessageType = 401 // 告警上报响应
    MessageTypeControlResponse MessageType = 500 // 电柜控制请求
    MessageTypeControlRequest  MessageType = 501 // 电柜控制响应
)

// RequestMessage 全局消息请求体
type RequestMessage struct {
    Message
    LoginRequest
    ReportRequest
    NoticeRequest
    ControlRequest
}

type Message struct {
    MessageType MessageType `json:"messageType"` // 消息类型
    TxnNo       int64       `json:"txnNo"`       // 流水号: 13位时间戳
    DevID       string      `json:"devId"`       // 设备ID
}

// Result 响应结果
type Result int

const (
    LoginResultFail Result = iota
    LoginResultSuccess
)

// Response 响应
type Response struct {
    Message
    Result Result `json:"result"` // 结果
}

// LoginRequest 上报登录请求
type LoginRequest struct {
    IMSI            string `json:"imsi,omitempty"`            // Sim卡IMSI信息
    HardVersion     string `json:"hardVersion,omitempty"`     // 硬件版本
    SoftVersion     string `json:"softVersion,omitempty"`     // 软件版本
    ProtocolVersion string `json:"protocolVersion,omitempty"` // 协议版本
}

// ReportRequest 上报电柜属性
type ReportRequest struct {
    AttrList []Attr `json:"attrList" json:"attrList,omitempty"` // 属性列表
    IsFull   int    `json:"isFull" json:"isFull,omitempty"`     // 是否全量上报: 0:增量 1:全量
}

// Attr 属性信息
type Attr struct {
    SignalData
    DoorID string `json:"doorId,omitempty"` // 柜门ID (可为空)
}

// NoticeRequest 上报告警信息
type NoticeRequest struct {
    AlarmList []Alarm `json:"alarmList,omitempty"` // 告警列表
}

// Alarm 告警属性
type Alarm struct {
    ID         string `json:"id,omitempty"`         // 信号量ID
    AlarmTime  int64  `json:"alarmTime,omitempty"`  // 告警时间: 13位时间戳
    AlarmmDesc string `json:"alarmmDesc,omitempty"` // 描述 (可为空)
    AlarmFlag  int    `json:"alarmFlag,omitempty"`  // 告警标识: 0:结束 1:开始
    DoorID     string `json:"doorId,omitempty"`     // 柜门ID (可为空, 对应信号量类别属于电池、柜门，字段必填)
    BatteryID  string `json:"batteryId,omitempty"`  // 电池ID (可为空, 对应信号量类别属于电池，字段必填)
    UserID     string `json:"userId,omitempty"`     // App用户ID (可为空, 首放、换电、退电的解绑、绑定事件必填)
}

// ControlRequest 下发控制请求
type ControlRequest struct {
    ParamList []ControlParam `json:"paramList,omitempty"` // 控制参数列表
}

type ScanBattery *int

var (
    ScanBatterySupport    ScanBattery = silk.Int(0)
    ScanBatteryNotSupport ScanBattery = silk.Int(1)
)

// ControlParam 控制参数
type ControlParam struct {
    SignalData
    DoorID      string      `json:"doorId,omitempty"`      // 柜门ID (可为空, 如果控制的是柜门，字段必填，否则可以省略该字段)
    BatteryID   string      `json:"batteryId,omitempty"`   // 电池设备ID (可为空, 如果控制的是电池，字段必填，否则可以省略该字段)
    Voltage     string      `json:"voltage,omitempty"`     // 电池电压 (可为空, 首放流程中这个字段必填)
    ScanBattery ScanBattery `json:"scanBattery,omitempty"` // 电池是否支持通讯 (可为空, 0:支持 1:不支持，扫换电柜这个字段必填)
    UserID      string      `json:"userId,omitempty"`      // App用户ID (可为空, 首放、换电、退电流程必填)
}
