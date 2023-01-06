// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-06
// Based on cabservd by liasica, magicrolan@qq.com.

package api

import (
    "github.com/auroraride/adapter/async"
    "github.com/auroraride/cabservd/internal/g"
    "github.com/labstack/echo/v4"
    "time"
)

type maintain struct{}

var Maintain = new(maintain)

func (*maintain) Update(echo.Context) (err error) {
    ticker := time.NewTicker(time.Second)
    defer ticker.Stop()

    for ; true; <-ticker.C {
        // 是否有进行中的异步业务
        if async.IsDone() {
            g.Quit <- true
            return
        }
    }
    return nil
}
