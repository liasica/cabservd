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
}
