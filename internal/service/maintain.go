// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-06
// Based on cabservd by liasica, magicrolan@qq.com.

package service

import (
    "github.com/auroraride/cabservd/internal/task"
)

type maintainService struct {
}

func NewMaintain() *maintainService {
    return &maintainService{}
}

func (s *maintainService) Update() bool {
    n := task.Bin.GetListenerCount() + task.Cabinet.GetListenerCount()
    return n == 0
}
