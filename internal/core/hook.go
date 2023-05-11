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
		OnConnect(c *Client)

		// OnMessage 收到消息
		// serial 需要注册的电柜编号
		// fields zap日志字段
		OnMessage(c *Client, b []byte) (serial string, res ResponseMessenger, fields []zap.Field, err error)

		// SendOperate 发送主要控制
		SendOperate(serial string, typ cabdef.Operate, ordinal int, times int) error

		// GetEmptyDeviation 获取空仓最大电压和电流
		// 空仓的时候有可能会有一定的电压和电流
		GetEmptyDeviation() (fakevoltage, fakecurrent float64)
	}

	Bean struct{}
)

func (h *Bean) OnConnect(*Client) {
	return
}

func (h *Bean) OnMessage(_ *Client, _ []byte) (serial string, _ ResponseMessenger, fields []zap.Field, err error) {
	return
}

func (h *Bean) SendOperate(_ string, _ cabdef.Operate, _ int, _ int) (err error) {
	return
}

func (h *Bean) GetEmptyDeviation() (voltage, current float64) {
	return
}
