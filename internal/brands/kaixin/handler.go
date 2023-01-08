// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-04
// Based on cabservd by liasica, magicrolan@qq.com.

package kaixin

import (
    "context"
    "github.com/auroraride/adapter"
    "github.com/auroraride/adapter/defs/cabdef"
    "github.com/auroraride/cabservd/internal/core"
    jsoniter "github.com/json-iterator/go"
)

type Hander struct {
    core.Bean
}

func New() *Hander {
    return &Hander{}
}

// GetEmptyDeviation TODO 后续做在数据库中
func (h *Hander) GetEmptyDeviation() (voltage, current float64) {
    voltage = 40
    current = 1
    return
}

// OnMessage 解析消息
func (h *Hander) OnMessage(b []byte, client *core.Client) (err error) {
    req := new(Request)
    err = jsoniter.Unmarshal(b, req)
    if err != nil {
        return
    }
    switch req.MsgType {
    case MessageTypeLoginRequest:
        err = h.LoginHandle(req, client)
    case MessageTypeReportRequest:
        err = h.ReportHandle(req)
    case MessageTypeNoticeRequest:
        err = h.NoticeHandle(req)
    case MessageTypeControlResponse:
        // TODO 控制成功逻辑
        // 收到成功逻辑处理完成后, 不发送反馈消息
        return
    }

    // 发送登录响应
    if err != nil {
        return client.SendMessage(req.Fail())
    }

    return client.SendMessage(req.Success())
}

// LoginHandle 登录请求
func (h *Hander) LoginHandle(req *Request, client *core.Client) (err error) {
    if req.DevID == "" {
        return adapter.ErrorCabinetSerialRequired
    }

    // // 清除仓位电池信息
    // // TODO 清除的时候会不会后来的消息先到
    // err = core.ResetBins(req.DevID)
    // if err != nil {
    //     return
    // }

    // 保存设备识别码
    client.Serial = req.DevID

    // 注册连接
    client.Register()

    // 查找或创建电柜
    core.LoadOrStoreCabinet(context.Background(), cabdef.BrandKaixin, req.DevID)

    // TODO: 保存其他信息
    return
}

// ReportHandle 状态上报请求
func (h *Hander) ReportHandle(req *Request) (err error) {
    if req.DevID == "" {
        return adapter.ErrorCabinetSerialRequired
    }
    core.UpdateCabinet(cabdef.BrandKaixin, req)
    return
}

// NoticeHandle 告警上报请求
func (h *Hander) NoticeHandle(req *Request) (err error) {
    // TODO 解读并保存所有告警信息
    return
}
