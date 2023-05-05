// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-04
// Based on cabservd by liasica, magicrolan@qq.com.

package api

import (
	"github.com/auroraride/adapter/app"
	"github.com/auroraride/adapter/defs/cabdef"
	"github.com/labstack/echo/v4"

	"github.com/auroraride/cabservd/internal/service"
)

type bin struct{}

var Bin = new(bin)

// Deactivate 仓位逻辑禁用或启用
func (*bin) Deactivate(c echo.Context) (err error) {
	ctx, req := app.ContextAndBinding[cabdef.BinDeactivateRequest](c)
	return ctx.SendResponse(map[string]bool{"status": service.NewBin(ctx.User).Deactivate(req) == nil})
}
