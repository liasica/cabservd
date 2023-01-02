// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-02
// Based on cabservd by liasica, magicrolan@qq.com.

package types

import "github.com/auroraride/cabservd/internal/ent"

// CabinetBinInfo 电柜仓位情况
type CabinetBinInfo struct {
    FullyBin *ent.Bin // 选取的满电仓位
    EmptyBin *ent.Bin // 选取的空电仓位
    Fully    int      // 正常的满电仓位数量
    Empty    int      // 正常的空电仓位数量
    Abnormal int      // 故障数量
    Disabled int      // 禁用数量
    Opened   int      // 开仓数量
}
