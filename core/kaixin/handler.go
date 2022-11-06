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
    switch req.MessageType {
    case MessageTypeLoginRequest:
        return h.login(req, client)
    }
    return
}

// 登录请求
func (h *Hander) login(req *Request, client *core.Client) (err error) {
    // TODO: 保存其他信息

    // 保存设备识别码
    client.SetDeviceID(req.DevID)

    // 发送登录响应
    return client.SendMessage(req.Success())
}
