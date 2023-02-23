// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-04
// Based on cabservd by liasica, magicrolan@qq.com.

package kaixin

import (
    "context"
    "github.com/auroraride/adapter"
    "github.com/auroraride/adapter/log"
    "github.com/auroraride/cabservd/internal/codec"
    "github.com/auroraride/cabservd/internal/core"
    jsoniter "github.com/json-iterator/go"
    "go.uber.org/zap"
)

type Handler struct {
    core.Bean
}

func New() (core.Hook, codec.Codec) {
    return &Handler{}, &codec.HeaderLength{}
}

// GetEmptyDeviation TODO 后续做在数据库中
func (h *Handler) GetEmptyDeviation() (voltage, current float64) {
    voltage = 40
    current = 1
    return
}

// OnMessage 解析消息
func (h *Handler) OnMessage(_ *core.Client, b []byte) (serial string, res core.ResponseMessenger, fields []zap.Field, err error) {
    fields = []zap.Field{
        zap.ByteString("decoded", b),
    }

    req := new(Request)
    err = jsoniter.Unmarshal(b, req)
    if err != nil {
        return
    }

    serial = req.DevID
    fields = append(fields, log.Payload(req))

    switch req.MsgType {
    case MessageTypeLoginRequest:
        err = h.LoginHandle(req)
    case MessageTypeReportRequest:
        err = h.ReportHandle(req)
    case MessageTypeNoticeRequest:
        err = h.NoticeHandle(req)
    case MessageTypeControlResponse:
        // TODO 控制成功逻辑
        // 收到成功逻辑处理完成后, 不发送反馈消息
        return
    }

    // 发送失败响应
    if err != nil {
        res = req.Fail()
        return
    }

    res = req.Success()

    return
}

// LoginHandle 登录请求
func (h *Handler) LoginHandle(req *Request) (err error) {
    if req.DevID == "" {
        return adapter.ErrorCabinetSerialRequired
    }

    // // 清除仓位电池信息
    // // TODO 清除的时候会不会后来的消息先到
    // err = core.ResetBins(req.DevID)
    // if err != nil {
    //     return
    // }

    // 查找或创建电柜
    go core.LoadOrStoreCabinet(context.Background(), req.DevID)

    // TODO: 保存其他信息
    return
}

// ReportHandle 状态上报请求
func (h *Handler) ReportHandle(req *Request) (err error) {
    if req.DevID == "" {
        return adapter.ErrorCabinetSerialRequired
    }
    core.UpdateCabinet(req)
    return
}

// NoticeHandle 告警上报请求
func (h *Handler) NoticeHandle(req *Request) (err error) {
    // TODO 解读并保存所有告警信息
    return
}
