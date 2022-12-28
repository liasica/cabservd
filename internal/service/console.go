// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on cabservd by liasica, magicrolan@qq.com.

package service

type consoleService struct {
    *BaseService
}

func NewConsole(params ...any) *consoleService {
    return &consoleService{
        BaseService: newService(params...),
    }
}
