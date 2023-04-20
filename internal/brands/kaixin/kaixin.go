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
	), &codec.HeaderLength{}
}
