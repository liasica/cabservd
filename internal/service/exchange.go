// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-30
// Based on cabservd by liasica, magicrolan@qq.com.

package service

import (
    "github.com/auroraride/cabservd/internal/types"
)

type exchange struct {
    user *types.User
}

func NewExchange(user *types.User) *exchange {
    return &exchange{
        user: user,
    }
}
