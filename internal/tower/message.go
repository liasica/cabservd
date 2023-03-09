// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-06
// Based on cabservd by liasica, magicrolan@qq.com.

package tower

import (
    "github.com/auroraride/adapter/log"
    "github.com/auroraride/cabservd/internal/codec"
    jsoniter "github.com/json-iterator/go"
    "github.com/liasica/go-helpers/silk"
    "go.uber.org/zap"
)

// MessageType 消息类型
type MessageType int

type MessageTypeList struct {
    LoginRequest    MessageType // 登录请求
    LoginResponse   MessageType // 登录响应
    ReportRequest   MessageType // 属性上报请求
    ReportResponse  MessageType // 属性上报响应
    NoticeRequest   MessageType // 告警上报
    NoticeResponse  MessageType // 告警上报响应
    ControlRequest  MessageType // 电柜控制请求
    ControlResponse MessageType // 电柜控制响应
}

type Message struct {
    MsgType MessageType `json:"msgType"` // 消息类型
    TxnNo   int64       `json:"txnNo"`   // 流水号: 13位时间戳
    DevID   string      `json:"devId"`   // 设备ID
}

// LoginRequest 上报登录请求
type LoginRequest struct {
    IMSI            string `json:"imsi,omitempty"`            // Sim卡IMSI信息
    HardVersion     string `json:"hardVersion,omitempty"`     // 硬件版本
    SoftVersion     string `json:"softVersion,omitempty"`     // 软件版本
    ProtocolVersion string `json:"protocolVersion,omitempty"` // 协议版本
}

// ReportCate 上报类型
type ReportCate int

const (
    ReportCateIncremental ReportCate = iota // 增量
    ReportCateFull                          // 全量
)

func (c ReportCate) String() string {
    switch c {
    case ReportCateIncremental:
        return "增量"
    case ReportCateFull:
        return "全量"
    }
    return ""
}

// ReportRequest 上报电柜属性
type ReportRequest struct {
    AttrList Attributes `json:"attrList,omitempty"` // 属性列表
    IsFull   ReportCate `json:"isFull,omitempty"`   // 是否全量上报: 0:增量 1:全量
}

// NoticeRequest 上报告警信息
type NoticeRequest struct {
    AlarmList []Alarm `json:"alarmList,omitempty"` // 告警列表
}

// AlarmFlag 告警标识
type AlarmFlag int

const (
    AlarmFlagEnd AlarmFlag = iota
    AlarmFlagStart
)

func (a AlarmFlag) String() string {
    switch a {
    case AlarmFlagEnd:
        return "结束"
    case AlarmFlagStart:
        return "开始"
    }
    return ""
}

// Alarm 告警属性
type Alarm struct {
    ID         string    `json:"id,omitempty"`         // 信号量ID
    AlarmTime  int64     `json:"alarmTime,omitempty"`  // 告警时间: 13位时间戳
    AlarmmDesc string    `json:"alarmmDesc,omitempty"` // 描述 (可为空)
    AlarmFlag  AlarmFlag `json:"alarmFlag,omitempty"`  // 告警标识: 0:结束 1:开始
    DoorID     string    `json:"doorId,omitempty"`     // 柜门ID (可为空, 对应信号量类别属于电池、柜门，字段必填)
    BatteryID  string    `json:"batteryId,omitempty"`  // 电池ID (可为空, 对应信号量类别属于电池，字段必填)
    UserID     string    `json:"userId,omitempty"`     // App用户ID (可为空, 首放、换电、退电的解绑、绑定事件必填)
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
    DoorID      string      `json:"doorId,omitempty"`      // 柜门ID (从1开始, 可为空, 如果控制的是柜门，字段必填，否则可以省略该字段)
    BatteryID   string      `json:"batteryId,omitempty"`   // 电池设备ID (可为空, 如果控制的是电池，字段必填，否则可以省略该字段)
    Voltage     string      `json:"voltage,omitempty"`     // 电池电压 (可为空, 首放流程中这个字段必填)
    ScanBattery ScanBattery `json:"scanBattery,omitempty"` // 电池是否支持通讯 (可为空, 0:支持 1:不支持，扫换电柜这个字段必填)
    UserID      string      `json:"userId,omitempty"`      // App用户ID (可为空, 首放、换电、退电流程必填)
}

// Result 响应结果
type Result int

const (
    LoginResultFail Result = iota
    LoginResultSuccess
)

func getMessage(c codec.Codec, data any) (message []byte, fields []zap.Field) {
    b, _ := jsoniter.Marshal(data)
    message = c.Encode(b)
    log.ResponseBody(b)
    fields = []zap.Field{
        zap.ByteString("data", b),
    }
    return
}

// Response 响应
type Response struct {
    Message
    Result Result `json:"result"` // 结果
}

func (r *Response) GetMessage(c codec.Codec) ([]byte, []zap.Field) {
    return getMessage(c, r)
}

// Request 全局消息请求体
type Request struct {
    Message
    LoginRequest
    ReportRequest
    NoticeRequest
    ControlRequest
}

// String TODO 转为String
func (r *Request) String() string {
    return "TODO"
}

func (r *Request) GetMessage(c codec.Codec) ([]byte, []zap.Field) {
    return getMessage(c, r)
}

// Success 成功响应
func (r *Request) Success() *Response {
    return &Response{
        Message: Message{
            MsgType: r.MsgType + 1,
            TxnNo:   r.TxnNo,
            DevID:   r.DevID,
        },
        Result: LoginResultSuccess,
    }
}

// Fail 响应失败
func (r *Request) Fail() *Response {
    return &Response{
        Message: Message{
            MsgType: r.MsgType + 1,
            TxnNo:   r.TxnNo,
            DevID:   r.DevID,
        },
        Result: LoginResultFail,
    }
}
