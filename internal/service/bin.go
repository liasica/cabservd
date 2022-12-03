// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-03
// Based on cabservd by liasica, magicrolan@qq.com.

package service

import "github.com/auroraride/cabservd/internal/types"

type binService struct {
}

func NewBin() *binService {
    return &binService{}
}

func (s *binService) SetEnable(req *types.ControlRequest) {
}
