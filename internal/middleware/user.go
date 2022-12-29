// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on cabservd by liasica, magicrolan@qq.com.

package middleware

import (
    errs "github.com/auroraride/adapter/errors"
    "github.com/auroraride/adapter/model"
    "github.com/auroraride/cabservd/internal/app"
    "github.com/labstack/echo/v4"
    "net/http"
)

func User() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            ctx := app.Context(c)
            header := c.Request().Header
            // 获取user信息
            user := &model.User{
                ID:   header.Get(model.HeaderUserID),
                Type: model.UserType(header.Get(model.HeaderUserType)),
            }
            if user.ID == "" || user.Type == "" {
                app.Panic(http.StatusUnauthorized, errs.UserRequired)
            }
            ctx.User = user
            return next(ctx)
        }
    }
}
