// Copyright (C) liasica. 2023-present.
//
// Created at 2023-03-14
// Based on cabservd by liasica, magicrolan@qq.com.

package yundong

import (
    "encoding/hex"
    "testing"
)

func TestParser_Login(t *testing.T) {
    p := &Parser{}
    b, _ := hex.DecodeString("03000a64108ae0276a4def4146")
    s, d := p.Login(b)
    t.Logf("%s => %s", s, d)
}
