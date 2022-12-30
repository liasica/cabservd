// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-04
// Based on cabservd by liasica, magicrolan@qq.com.

package main

import (
    "context"
    "github.com/auroraride/cabservd/internal"
    "github.com/auroraride/cabservd/internal/brands/kaixin"
    "github.com/auroraride/cabservd/internal/core"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/cabinet"
    "github.com/auroraride/cabservd/internal/g"
    "github.com/auroraride/cabservd/internal/mem"
    "github.com/auroraride/cabservd/internal/notice"
    "github.com/auroraride/cabservd/internal/router"
    "github.com/auroraride/cabservd/internal/service"
)

func main() {
    // core boot
    internal.Boot()

    // 标记所有电柜为离线和空闲
    _ = ent.Database.Cabinet.Update().SetOnline(false).SetStatus(cabinet.StatusIdle).Exec(context.Background())

    // TODO 缓存数据?
    // cache()

    // 加载hooks
    notice.Start()

    // 启动 http server
    go router.Start()

    // 启动socket hub
    go core.Start(
        g.Config.Tcp.Bind,
        g.Config.Brand,
        kaixin.New(),
        new(core.HeaderLength),
    )

    select {}
}

func cache() {
    cabs := service.NewCabinet(service.PermissionNotRequired).QuerySerialWithBinAll()
    for _, cab := range cabs {
        mem.SetCabinet(cab)
    }

    bins := service.NewBin(service.PermissionNotRequired).QueryAllBin()
    for _, b := range bins {
        mem.SetBin(b)
    }
}
