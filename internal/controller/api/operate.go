// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-31
// Based on cabservd by liasica, magicrolan@qq.com.

package api

import (
    "github.com/auroraride/cabservd/internal/app"
    "github.com/labstack/echo/v4"
)

type operate struct{}

var Operate = new(operate)

func (*operate) Bin(c echo.Context) (err error) {
    ctx := app.Context(c)
    return ctx.SendResponse()
}
