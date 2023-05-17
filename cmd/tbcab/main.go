// Copyright (C) liasica. 2023-present.
//
// Created at 2023-03-09
// Based on cabservd by liasica, magicrolan@qq.com.

package main

import (
	"github.com/auroraride/cabservd/internal"
	"github.com/auroraride/cabservd/internal/core"
	"github.com/auroraride/cabservd/internal/g"
	"github.com/auroraride/cabservd/internal/tower"
)

func main() {
	// 设定变量
	g.Fakevoltage = 40
	g.AutoResetWithoutBatterySN = true
	g.CalculateMonVoltage = true

	internal.Boot(
		func() {
			// 启动socket hub
			go core.Start(
				g.Config.Tcp.Bind,
				tower.New(
					tower.WithMessageTypeList(&tower.MessageTypeList{
						LoginRequest:    110,
						LoginResponse:   111,
						ReportRequest:   310,
						ReportResponse:  311,
						NoticeRequest:   410,
						NoticeResponse:  411,
						ControlRequest:  500,
						ControlResponse: 501,
					}),
				),
				&core.Linebreak{},
			)
		},
	)
}
