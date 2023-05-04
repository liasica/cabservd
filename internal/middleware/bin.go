// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-04
// Based on cabservd by liasica, magicrolan@qq.com.

package middleware

import (
	"github.com/auroraride/adapter"
	"github.com/labstack/echo/v4"

	"github.com/auroraride/cabservd/internal/service"
)

func BinOperateExclusive() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 获取电柜编码
			serial := c.Get("serial").(string)
			if serial == "" {
				return adapter.ErrorCabinetSerialRequired
			}
			// 是否有进行中的任务
			if service.NewConsole().InJob(serial) {
				return adapter.ErrorCabinetBusy
			}
			return next(c)
		}
	}
}
