// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-29
// Based on cabservd by liasica, magicrolan@qq.com.

package core

import "github.com/auroraride/cabservd/internal/ent"

type BinParser interface {
    Bins() ent.BinPointers
    // TODO Cabinet
}
