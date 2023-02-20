// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-04
// Based on cabservd by liasica, magicrolan@qq.com.

package core

import (
    "github.com/auroraride/adapter/defs/cabdef"
    "go.uber.org/zap"
)

type (
    Hook interface {
        // OnConnect 连接接口
        OnConnect() (err error)

        // OnMessage 收到消息
        // serial 需要注册的电柜编号
        // fields zap日志字段
        OnMessage(b []byte, client *Client) (serial string, fields []zap.Field, err error)

        // SendControl 发送控制
        SendControl(serial string, typ cabdef.Operate, ordinal int) error

        // GetEmptyDeviation 获取空仓最大电压和电流
        // 空仓的时候有可能会有一定的电压和电流
        GetEmptyDeviation() (fakevoltage, fakecurrent float64)
    }

    Bean struct{}
)

func (h *Bean) OnConnect() (err error) {
    return
}

func (h *Bean) OnMessage(_ []byte, _ *Client) (serial string, fields []zap.Field, err error) {
    return
}

func (h *Bean) SendControl(serial string, typ cabdef.Operate, ordinal int) (err error) {
    return
}

func (h *Bean) GetEmptyDeviation() (voltage, current float64) {
    return
}
