// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-15
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

import (
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"

	"github.com/auroraride/cabservd/internal/core"
)

// ApiResponse 接口响应, 西六楼 -> 平台
type ApiResponse[T any] struct {
	Code      int    `json:"code"`
	RequestID string `json:"requestID"`
	ErrCode   string `json:"errCode,omitempty"`
	ErrMsg    string `json:"errMsg,omitempty"`
	Data      T      `json:"data,omitempty"`
	SysTime   int64  `json:"sysTime"`
}

// NotifyResult 回调响应, 平台 -> 西六楼
type NotifyResult[T any] struct {
	Code      int    `json:"code"`
	RequestID string `json:"requestID"`
	ErrorMsg  string `json:"errorMsg,omitempty"`
	Data      T      `json:"data,omitempty"`
}

func (r *NotifyResult[T]) Bytes() []byte {
	b, _ := jsoniter.Marshal(r)
	return b
}

func (r *NotifyResult[T]) GetMessage(c core.Codec) ([]byte, []zap.Field) {
	b := r.Bytes()
	return c.Encode(b), []zap.Field{zap.ByteString("data", b)}
}

// CommandResponse 硬件操作结果, 西六楼 -> 平台
type CommandResponse[T any] struct {
	Command   string `json:"command"`        // 指令
	RequestID string `json:"requestID"`      // 每次调用硬件操作接口使用的requestId
	Result    int    `json:"result"`         // 0--成功 1--失败
	Code      string `json:"code,omitempty"` // 错误编号
	Data      T      `json:"data,omitempty"` // 结果
}
