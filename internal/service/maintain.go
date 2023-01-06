// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-06
// Based on cabservd by liasica, magicrolan@qq.com.

package service

import (
    "github.com/auroraride/cabservd/internal/g"
    "github.com/auroraride/cabservd/internal/task"
    "time"
)

type maintainService struct {
}

func NewMaintain() *maintainService {
    return &maintainService{}
}

func (s *maintainService) Update() {
    ticker := time.NewTicker(time.Second)
    defer ticker.Stop()

    for ; true; <-ticker.C {
        n := task.Bin.GetListenerCount() + task.Cabinet.GetListenerCount()
        if n == 0 {
            g.Quit <- true
            return
        }
    }
}
