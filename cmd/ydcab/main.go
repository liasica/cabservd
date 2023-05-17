// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-18
// Based on cabservd by liasica, magicrolan@qq.com.

package main

import (
	"github.com/auroraride/cabservd/internal"
	"github.com/auroraride/cabservd/internal/brands/yundong"
	"github.com/auroraride/cabservd/internal/core"
	"github.com/auroraride/cabservd/internal/g"
)

func main() {
	// 设定变量
	g.Fakevoltage = 45

	internal.Boot(func() {
		hook, codecor := yundong.New()
		// 启动socket hub
		go core.Start(
			g.Config.Tcp.Bind,
			hook,
			codecor,
		)
	})
}
