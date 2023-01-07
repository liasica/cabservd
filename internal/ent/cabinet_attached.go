// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-03
// Based on cabservd by liasica, magicrolan@qq.com.

package ent

import (
    "fmt"
    "github.com/auroraride/adapter"
    "strings"
)

func (c *Cabinet) DetectOnline() (err error) {
    if !c.Online {
        err = adapter.ErrorCabinetOffline
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

func (c *Cabinet) GetListenerKey() string {
    return fmt.Sprintf("%s-%d", c.GetTableName(), c.ID)
}

func (c *CabinetPointer) String() string {
    var builder strings.Builder
    builder.WriteString("电柜[")
    builder.WriteString(*c.Serial)
    builder.WriteString("]变动 ->")

    if c.Online != nil {
        builder.WriteString(" 在线=")
        builder.WriteString(adapter.Bool(*c.Online).String())
    }

    if c.Power != nil {
        builder.WriteString(" 市电=")
        builder.WriteString(adapter.Bool(*c.Power).String())
    }

    if c.Status != nil {
        builder.WriteString(" 状态=")
        builder.WriteString(c.Status.String())
    }

    if c.Enable != nil {
        builder.WriteString(" 启用=")
        builder.WriteString(adapter.Bool(*c.Enable).String())
    }

    return builder.String()
}
