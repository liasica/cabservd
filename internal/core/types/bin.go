// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-28
// Based on cabservd by liasica, magicrolan@qq.com.

package types

type ChargeStatus string

const (
    ChargeStatusNoBattery ChargeStatus = "0" // 无电池
    ChargeStatusCharging  ChargeStatus = "1" // 充电中
    ChargeStatusFull      ChargeStatus = "2" // 已充满
    ChargeStatusException ChargeStatus = "5" // 异常
)
