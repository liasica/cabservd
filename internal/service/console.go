// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on cabservd by liasica, magicrolan@qq.com.

package service

import (
    "github.com/auroraride/adapter/app"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/console"
    "time"
)

type consoleService struct {
    *app.BaseService

    orm *ent.ConsoleClient
}

func NewConsole(params ...any) *consoleService {
    return &consoleService{
        BaseService: app.NewService(params...),
        orm:         ent.Database.Console,
    }
}

// Update 更新记录
func (s *consoleService) Update(ec *ent.Console, b *ent.Bin, err error) *ent.Console {
    now := time.Now()
    cr := ec.Update().SetStopAt(now)
    if ec.StartAt != nil {
        cr.SetDuration(now.Sub(*ec.StartAt).Seconds())
    }

    // 仓位信息
    if b != nil {
        cr.SetAfterBin(b.Info())
    }

    if err != nil {
        cr.SetStatus(console.StatusFailed).SetMessage(err.Error())
    } else {
        cr.SetStatus(console.StatusSuccess)
    }

    ec, _ = cr.Save(s.GetContext())
    return ec
}
