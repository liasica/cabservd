// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-15
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

type MessageToCabinet struct {
	Version   string `json:"version"`
	Timestamp int64  `json:"timestamp"`
	RequestId string `json:"requestId"`
	AppId     string `json:"appId"`
	Sign      string `json:"sign"`
	Biz       string `json:"biz"`
}
