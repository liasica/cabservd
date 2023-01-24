// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-31
// Based on cabservd by liasica, magicrolan@qq.com.

package api

import (
    "github.com/auroraride/adapter/app"
    "github.com/auroraride/adapter/defs/cabdef"
    "github.com/auroraride/cabservd/internal/service"
    "github.com/labstack/echo/v4"
)

type operate struct{}

var Operate = new(operate)

func (*operate) Bin(c echo.Context) (err error) {
    ctx, req := app.ContextAndBinding[cabdef.OperateBinRequest](c)
    return ctx.SendResponse(service.NewOperate(ctx.User).Bin(req))
}
