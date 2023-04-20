// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-15
// Based on yundong by liasica, magicrolan@qq.com.

package yundong

import "errors"

var (
	ErrInvaildCommand     = errors.New("指令错误")
	ErrInvaildFieldData   = errors.New("字段数据错误")
	ErrIncompletePacket   = errors.New("incomplete packet")
	ErrIncorrectPacket    = errors.New("消息错误")
	ErrTimeout            = errors.New("客户端空闲超时")
	ErrClientNotFound     = errors.New("未获取到在线客户端")
	ErrPassthroughCommand = errors.New("透传命令错误")
)
