// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-04
// Based on cabservd by liasica, magicrolan@qq.com.

package main

import (
    "github.com/auroraride/cabservd/internal"
    "github.com/auroraride/cabservd/internal/core"
    "github.com/auroraride/cabservd/internal/core/kaixin"
    "github.com/auroraride/cabservd/router"
)

func main() {
    // core boot
    internal.Boot()

    // 启动 http server
    go router.Start()

    // 启动socket hub
    core.Start(
        "0.0.0.0:18511",
        "凯信",
        kaixin.New(),
        new(core.HeaderLength),
    )
}
