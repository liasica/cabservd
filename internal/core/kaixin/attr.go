// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-21
// Based on cabservd by liasica, magicrolan@qq.com.

package kaixin

import (
    log "github.com/sirupsen/logrus"
    "strconv"
)

// Attr 属性信息
type Attr struct {
    SignalData
    DoorID string `json:"doorId,omitempty"` // 柜门ID (可为空)
}

func (attr *Attr) GetBrand() (string, bool) {
    return Brand
}

func (attr *Attr) GetSN() (string, bool) {
    return Brand
}

func (attr *Attr) GetOpen() (bool, bool) {
    return true
}

func (attr *Attr) GetDoorIndex() (index int, exists bool) {
    if attr.DoorID == "" {
        return
    }

    exists = true
    id, err := strconv.Atoi(attr.DoorID)
    if err != nil {
        log.Errorf("仓位解析失败")
    }
    index = id - 1
    return
}
