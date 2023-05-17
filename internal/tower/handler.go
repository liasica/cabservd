// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-04
// Based on cabservd by liasica, magicrolan@qq.com.

package tower

import (
	"context"

	"github.com/auroraride/adapter"
	"github.com/auroraride/adapter/log"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"

	"github.com/auroraride/cabservd/internal/core"
)

var (
	cabinetSignals map[Signal]CabinetSignalFunc
	binSignals     map[Signal]BinSignalFunc
)

type Handler struct {
	core.Bean
	mtl *MessageTypeList
}

func New(options ...Option) (h *Handler) {
	h = &Handler{}
	for _, o := range options {
		o.apply(h)
	}

	if h.mtl == nil {
		h.mtl = &MessageTypeList{
			LoginRequest:    100,
			LoginResponse:   101,
			ReportRequest:   300,
			ReportResponse:  301,
			NoticeRequest:   400,
			NoticeResponse:  401,
			ControlRequest:  500,
			ControlResponse: 501,
		}
	}

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
	case h.mtl.LoginRequest:
		err = h.LoginHandle(req)
	case h.mtl.ReportRequest:
		err = h.ReportHandle(req)
	case h.mtl.NoticeRequest:
		err = h.NoticeHandle(req)
	case h.mtl.ControlResponse:
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
	core.UpdateCabinet(h, req)
	return
}

// NoticeHandle 告警上报请求
func (h *Handler) NoticeHandle(req *Request) (err error) {
	// TODO 解读并保存所有告警信息
	return
}
