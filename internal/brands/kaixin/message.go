// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-06
// Based on cabservd by liasica, magicrolan@qq.com.

package kaixin

import (
    "github.com/auroraride/adapter/log"
    "github.com/auroraride/cabservd/internal/codec"
    jsoniter "github.com/json-iterator/go"
    "go.uber.org/zap"
)

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
