// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-31
// Based on cabservd by liasica, magicrolan@qq.com.

package service

type operateService struct {
    *BaseService
}

func NewOperate(params ...any) *operateService {
    return &operateService{
        BaseService: newService(params...),
    }
}
