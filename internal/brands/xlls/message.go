// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-15
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

import (
	"github.com/google/uuid"
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

func generateRequestID() string {
	return uuid.New().String()
}
