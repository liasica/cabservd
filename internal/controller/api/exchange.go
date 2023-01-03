// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on cabservd by liasica, magicrolan@qq.com.

package api

import (
    "github.com/auroraride/adapter/defs/cabdef"
    "github.com/auroraride/cabservd/internal/app"
    "github.com/auroraride/cabservd/internal/service"
    "github.com/labstack/echo/v4"
)

type exchange struct{}

var Exchange = new(exchange)

func (*exchange) Usable(c echo.Context) (err error) {
    ctx, req := app.ContextAndBinding[cabdef.ExchangeUsableRequest](c)
    return ctx.SendResponse(service.NewExchange(ctx.User).Usable(req))
}

func (*exchange) Do(c echo.Context) (err error) {
    ctx, req := app.ContextAndBinding[cabdef.ExchangeRequest](c)
    return ctx.SendResponse(service.NewExchange(ctx.User).Do(req))
}
