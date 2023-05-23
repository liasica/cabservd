// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-04
// Based on cabservd by liasica, magicrolan@qq.com.

package main

import (
	"github.com/auroraride/cabservd/internal"
	"github.com/auroraride/cabservd/internal/brands/kaixin"
	"github.com/auroraride/cabservd/internal/core"
	"github.com/auroraride/cabservd/internal/g"
)

func main() {
	g.ExchangeFirstStepRetryTimes = 3
	g.ExchangeThirdStepRetryTimes = 3

	internal.Boot(func() {
		hook, codecor := kaixin.New()
		// 启动socket hub
		go core.Start(
			g.Config.Tcp.Bind,
			hook,
			codecor,
		)
	})
}
