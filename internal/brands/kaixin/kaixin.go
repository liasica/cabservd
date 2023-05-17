// Copyright (C) liasica. 2023-present.
//
// Created at 2023-03-09
// Based on cabservd by liasica, magicrolan@qq.com.

package kaixin

import (
	"github.com/auroraride/cabservd/internal/core"
	"github.com/auroraride/cabservd/internal/g"
	"github.com/auroraride/cabservd/internal/tower"
)

func New() (core.Hook, core.Codec) {
	// 设定变量
	g.Fakevoltage = 40
	g.BatteryReign = true

	return tower.New(
		tower.WithBinSignals(binSignals),
	), &core.HeaderLength{}
}

func NewNonIntelligent() (core.Hook, core.Codec) {
	// 设定变量
	g.Fakevoltage = 44 // 2023年05月08日20:09:08 曹博文说把虚拟电压调整为44V
	g.BatteryReign = true

	return tower.New(
		tower.WithBinSignals(binSignals),
	), &core.HeaderLength{}
}
