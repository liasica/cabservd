// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-02
// Based on cabservd by liasica, magicrolan@qq.com.

package core

import (
    "github.com/auroraride/adapter"
    "github.com/panjf2000/gnet/v2"
    log "github.com/sirupsen/logrus"
)

func Start(addr string, brand adapter.Brand, bean Hook, codec Codec) {
    Hub = &hub{
        addr:       addr,
        Bean:       bean,
        brand:      brand,
        codec:      codec,
        connect:    make(chan *Client),
        disconnect: make(chan *Client),
    }

    go Hub.run()

    log.Fatal(gnet.Run(
        Hub,
        Hub.addr,
        gnet.WithMulticore(true),
        gnet.WithReuseAddr(true),
        gnet.WithLogger(log.StandardLogger()),
    ))
}

func (h *hub) run() {
    for {
        select {
        case client := <-h.connect:
            h.clients.Store(client, "")
        case client := <-h.disconnect:
            if _, ok := h.clients.Load(client); ok {
                h.clients.Delete(client)
            }
            close(client.receiver)
        }
    }
}
