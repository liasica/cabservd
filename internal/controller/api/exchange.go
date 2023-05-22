// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on cabservd by liasica, magicrolan@qq.com.

package api

import (
	"github.com/auroraride/adapter"
	"github.com/auroraride/adapter/app"
	"github.com/auroraride/adapter/async"
	"github.com/auroraride/adapter/defs/cabdef"
	"github.com/auroraride/adapter/maintain"
	"github.com/auroraride/adapter/snag"
	"github.com/labstack/echo/v4"

	"github.com/auroraride/cabservd/internal/service"
)

type exchange struct{}

var Exchange = new(exchange)

func (*exchange) Usable(c echo.Context) (err error) {
	if maintain.Exists() {
		app.Panic(adapter.ErrorMaintain)
	}

	ctx, req := app.ContextAndBinding[cabdef.ExchangeUsableRequest](c)
	return ctx.SendResponse(service.NewExchange(ctx.User).Usable(req))
}

func (*exchange) Do(c echo.Context) (err error) {
	if maintain.Exists() {
		app.Panic(adapter.ErrorMaintain)
	}

	return async.WithTaskReturn[error](func() error {
		return snag.WithRecoverError(func() error {
			ctx, req := app.ContextAndBinding[cabdef.ExchangeRequest](c)
			return ctx.SendResponse(service.NewExchange(ctx.User).Do(req))
		})
	})
}
