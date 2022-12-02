// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-02
// Based on cabservd by liasica, magicrolan@qq.com.

package types

// DetectBattery 检测电池
type DetectBattery uint8

const (
    DetectBatteryIgnore DetectBattery = iota // 忽略
)
