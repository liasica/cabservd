// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-06
// Based on cabservd by liasica, magicrolan@qq.com.

package api

import (
    "github.com/auroraride/cabservd/internal/service"
    "github.com/labstack/echo/v4"
)

type maintain struct{}

var Maintain = new(maintain)

func (*maintain) Update(c echo.Context) (err error) {
    service.NewMaintain().Update()
    return nil
}
