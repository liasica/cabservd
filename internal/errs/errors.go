// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-10
// Based on cabservd by liasica, magicrolan@qq.com.

package errs

import "errors"

var (
    IncompletePacket        = errors.New("incomplete packet") // 数据包不完整
    ClientNotFound          = errors.New("未找到客户端")
    CabinetDeviceIDRequired = errors.New("电柜序号不存在")
    CabinetBinIndexRequired = errors.New("仓位序号不存在")
)
