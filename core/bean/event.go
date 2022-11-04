// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-04
// Based on cabservd by liasica, magicrolan@qq.com.

package bean

type (
    Event interface {
        // OnConnect 连接接口
        OnConnect() (err error)

        // OnMessage 收到消息
        OnMessage(b []byte) (err error)

        // OnRegister 注册
        OnRegister(body []byte) (err error)

        // OnReport 属性上报
        OnReport(body []byte) (err error)

        // OnNotice 告警上报
        OnNotice(body []byte) (err error)

        // OnControl 控制回报
        OnControl(body []byte) (err error)
    }

    Hooks struct{}
)

func (h *Hooks) OnConnect() (err error) {
    return
}

func (h *Hooks) OnMessage(_ []byte) (err error) {
    return
}

func (h *Hooks) OnRegister(_ []byte) (err error) {
    return
}

func (h *Hooks) OnReport(_ []byte) (err error) {
    return
}

func (h *Hooks) OnNotice(_ []byte) (err error) {
    return
}

func (h *Hooks) OnControl(_ []byte) (err error) {
    return
}
