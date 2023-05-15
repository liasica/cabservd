// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-15
// Based on cabservd by liasica, magicrolan@qq.com.

package core

type Option interface {
	apply(h *hub)
}

type optionFunc func(h *hub)

func (f optionFunc) apply(h *hub) {
	f(h)
}
