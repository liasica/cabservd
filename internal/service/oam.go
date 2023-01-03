// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-03
// Based on cabservd by liasica, magicrolan@qq.com.

package service

import (
    "context"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/console"
)

type oamService struct {
}

func NewOam() *oamService {
    return &oamService{}
}

func (s *oamService) Business() (running bool) {
    running, _ = ent.Database.Console.Query().Where(console.StatusIn(console.StatusRunning)).Exist(context.Background())
    return
}
