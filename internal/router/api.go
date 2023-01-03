// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-11
// Based on cabservd by liasica, magicrolan@qq.com.

package router

import (
    "fmt"
    "github.com/auroraride/adapter"
    "github.com/auroraride/cabservd/internal/app"
    "github.com/auroraride/cabservd/internal/controller/api"
    "github.com/auroraride/cabservd/internal/g"
    mw "github.com/auroraride/cabservd/internal/middleware"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    log "github.com/sirupsen/logrus"
    "net/http"
)

func Start() {
    r := echo.New()

    r.HTTPErrorHandler = func(err error, c echo.Context) {
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

    r.Validator = app.NewValidator()

    r.Use(
        mw.Context(),
        mw.Recover(),
        mw.User(),

        middleware.GzipWithConfig(middleware.GzipConfig{
            Level: 5,
        }),
        // TODO body dump middleware
    )

    r.POST("/operate/bin", api.Operate.Bin)

    r.POST("/business/usable", api.Business.Usable)
    r.POST("/business/do", api.Business.Do)

    r.POST("/exchange/usable", api.Exchange.Usable)
    r.POST("/exchange/do", api.Exchange.Do)

    // operation and maintenance 运维接口
    r.GET("oam/business", api.Oam.Business)

    log.Fatal(r.Start(g.Config.Api.Bind))
}
