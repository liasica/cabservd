// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-15
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

import (
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"

	"github.com/auroraride/cabservd/internal/codec"
)

// ApiResponse 平台响应
type ApiResponse[T any] struct {
	Code      int    `json:"code"`
	RequestID string `json:"requestID"`
	ErrCode   string `json:"errCode,omitempty"`
	ErrMsg    string `json:"errMsg,omitempty"`
	Data      T      `json:"data,omitempty"`
	SysTime   int64  `json:"sysTime"`
}

// NotifyResponse 回调响应
type NotifyResponse[T any] struct {
	Code      int    `json:"code"`
	RequestID string `json:"requestID"`
	ErrorMsg  string `json:"errorMsg,omitempty"`
	Data      T      `json:"data,omitempty"`
}

func (r *NotifyResponse[T]) Bytes() []byte {
	b, _ := jsoniter.Marshal(r)
	return b
}

func (r *NotifyResponse[T]) GetMessage(c codec.Codec) ([]byte, []zap.Field) {
	b := r.Bytes()
	return c.Encode(b), []zap.Field{zap.ByteString("data", b)}
}
