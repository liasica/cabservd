// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-04
// Based on cabservd by liasica, magicrolan@qq.com.

package main

import (
    "github.com/auroraride/cabservd/internal"
    "github.com/auroraride/cabservd/internal/brands/kaixin"
    "github.com/auroraride/cabservd/internal/core"
    "github.com/auroraride/cabservd/internal/g"
    "github.com/auroraride/cabservd/internal/hook"
    "github.com/auroraride/cabservd/internal/mem"
    "github.com/auroraride/cabservd/internal/router"
    "github.com/auroraride/cabservd/internal/service"
)

func main() {
    // core boot
    internal.Boot()

    // TODO 缓存数据?
    // cache()

    // TODO 启动bridge?
    // go bridge.Start()

    // 启动 http server
    go router.Start()

    // 启动socket hub
    go core.Start(
        g.Config.Tcp.Bind,
        g.Config.Brand,
        kaixin.New(),
        new(core.HeaderLength),
    )

    // 加载hooks
    hook.Start()

    select {}
}

func cache() {
    cabs := service.NewCabinet().QueryAllCabinetWithBin()
    for _, cab := range cabs {
        mem.SetCabinet(cab)
    }

    bins := service.NewBin().QueryAllBin()
    for _, b := range bins {
        mem.SetBin(b)
    }
}
