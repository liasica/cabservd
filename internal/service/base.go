// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on cabservd by liasica, magicrolan@qq.com.

package service

import (
    "context"
    "github.com/auroraride/adapter"
    "github.com/auroraride/cabservd/internal/app"
    "net/http"
)

type Permission bool

const (
    PermissionRequired    Permission = true
    PermissionNotRequired Permission = false
)

type BaseService struct {
    User *adapter.User
    ctx  context.Context
}

func newService(params ...any) *BaseService {
    nq := PermissionRequired
    s := &BaseService{
        ctx: context.Background(),
    }
    for _, param := range params {
        switch v := param.(type) {
        case *adapter.User:
            s.User = v
        case Permission:
            nq = v
        }
    }
    if s.User == nil && nq {
        app.Panic(http.StatusUnauthorized, adapter.ErrorUserRequired)
    }
    return s
}
