// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-18
// Based on cabservd by liasica, magicrolan@qq.com.

package main

import (
    "github.com/auroraride/cabservd/internal"
    "github.com/auroraride/cabservd/internal/brands/yundong"
)

func main() {
    internal.Boot(yundong.New)
}
