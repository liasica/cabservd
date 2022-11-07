// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-04
// Based on cabservd by liasica, magicrolan@qq.com.

package kaixin

import (
    "cabservd/core"
    jsoniter "github.com/json-iterator/go"
)

type Hander struct {
    core.Bean
}

func New() *Hander {
    return &Hander{}
}

// OnMessage 解析消息
func (h *Hander) OnMessage(b []byte, client *core.Client) (err error) {
    req := new(Request)
    // err = req.UnmarshalBinary(b)
    err = jsoniter.Unmarshal(b, req)
    if err != nil {
        return
    }
    switch req.MsgType {
    case MessageTypeLoginRequest:
        err = h.login(req, client)
    case MessageTypeReportRequest:
        err = h.report(req)
    case MessageTypeNoticeRequest:
        err = h.notice(req)
    case MessageTypeControlResponse:
        // TODO 控制成功逻辑
    }

    // 发送登录响应
    if err != nil {
        return client.SendMessage(req.Fail())
    }

    return client.SendMessage(req.Success())
}

// 登录请求
func (h *Hander) login(req *Request, client *core.Client) (err error) {
    // TODO: 保存其他信息

    // 保存设备识别码
    client.SetDeviceID(req.DevID)

    return
}

// 状态上报请求
func (h *Hander) report(req *Request) (err error) {
    // TODO 解读所有信号量
    return
}

// 告警上报请求
func (h *Hander) notice(req *Request) (err error) {
    // TODO 解读所有告警信息
    return
}
