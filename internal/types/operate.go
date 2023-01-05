// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-01
// Based on cabservd by liasica, magicrolan@qq.com.

package types

import (
    "github.com/auroraride/adapter/defs/cabdef"
)

type StepCallback func(*cabdef.BinOperateResult)
