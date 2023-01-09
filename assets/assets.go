// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-17
// Based on cabservd by liasica, magicrolan@qq.com.

package assets

import (
    _ "embed"
)

var (
    //go:embed config/config.yaml
    DefaultConfig []byte
)
