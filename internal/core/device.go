// Copyright (C) liasica. 2023-present.
//
// Created at 2023-03-12
// Based on cabservd by liasica, magicrolan@qq.com.

package core

type Device struct {
    BatteryReign              bool // 是否有电池在位检测
    AutoResetWithoutBatterySN bool // 当电池编号不存在的时候是否自动清除电池信息 (soc / soh / current / voltage)
}
