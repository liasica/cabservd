// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-11
// Based on cabservd by liasica, magicrolan@qq.com.

package router

import (
    amw "github.com/auroraride/adapter/middleware"
    "github.com/auroraride/adapter/zlog"
    "github.com/auroraride/cabservd/assets"
    "github.com/auroraride/cabservd/internal/controller/api"
    "github.com/auroraride/cabservd/internal/g"
    mw "github.com/auroraride/cabservd/internal/middleware"
    "github.com/labstack/echo/v4"
    "net/http"
)

func Start(e *echo.Echo) {
    e.Renderer = assets.Templates

    dumpFile := amw.NewDumpLoggerMiddleware(zlog.StandardLogger())

    // 运维接口
    m := e.Group("/maintain")
    m.GET("/update", api.Maintain.Update)
    m.GET("/clients", api.Maintain.Clients)

    r := e.Group("/")
    r.Use(
        mw.Context(),
        mw.Recover(),
        mw.User(),

        dumpFile.WithDefaultConfig(),

        // middleware.GzipWithConfig(middleware.GzipConfig{
        //     Level: 5,
        // }),
    )

    // 仓位操作 <管理员权限>
    r.POST("operate/bin", api.Operate.Bin, mw.Manager())

    r.POST("business/usable", api.Business.Usable)
    r.POST("business/do", api.Business.Do)

    r.POST("exchange/usable", api.Exchange.Usable)
    r.POST("exchange/do", api.Exchange.Do)

    r.POST("device/bininfo", api.Device.BinInfo)

    if err := e.Start(g.Config.Api.Bind); err != nil && err != http.ErrServerClosed {
        zlog.Fatal(err.Error())
    }
}
