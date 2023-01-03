// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on cabservd by liasica, magicrolan@qq.com.

package middleware

import (
    "github.com/auroraride/adapter"
    "github.com/auroraride/cabservd/internal/app"
    "github.com/labstack/echo/v4"
    "net/http"
)

var (
    userSkipper = map[string]bool{
        "/oam/running": true,
        "/oam/status":  true,
    }
)

func User() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            ctx := app.Context(c)
            p := c.Path()
            if userSkipper[p] {
                return next(ctx)
            }

            header := c.Request().Header
            // 获取user信息
            user := &adapter.User{
                ID:   header.Get(adapter.HeaderUserID),
                Type: adapter.UserType(header.Get(adapter.HeaderUserType)),
            }
            if user.ID == "" || user.Type == "" {
                app.Panic(http.StatusUnauthorized, adapter.ErrorUserRequired)
            }
            ctx.User = user
            return next(ctx)
        }
    }
}
