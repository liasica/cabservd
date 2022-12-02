// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-30
// Based on cabservd by liasica, magicrolan@qq.com.

package service

import "github.com/auroraride/cabservd/internal/ent"

type exchange struct {
    serial string
    cab    *ent.Cabinet
    bins   ent.Bins
}

func NewExchange(serial string, userID uint64) *exchange {
    return &exchange{}
}
