// Copyright (C) liasica. 2022-present.
//
// Created at 2022-10-31
// Based on cabservd by liasica, magicrolan@qq.com.

package hub

import (
    "github.com/panjf2000/gnet/v2"
    "sync"
)

type Hub struct {
    gnet.BuiltinEventEngine

    // address to listen
    addr string

    // 在线的客户端
    // key = devId
    // value = gnet.Conn
    clients sync.Map
}

func (h *Hub) Run(addr string) {
}
