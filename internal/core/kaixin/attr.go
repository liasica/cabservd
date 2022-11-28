// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-21
// Based on cabservd by liasica, magicrolan@qq.com.

package kaixin

import (
	"github.com/auroraride/cabservd/internal/core/types"
	log "github.com/sirupsen/logrus"
	"strconv"
)

// Attr 属性信息
type Attr struct {
	SignalData
	DoorID string `json:"doorId,omitempty"` // 柜门ID (可为空)
}

func (a *Attr) GetDoorIndex() (index int, exists bool) {
	if a.DoorID == "" {
		return
	}

	exists = true
	id, err := strconv.Atoi(a.DoorID)
	if err != nil {
		log.Errorf("仓位解析失败")
	}
	index = id - 1
	return
}

func (a *Attr) GetOpen() (open bool, exists bool) {
	exists = a.ID == SignalDoorStatus
	if exists {
		open = a.StringValue() == DoorStatusOpen
	}
	return
}

func (a *Attr) GetEnable() (enable bool, exists bool) {
	exists = a.ID == SignalBinEnable
	if exists {
		enable = a.StringValue() == BinEnable
	}
	return
}

func (a *Attr) GetBattery() (v string, exists bool) {
	exists = a.ID == SignalBatterySN
	if exists {
		v = a.StringValue()
	}
	return
}

func (a *Attr) GetChargeStatus() (v types.ChargeStatus, exists bool) {
	exists = a.ID == SignalBinChargeStatus
	if exists {
		v = types.ChargeStatus(a.StringValue())
	}
	return
}

func (a *Attr) GetVoltage() (v float64, exists bool) {
	exists = a.ID == SignalBatteryVoltage
	if exists {
		v = a.ValueFloat64()
	}
	return
}

func (a *Attr) GetCurrent() (v float64, exists bool) {
	exists = a.ID == SignalBatteryCurrent
	if exists {
		v = a.ValueFloat64()
	}
	return
}

func (a *Attr) GetSoC() (v float64, exists bool) {
	exists = a.ID == SignalSOC
	if exists {
		v = a.ValueFloat64()
	}
	return
}

func (a *Attr) GetSoH() (v float64, exists bool) {
	exists = a.ID == SignalSOH
	if exists {
		v = a.ValueFloat64()
	}
	return
}
