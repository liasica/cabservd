// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on cabservd by liasica, magicrolan@qq.com.

package middleware

import (
    "github.com/auroraride/adapter/app"
    "github.com/labstack/echo/v4"
)

func Context() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            return next(app.NewBaseContext(c))
        }
    }
}
