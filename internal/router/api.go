// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-11
// Based on cabservd by liasica, magicrolan@qq.com.

package router

import (
	"net/http"

	"github.com/auroraride/adapter"
	"github.com/auroraride/adapter/app"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/auroraride/cabservd/assets"
	"github.com/auroraride/cabservd/internal/controller/api"
	"github.com/auroraride/cabservd/internal/g"
	"github.com/auroraride/cabservd/internal/middleware"
)

func Start(e *echo.Echo, brand adapter.CabinetBrand) {
	e.Renderer = assets.Templates

	e.GET("/maintain/clients", api.Maintain.Clients)

	// 仓位操作 <管理员权限>
	e.POST("/operate/bin", api.Operate.Bin, app.UserTypeManagerMiddleware(), middleware.BinOperateExclusive())
	e.POST("/bin/deactivate", api.Bin.Deactivate, app.UserTypeManagerMiddleware(), middleware.BinOperateExclusive())

	// 代理操作
	e.POST("/agent/operate/bin", api.Operate.Bin, app.UserTypeAgentMiddleware(), middleware.BinOperateExclusive())

	// 业务
	e.POST("/business/usable", api.Business.Usable)
	e.POST("/business/do", api.Business.Do, middleware.BinOperateExclusive())

	// 换电
	e.POST("/exchange/usable", api.Exchange.Usable)
	e.POST("/exchange/do", api.Exchange.Do, middleware.BinOperateExclusive())

	// 电柜信息
	e.POST("/device/bininfo", api.Device.BinInfo)

	// 西六楼电柜
	if brand == adapter.CabinetBrandXiliulouServer {
		e.Any("/battery", api.Xllscab.Battery)
	}

	if err := e.Start(g.Config.Api.Bind); err != nil && err != http.ErrServerClosed {
		zap.L().Fatal(err.Error())
	}
}
