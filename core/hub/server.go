// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-02
// Based on cabservd by liasica, magicrolan@qq.com.

package hub

import (
    "github.com/panjf2000/gnet/v2"
    log "github.com/sirupsen/logrus"
)

func Run(addr string) {
    h := &hub{
        addr: addr,
    }
    log.Fatal(gnet.Run(
        h,
        h.addr,
        gnet.WithMulticore(true),
        gnet.WithReuseAddr(true),
        gnet.WithLogger(log.New()),
    ))
}
