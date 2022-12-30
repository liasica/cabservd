// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-30
// Based on cabservd by liasica, magicrolan@qq.com.

package api

import (
    "github.com/auroraride/adapter/model"
    "github.com/auroraride/cabservd/internal/app"
    "github.com/auroraride/cabservd/internal/service"
    "github.com/labstack/echo/v4"
)

type binApi struct{}

var Bin = new(binApi)

func (*binApi) Operate(c echo.Context) (err error) {
    ctx, req := app.ContextAndBinding[model.OperateRequest](c)
    return ctx.SendResponse(service.NewBin(ctx.User).Operate(req))
}
