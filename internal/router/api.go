// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-11
// Based on cabservd by liasica, magicrolan@qq.com.

package router

import (
    "github.com/auroraride/adapter/app"
    "github.com/auroraride/adapter/zlog"
    "github.com/auroraride/cabservd/assets"
    "github.com/auroraride/cabservd/internal/controller/api"
    "github.com/auroraride/cabservd/internal/g"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "net/http"
)

func Start(e *echo.Echo) {
    e.Renderer = assets.Templates

    dump := app.NewDumpLoggerMiddleware(zlog.StandardLogger())

    // 运维接口
    m := e.Group("/maintain")
    m.GET("/update", api.Maintain.Update)
    m.GET("/clients", api.Maintain.Clients)

    userSkipper := map[string]bool{
        "/oam/running": true,
        "/oam/status":  true,
    }

    r := e.Group("/")
    r.Use(
        app.ContextMiddleware(),
        app.RecoverMiddleware(zlog.StandardLogger()),
        app.UserMiddleware(func(c echo.Context) bool {
            return userSkipper[c.Path()]
        }),

        dump.WithDefaultConfig(),

        middleware.GzipWithConfig(middleware.GzipConfig{
            Level: 5,
        }),
    )

    // 仓位操作 <管理员权限>
    r.POST("operate/bin", api.Operate.Bin, app.UserTypeManagerMiddleware())

    r.POST("business/usable", api.Business.Usable)
    r.POST("business/do", api.Business.Do)

    r.POST("exchange/usable", api.Exchange.Usable)
    r.POST("exchange/do", api.Exchange.Do)

    r.POST("device/bininfo", api.Device.BinInfo)

    if err := e.Start(g.Config.Api.Bind); err != nil && err != http.ErrServerClosed {
        zlog.Fatal(err.Error())
    }
}
