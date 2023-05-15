// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-15
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

// Response 平台响应
type Response[T any] struct {
	Code      int    `json:"code,omitempty"`
	RequestID string `json:"requestID,omitempty"`
	ErrCode   string `json:"errCode,omitempty"`
	ErrMsg    string `json:"errMsg,omitempty"`
	Data      T      `json:"data,omitempty"`
	SysTime   int64  `json:"sysTime,omitempty"`
}

// NotifyResponse 回调响应
type NotifyResponse struct {
	Code      int    `json:"code,omitempty"`
	RequestID string `json:"requestID,omitempty"`
	ErrorMsg  string `json:"errorMsg,omitempty"`
	Data      string `json:"data,omitempty"`
}
