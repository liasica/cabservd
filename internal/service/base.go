// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on cabservd by liasica, magicrolan@qq.com.

package service

import (
    "context"
    errs "github.com/auroraride/adapter/errors"
    "github.com/auroraride/adapter/model"
    "github.com/auroraride/cabservd/internal/app"
    "net/http"
)

type Permission bool

const (
    PermissionRequired    Permission = true
    PermissionNotRequired Permission = false
)

type BaseService struct {
    User *model.User
    ctx  context.Context
}

func newService(params ...any) *BaseService {
    nq := PermissionRequired
    s := &BaseService{
        ctx: context.Background(),
    }
    for _, param := range params {
        switch v := param.(type) {
        case *model.User:
            s.User = v
        case Permission:
            nq = v
        }
    }
    if s.User == nil && nq {
        app.Panic(http.StatusUnauthorized, errs.UserRequired)
    }
    return s
}
