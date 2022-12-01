// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-04
// Based on cabservd by liasica, magicrolan@qq.com.

package core

import "github.com/auroraride/cabservd/types"

type (
    Hook interface {
        // OnConnect 连接接口
        OnConnect() (err error)

        // OnMessage 收到消息
        OnMessage(b []byte, client *Client) (err error)

        // OnControl 发送控制
        OnControl(serial string, typ types.ControlType, index int) error
    }

    Bean struct{}
)

func (h *Bean) OnConnect() (err error) {
    return
}

func (h *Bean) OnMessage(_ []byte, _ *Client) (err error) {
    return
}

func OnControl(serial string, typ types.ControlType, index int) (err error) {
    return
}
