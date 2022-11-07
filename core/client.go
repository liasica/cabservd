// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-06
// Based on cabservd by liasica, magicrolan@qq.com.

package core

import (
    jsoniter "github.com/json-iterator/go"
    "github.com/panjf2000/gnet/v2"
)

var (
    newline = []byte("\n")
)

type Client struct {
    // gnet 连接
    gnet.Conn

    Hub *Hub
}

// SetDeviceID 设置deviceID
func (c *Client) SetDeviceID(id string) {
    c.Hub.clients.Store(c, id)
}

// SendMessage 向客户端发送消息
func (c *Client) SendMessage(data any) (err error) {
    var b []byte
    b, err = jsoniter.Marshal(data)
    if err != nil {
        return
    }
    _, err = c.Write(b)
    _, err = c.Write(newline)
    return
}
