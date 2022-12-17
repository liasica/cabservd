// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-02
// Based on cabservd by liasica, magicrolan@qq.com.

package types

import (
    "database/sql/driver"
)

// DetectBattery 检测电池
type DetectBattery uint8

const (
    DetectBatteryIgnore DetectBattery = iota // 忽略
    DetectBatteryPutin                       // 检测放入
    DetectBatteryPutout                      // 检测取走
)

// ExchangeStep 换电步骤
type ExchangeStep uint8

const (
    ExchangeStepFirst ExchangeStep = iota + 1
    ExchangeStepSecond
    ExchangeStepThird
    ExchangeStepFourth
)

func (s *ExchangeStep) Scan(src interface{}) error {
    switch src := src.(type) {
    case nil:
        return nil
    case int64:
        *s = ExchangeStep(src)
    }
    return nil
}

func (s ExchangeStep) Value() (driver.Value, error) {
    return s, nil
}

func (s ExchangeStep) String() string {
    switch s {
    case ExchangeStepFirst:
        return "第1步, 开启空电仓门"
    case ExchangeStepSecond:
        return "第2步, 放入电池关仓"
    case ExchangeStepThird:
        return "第3步, 开启满电仓门"
    case ExchangeStepFourth:
        return "第4步, 取出电池关仓"
    }
    return "-"
}

type ExchangeRequest struct {
    User   *User  `json:"user" form:"user" binding:"required"` // 用户信息
    Serial string `json:"serial" form:"serial"`
}
