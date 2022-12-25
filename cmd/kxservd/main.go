// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-04
// Based on cabservd by liasica, magicrolan@qq.com.

package main

import (
    "github.com/auroraride/cabservd/bridge"
    "github.com/auroraride/cabservd/internal"
    "github.com/auroraride/cabservd/internal/brands/kaixin"
    "github.com/auroraride/cabservd/internal/core"
    "github.com/auroraride/cabservd/internal/g"
    "github.com/auroraride/cabservd/internal/router"
)

func main() {
    // core boot
    internal.Boot()

    // 启动bridge
    go bridge.RunCabinet()

    // 启动 http server
    go router.Start()

    // 启动socket hub
    core.Start(
        g.Config.Tcp.Bind,
        g.Config.Brand,
        kaixin.New(),
        new(core.HeaderLength),
    )
}
