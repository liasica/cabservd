// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-31
// Based on cabservd by liasica, magicrolan@qq.com.

package api

import (
    "github.com/auroraride/adapter"
    "github.com/auroraride/cabservd/internal/app"
    "github.com/auroraride/cabservd/internal/service"
    "github.com/labstack/echo/v4"
)

type operate struct{}

var Operate = new(operate)

func (*operate) Do(c echo.Context) (err error) {
    ctx, req := app.ContextAndBinding[adapter.OperateRequest](c)
    return ctx.SendResponse(service.NewOperate(ctx.User).Do(req))
}
