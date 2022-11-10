// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-04
// Based on cabservd by liasica, magicrolan@qq.com.

package main

import (
    "cabservd/core"
    "cabservd/core/kaixin"
    "cabservd/internal"
)

func main() {
    // core boot
    internal.Boot()

    // 启动socket hub
    core.Start(
        "0.0.0.0:18531",
        "凯信",
        kaixin.New(),
        new(core.HeaderLength),
    )
}
