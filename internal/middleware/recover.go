// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on cabservd by liasica, magicrolan@qq.com.

package middleware

import (
    "fmt"
    "github.com/auroraride/adapter/model"
    "github.com/auroraride/cabservd/internal/app"
    "github.com/labstack/echo/v4"
    log "github.com/sirupsen/logrus"
    "runtime/debug"
)

func Recover() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            ctx := app.Context(c)

            defer func() {
                if r := recover(); r != nil {
                    switch v := r.(type) {
                    case *model.Response:
                        _ = ctx.SendResponse(v.Code, v.Message, v.Data)
                    default:
                        err := fmt.Errorf("%v", r)
                        log.Error(fmt.Sprintf("%v\n%s", r, debug.Stack()))
                        c.Error(err)
                    }
                }
            }()
            return next(c)
        }
    }
}
