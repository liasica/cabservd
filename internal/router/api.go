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
    "net/http"
)

func Start(e *echo.Echo) {
    e.Renderer = assets.Templates

    // 仓位操作 <管理员权限>
    e.POST("/operate/bin", api.Operate.Bin, app.UserTypeManagerMiddleware())

    e.POST("/business/usable", api.Business.Usable)
    e.POST("/business/do", api.Business.Do)

    e.POST("/exchange/usable", api.Exchange.Usable)
    e.POST("/exchange/do", api.Exchange.Do)

    e.POST("device/bininfo", api.Device.BinInfo)

    if err := e.Start(g.Config.Api.Bind); err != nil && err != http.ErrServerClosed {
        zlog.Fatal(err.Error())
    }
}
