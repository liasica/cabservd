// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-01
// Based on cabservd by liasica, magicrolan@qq.com.

package types

type ControlType string

const (
    ControlTypeOpenDoor ControlType = "bin_open"    // 开门请求
    BinDisable                      = "bin_disable" // 禁用
    BinEnable                       = "bin_enable"  // 启用
)

type ControlRequest struct {
    Brand  string      `json:"brand" binding:"required"`  // 电柜品牌
    Type   ControlType `json:"type" binding:"required"`   // 控制类型
    Serial string      `json:"serial" binding:"required"` // 待控制的柜子编号
    Index  *int        `json:"index" binding:"required"`  // 待控制的仓位index
}