// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-01
// Based on cabservd by liasica, magicrolan@qq.com.

package internal

import (
    "context"
    "github.com/auroraride/adapter/codec"
    "github.com/auroraride/adapter/maintain"
    "github.com/auroraride/adapter/snag"
    "github.com/auroraride/adapter/zlog"
    "github.com/auroraride/cabservd/internal/brands/kaixin"
    "github.com/auroraride/cabservd/internal/core"
    "github.com/auroraride/cabservd/internal/demo"
    "github.com/auroraride/cabservd/internal/g"
    "github.com/auroraride/cabservd/internal/router"
    "github.com/auroraride/cabservd/internal/task"
    "github.com/labstack/echo/v4"
)

func Boot() {
    // 初始化系统
    initialize()

    snag.WithRecover(func() {
        // 加载hooks
        task.Start()

        // 启动 http server
        e := echo.New()
        go router.Start(e)

        // 启动socket hub
        go core.Start(
            g.Config.Tcp.Bind,
            g.Config.Brand,
            kaixin.New(),
            &codec.HeaderLength{},
        )

        // debug
        go demo.Debug()

        // maintain
        if maintain.Exists() {
            _ = maintain.Remove()
        }

        select {
        case <-g.Quit:
            _ = e.Shutdown(context.Background())
        }

    }, zlog.StandardLogger())
}
