// Copyright (C) liasica. 2022-present.
//
// Created at 2022-10-31
// Based on cabservd by liasica, magicrolan@qq.com.

package main

import (
    "cabservd/core/hub"
    "cabservd/internal"
)

func main() {
    // core boot
    internal.Boot()

    // 启动socket hub
    hub.Run("0.0.0.0:18531")
}
