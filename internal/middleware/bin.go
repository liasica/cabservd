// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-04
// Based on cabservd by liasica, magicrolan@qq.com.

package middleware

import (
	"bytes"
	"io"

	"github.com/auroraride/adapter"
	"github.com/auroraride/adapter/app"
	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"

	"github.com/auroraride/cabservd/internal/service"
)

func BinOperateExclusive() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := app.Context(c)
			var req struct {
				Serial string `json:"serial"`
			}
			// 获取电柜编码
			b, _ := io.ReadAll(c.Request().Body)
			_ = jsoniter.Unmarshal(b, &req)
			_ = ctx.Request().Body.Close()
			if req.Serial == "" {
				return adapter.ErrorCabinetSerialRequired
			}

			ctx.Request().Body = io.NopCloser(bytes.NewBuffer(b))

			// 是否有进行中的任务
			if service.NewConsole(ctx.User).InJob(req.Serial) {
				return adapter.ErrorCabinetBusy
			}
			return next(ctx)
		}
	}
}
