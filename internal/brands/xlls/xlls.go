// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-15
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

import (
	"net/http"

	"github.com/auroraride/adapter"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/color"
	"go.uber.org/zap"

	"github.com/auroraride/cabservd/internal/g"
)

var (
	appID     string
	appSecret []byte
	baseURL   string
	version   string
)

func Start() {
	conf := adapter.GetKoanf()

	appID = conf.Get("xiliulou.appId").(string)
	appSecret = []byte(conf.Get("xiliulou.appSecret").(string))
	baseURL = conf.Get("xiliulou.server").(string)
	version = conf.Get("xiliulou.version").(string)

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	colorer := color.New()
	bind := g.Config.Tcp.Bind
	colorer.Printf("⇨ 西六楼对接启动于 %s\n", colorer.Green(bind))

	r := new(receiver)

	e.POST(pathHardwareOperation, func(c echo.Context) error {
		return nil
	})

	e.POST(pathBusinesss, func(c echo.Context) error {
		return nil
	})

	e.POST(pathOfflineExchange, func(c echo.Context) error {
		return nil
	})

	// 格挡状态变化通知
	e.POST(pathCellChange, r.onBin)

	// 电池和充电器状态变化通知
	e.POST(pathBatteryChargeChange, r.onBat)

	// 柜机状态变化通知
	e.POST(pathCabinetChange, r.onCab)

	e.POST(pathHardwareFault, func(c echo.Context) error {
		return nil
	})

	e.POST(pathSelfServiceOpen, func(c echo.Context) error {
		return nil
	})

	if err := e.Start(bind); err != nil && err != http.ErrServerClosed {
		zap.L().Fatal(err.Error())
	}
}
