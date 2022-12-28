// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-10
// Based on cabservd by liasica, magicrolan@qq.com.

package errs

import "errors"

var (
    BadRequest          = errors.New("请求参数错误")
    InternalServerError = errors.New("服务器未知错误")

    UserRequired = errors.New("需要用户信息")

    IncompletePacket = errors.New("incomplete packet") // 数据包不完整

    ParamValidateFailed = errors.New("数据校验失败")

    CabinetSerialRequired     = errors.New("电柜序号不存在")
    CabinetBrandRequired      = errors.New("电柜型号不存在")
    CabinetBinOrdinalRequired = errors.New("仓位序号不存在")
    CabinetNotFound           = errors.New("电柜未找到")
    CabinetOffline            = errors.New("电柜不在线")
    CabinetInitializing       = errors.New("电柜初始化中")
    CabinetAbnormal           = errors.New("电柜状态异常")
    CabinetClientNotFound     = errors.New("未找到在线电柜")
    CabinetNoFully            = errors.New("无可换电池")
    CabinetNoEmpty            = errors.New("无空仓位")
    CabinetBusy               = errors.New("电柜忙")
    CabinetControlParamError  = errors.New("电柜控制参数错误")
    CabinetDoorOpened         = errors.New("有开启中的仓门")

    ExchangeTaskNotExist = errors.New("换电任务不存在")
    ExchangeFailed       = errors.New("换电失败")
    ExchangeTimeOut      = errors.New("换电超时")
    ExchangeBatteryLost  = errors.New("电池未放入")
    ExchangeBatteryExist = errors.New("电池未取走")
)
