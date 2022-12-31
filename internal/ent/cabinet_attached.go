// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-03
// Based on cabservd by liasica, magicrolan@qq.com.

package ent

import "github.com/auroraride/adapter"

func (c *Cabinet) DetectOnline() (err error) {
    if !c.Online {
        err = adapter.CabinetOffline
    }
    return
}

func (c *Cabinet) DetectExchangeTask() (err error) {
    err = c.DetectOnline()
    if err != nil {
        return
    }
    return
}

func (c *Cabinet) GetID() uint64 {
    return c.ID
}

func (c *Cabinet) GetSerial() string {
    return c.Serial
}
