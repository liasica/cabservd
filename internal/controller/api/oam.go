// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-03
// Based on cabservd by liasica, magicrolan@qq.com.

package api

import (
    "github.com/auroraride/cabservd/internal/service"
    "github.com/labstack/echo/v4"
    "net/http"
)

type oam struct{}

var Oam = new(oam)

func (*oam) Business(c echo.Context) (err error) {
    return c.JSON(http.StatusOK, service.NewOam().Business())
}
