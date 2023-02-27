// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-27
// Based on cabservd by liasica, magicrolan@qq.com.

package yundong

import (
    "fmt"
    "github.com/auroraride/adapter"
    "testing"
)

func TestGetQRCode(t *testing.T) {
    sn := "NCAWDFA0L751N027"
    var (
        sb [16]byte
        qr [128]byte
    )
    copy(sb[:], adapter.ConvertString2Bytes(sn))
    copy(qr[:], adapter.ConvertString2Bytes(sn))
    message := append(sb[:], qr[:]...)
    fmt.Println(message)
}
