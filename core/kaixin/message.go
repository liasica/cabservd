// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-06
// Based on cabservd by liasica, magicrolan@qq.com.

package kaixin

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

// Request 全局消息请求体
type Request struct {
    Message
    LoginRequest
    ReportRequest
    NoticeRequest
    ControlRequest
}

// String TODO 转为String
func (m *Request) String() string {
    return "TODO"
}

// Success 成功响应
func (m *Request) Success() *Response {
    return &Response{
        Message: Message{
            MessageType: m.MessageType + 1,
            TxnNo:       m.TxnNo,
            DevID:       m.DevID,
        },
        Result: LoginResultSuccess,
    }
}

// Fail 响应失败
func (m *Request) Fail() *Response {
    return &Response{
        Message: Message{
            MessageType: m.MessageType + 1,
            TxnNo:       m.TxnNo,
            DevID:       m.DevID,
        },
        Result: LoginResultFail,
    }
}
