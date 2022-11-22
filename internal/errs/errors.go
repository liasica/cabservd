// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-10
// Based on cabservd by liasica, magicrolan@qq.com.

package errs

import "errors"

var (
    IncompletePacket = errors.New("incomplete packet") // 数据包不完整
    ClientNotFound   = errors.New("未找到客户端")

    ParamValidateFailed = errors.New("数据校验失败")

    CabinetDeviceIDRequired = errors.New("电柜序号不存在")
    CabinetBinIndexRequired = errors.New("仓位序号不存在")
    CabinetNotFound         = errors.New("电柜未找到")
    CabinetNoFully          = errors.New("无可换电池")
    CabinetNoEmpty          = errors.New("无空仓位")
    CabinetBusy             = errors.New("电柜忙")

    ExchangeTaskNotExist = errors.New("换电任务不存在")
    ExchangeFailed       = errors.New("换电失败")
    ExchangeTimeOut      = errors.New("换电超时")
    ExchangeBatteryLost  = errors.New("电池未放入")
    ExchangeBatteryExist = errors.New("电池未取走")
)
