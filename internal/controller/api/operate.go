// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-31
// Based on cabservd by liasica, magicrolan@qq.com.

package api

import (
	"github.com/auroraride/adapter/app"
	"github.com/auroraride/adapter/defs/cabdef"
	"github.com/labstack/echo/v4"

	"github.com/auroraride/cabservd/internal/service"
)

type operate struct{}

var Operate = new(operate)

func (*operate) Bin(c echo.Context) (err error) {
	ctx, req := app.ContextAndBinding[cabdef.OperateBinRequest](c)
	return ctx.SendResponse(service.NewOperate(ctx.User).Bin(req))
}

func (*operate) BinOpenAndClose(c echo.Context) (err error) {
	ctx, req := app.ContextAndBinding[cabdef.OperateBinRequest](c)
	return ctx.SendResponse(service.NewOperate(ctx.User).BinOpenAndClose(req))
}
