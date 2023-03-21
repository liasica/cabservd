// Copyright (C) liasica. 2023-present.
//
// Created at 2023-03-17
// Based on cabservd by liasica, magicrolan@qq.com.

package main

import (
    "github.com/auroraride/cabservd/internal"
    "github.com/auroraride/cabservd/internal/brands/kaixin"
)

func main() {
    internal.Boot(kaixin.NewNonIntelligent)
}
