// Copyright (C) liasica. 2023-present.
//
// Created at 2023-03-09
// Based on cabservd by liasica, magicrolan@qq.com.

package kaixin

import (
	"github.com/auroraride/cabservd/internal/codec"
	"github.com/auroraride/cabservd/internal/core"
	"github.com/auroraride/cabservd/internal/tower"
)

func New() (core.Hook, codec.Codec) {
	return tower.New(
		tower.WithBinSignals(binSignals),
		tower.WithAutoResetBattery(false),
		tower.WithBatteryReign(true),
	), &codec.HeaderLength{}
}

func NewNonIntelligent() (core.Hook, codec.Codec) {
	return tower.New(
		tower.WithBinSignals(binSignals),
		tower.WithAutoResetBattery(false),
		tower.WithBatteryReign(true),
		tower.WithFakeVoltage(44), // 2023年05月08日20:09:08 曹博文说把虚拟电压调整为44V
	), &codec.HeaderLength{}
}
