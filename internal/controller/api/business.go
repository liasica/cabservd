// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-02
// Based on cabservd by liasica, magicrolan@qq.com.

package api

import (
	"github.com/auroraride/adapter"
	"github.com/auroraride/adapter/app"
	"github.com/auroraride/adapter/async"
	"github.com/auroraride/adapter/defs/cabdef"
	"github.com/auroraride/adapter/maintain"
	"github.com/auroraride/cabservd/internal/service"
	"github.com/labstack/echo/v4"
)

type business struct{}

var Business = new(business)

func (*business) Usable(c echo.Context) (err error) {
	if maintain.Exists() {
		app.Panic(adapter.ErrorMaintain)
	}

	ctx, req := app.ContextAndBinding[cabdef.BusinuessUsableRequest](c)
	return ctx.SendResponse(service.NewBusiness(ctx.User).Usable(req))
}

func (*business) Do(c echo.Context) (err error) {
	if maintain.Exists() {
		app.Panic(adapter.ErrorMaintain)
	}

	return async.WithTaskReturn[error](func() error {
		ctx, req := app.ContextAndBinding[cabdef.BusinessRequest](c)
		return ctx.SendResponse(service.NewBusiness(ctx.User).Do(req))
	})
}
