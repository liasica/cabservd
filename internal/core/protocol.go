// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-15
// Based on cabservd by liasica, magicrolan@qq.com.

package core

type Protocol int

const (
	ProtocolTcp Protocol = iota
	ProtocolHttp
)

func (p Protocol) Http() bool {
	return p == ProtocolHttp
}

func (p Protocol) Tcp() bool {
	return p == ProtocolTcp
}
