// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-31
// Based on cabservd by liasica, magicrolan@qq.com.

package service

import "github.com/auroraride/adapter"

type operateService struct {
    *BaseService
}

func NewOperate(params ...any) *operateService {
    return &operateService{
        BaseService: newService(params...),
    }
}

func (s *operateService) Do(req *adapter.OperateRequest) error {
    switch req.Type {
    case adapter.OperateBinOpen, adapter.OperateBinEnable, adapter.OperateBinDisable:
        return NewBin(s.ctx, s.User).Operate(req)
    }
    return adapter.ErrorBadRequest
}
