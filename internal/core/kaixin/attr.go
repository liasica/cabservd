// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-21
// Based on cabservd by liasica, magicrolan@qq.com.

package kaixin

import (
    "fmt"
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
        open = fmt.Sprintf("%v", a.Value) == DoorStatusOpen
    }
    return
}

func (a *Attr) GetEnable() (enable bool, exists bool) {
    exists = a.ID == SignalBinEnable
    if exists {
        enable = fmt.Sprintf("%v", a.Value) == BinEnable
    }
    return
}

func (a *Attr) GetBatterySN() (sn string, exists bool) {
    exists = a.ID == SignalBatterySN
    if exists {
        sn = fmt.Sprintf("%v", a.Value)
    }
    return
}

func (a *Attr) GetVoltage() (v float64, exists bool) {
    exists = a.ID == SignalBatteryVoltage
    if exists {
        v = a.Value.(float64)
    }
    return
}

func (a *Attr) GetCurrent() (v float64, exists bool) {
    exists = a.ID == SignalBatteryCurrent
    if exists {
        v = a.Value.(float64)
    }
    return
}
