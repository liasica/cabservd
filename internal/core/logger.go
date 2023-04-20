// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-15
// Based on yundong by liasica, magicrolan@qq.com.

package core

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func (c *Client) logPrefix() (str string) {
	defer func() {
		if v := recover(); v != nil {
			zap.L().WithOptions(zap.WithCaller(false)).Error("捕获错误", zap.Error(fmt.Errorf("%v", v)))
		}
	}()
	str = "[" + c.Conn.RemoteAddr().String() + "(" + c.Serial + ")" + "] "
	return
}

func (c *Client) Info(msg string, fields ...zap.Field) {
	zap.L().Info(c.logPrefix()+msg, fields...)
}

func (c *Client) Error(msg string, fields ...zap.Field) {
	zap.L().Error(c.logPrefix()+msg, fields...)
}

func (c *Client) Log(lvl zapcore.Level, msg string, fields ...zap.Field) {
	zap.L().Log(lvl, c.logPrefix()+msg, fields...)
}
