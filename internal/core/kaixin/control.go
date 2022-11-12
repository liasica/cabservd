// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-11
// Based on cabservd by liasica, magicrolan@qq.com.

package kaixin

import (
    "github.com/auroraride/cabservd/internal/core"
    "time"
)

func Control(deviceId string, req ControlRequest) (err error) {
    msg := &Request{
        Message: Message{
            MsgType: MessageTypeControlRequest,
            TxnNo:   time.Now().UnixMilli(),
            DevID:   deviceId,
        },
        ControlRequest: req,
    }

    var c *core.Client
    c, err = core.GetClient(deviceId)
    if err != nil {
        return
    }

    return c.SendMessage(msg, true)
}
