// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-28
// Based on cabservd by liasica, magicrolan@qq.com.

package ent

// ResetBatteryInfo 无电池的时候清除电池信息
// TODO: 是否有必要?
func (u *BinUpsert) ResetBatteryInfo() *BinUpsert {
    u.SetCurrent(0).SetVoltage(0).SetSoc(0).SetSoh(0)
    return u
}
