// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-04
// Based on cabservd by liasica, magicrolan@qq.com.

package core

import "github.com/auroraride/cabservd/internal/types"

type (
    Hook interface {
        // OnConnect 连接接口
        OnConnect() (err error)

        // OnMessage 收到消息
        OnMessage(b []byte, client *Client) (err error)

        // OnControl 发送控制
        OnControl(serial string, typ types.ControlType, ordinal int) error

        // GetEmptyFake 获取空仓最大电压和电流
        // 空仓的时候有可能会有一定的电压和电流
        GetEmptyFake() (fakevoltage, fakecurrent float64)
    }

    Bean struct{}
)

func (h *Bean) OnConnect() (err error) {
    return
}

func (h *Bean) OnMessage(_ []byte, _ *Client) (err error) {
    return
}

func (h *Bean) OnControl(serial string, typ types.ControlType, ordinal int) (err error) {
    return
}

func (h *Bean) GetEmptyFake() (voltage, current float64) {
    return
}
