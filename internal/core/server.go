// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-02
// Based on cabservd by liasica, magicrolan@qq.com.

package core

import (
    "github.com/auroraride/adapter/defs/cabdef"
    "github.com/panjf2000/gnet/v2"
    log "github.com/sirupsen/logrus"
)

func Start(addr string, brand cabdef.Brand, bean Hook, codec Codec) {
    Hub = &hub{
        addr:       addr,
        Bean:       bean,
        brand:      brand,
        codec:      codec,
        register:   make(chan *Client),
        disconnect: make(chan *Client),
    }

    go Hub.run()
    // go Hub.deadCheck()

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
        case client := <-h.register:
            h.clients.Store(client.Serial, client)
        case client := <-h.disconnect:
            go client.Close()
        }
    }
}

// 每隔1分钟标记20分之前更新的电柜为离线
// TODO 是否发送消息
func (h *hub) deadCheck() {
    // ticker := time.NewTicker(time.Minute)
    // for {
    //     select {
    //     case t := <-ticker.C:
    //         _ = ent.Database.GetCabinet.Update().
    //             Where(
    //                 cabinet.Brand(g.Config.Brand),
    //                 cabinet.UpdatedAtLT(t.Add(-20*time.Minute)),
    //             ).
    //             SetOnline(false).
    //             Exec(context.Background())
    //     }
    // }
}
