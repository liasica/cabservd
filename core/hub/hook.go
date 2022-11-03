// Copyright (C) liasica. 2022-present.
//
// Created at 2022-10-31
// Based on cabservd by liasica, magicrolan@qq.com.

package hub

type Hooks interface {
    // OnConnect 连接接口
    OnConnect() error

    // OnRegister 注册
    OnRegister(body []byte) error

    // OnReport 属性上报
    OnReport(body []byte) error

    // OnNotice 告警上报
    OnNotice(body []byte) error

    // OnControl 控制回报
    OnControl(body []byte) error
}
