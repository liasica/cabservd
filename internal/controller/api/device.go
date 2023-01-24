// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-10
// Based on cabservd by liasica, magicrolan@qq.com.

package api

import (
    "github.com/auroraride/adapter/app"
    "github.com/auroraride/adapter/defs/cabdef"
    "github.com/auroraride/cabservd/internal/service"
    "github.com/labstack/echo/v4"
)

type device struct{}

var Device = new(device)

func (*device) BinInfo(c echo.Context) (err error) {
    ctx, req := app.ContextAndBinding[cabdef.BinInfoRequest](c)
    var info *cabdef.BinInfo
    info, err = service.NewBin(ctx.User).BinInfo(req)
    if err != nil {
        return
    }
    return ctx.SendResponse(info)
}
