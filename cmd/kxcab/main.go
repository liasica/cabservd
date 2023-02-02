// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-04
// Based on cabservd by liasica, magicrolan@qq.com.

package main

import (
    "github.com/auroraride/adapter/codec"
    "github.com/auroraride/cabservd/internal"
    "github.com/auroraride/cabservd/internal/brands/kaixin"
)

func main() {
    internal.Boot(kaixin.New(), &codec.HeaderLength{})
}
