// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-04
// Based on cabservd by liasica, magicrolan@qq.com.

package kaixin

import (
    "cabservd/core"
    jsoniter "github.com/json-iterator/go"
)

type Hander struct {
    core.Bean
}

func New() *Hander {
    return &Hander{}
}

func (h *Hander) OnMessage(b []byte) (err error) {
    var message Message
    err = jsoniter.Unmarshal(b, &message)
    if err != nil {
        return
    }
    switch message.MessageType {
    case MessageTypeLoginRequest:
    }
    return
}
