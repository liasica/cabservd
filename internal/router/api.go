// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-11
// Based on cabservd by liasica, magicrolan@qq.com.

package router

import (
    "fmt"
    "github.com/auroraride/adapter"
    amw "github.com/auroraride/adapter/middleware"
    "github.com/auroraride/cabservd/assets"
    "github.com/auroraride/cabservd/internal/app"
    "github.com/auroraride/cabservd/internal/controller/api"
    "github.com/auroraride/cabservd/internal/g"
    mw "github.com/auroraride/cabservd/internal/middleware"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    log "github.com/sirupsen/logrus"
    "net/http"
)

func Start(e *echo.Echo) {
    e.JSONSerializer = &adapter.DefaultJSONSerializer{}

    e.HTTPErrorHandler = func(err error, c echo.Context) {
        ctx := app.Context(c)
        message := err
        code := http.StatusInternalServerError
        var data any
        switch err.(type) {
        case *echo.HTTPError:
            target := err.(*echo.HTTPError)
            message = fmt.Errorf("%v", target.Message)
            break
        }
        _ = ctx.SendResponse(code, message, data)
    }

    echo.NotFoundHandler = func(c echo.Context) error {
        return app.Context(c).SendResponse(http.StatusNotFound, adapter.ErrorNotFound)
    }

    echo.MethodNotAllowedHandler = func(c echo.Context) error {
        routerAllowMethods, ok := c.Get(echo.ContextKeyHeaderAllow).(string)
        if ok && routerAllowMethods != "" {
            c.Response().Header().Set(echo.HeaderAllow, routerAllowMethods)
        }
        return app.Context(c).SendResponse(http.StatusBadRequest, fmt.Errorf("%v", echo.ErrMethodNotAllowed.Message))
    }

    e.Validator = app.NewValidator()

    e.Renderer = assets.Templates

    dumpFile := amw.NewDumpFile()

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

        middleware.GzipWithConfig(middleware.GzipConfig{
            Level: 5,
        }),
        // TODO body dump middleware
    )

    // 仓位操作 <管理员权限>
    r.POST("operate/bin", api.Operate.Bin, mw.Manager())

    r.POST("business/usable", api.Business.Usable)
    r.POST("business/do", api.Business.Do)

    r.POST("exchange/usable", api.Exchange.Usable)
    r.POST("exchange/do", api.Exchange.Do)

    r.POST("device/bininfo", api.Device.BinInfo)

    if err := e.Start(g.Config.Api.Bind); err != nil && err != http.ErrServerClosed {
        log.Fatal(err)
    }
}
