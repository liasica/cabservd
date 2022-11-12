// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-06
// Based on cabservd by liasica, magicrolan@qq.com.

package core

import (
    "github.com/auroraride/cabservd/internal/errs"
    jsoniter "github.com/json-iterator/go"
    "github.com/panjf2000/gnet/v2"
    log "github.com/sirupsen/logrus"
)

type Client struct {
    // gnet 连接
    gnet.Conn

    Hub *hub
}

// SetDeviceID 设置deviceID
func (c *Client) SetDeviceID(id string) {
    c.Hub.clients.Store(c, id)
}

// SendMessage 向客户端发送消息
func (c *Client) SendMessage(data any, params ...any) (err error) {
    // return jsoniter.NewEncoder(c).Encode(c)

    b, _ := jsoniter.Marshal(data)

    var logMessage bool
    if len(params) > 0 {
        logMessage = params[0].(bool)
    }

    _, err = c.Write(append(b, newline...))
    if err != nil {
        log.Errorf("[FD=%d / %s] 发送失败, message: %s", c.Fd(), c.RemoteAddr(), b)
    } else if logMessage {
        log.Infof("[FD=%d / %s] 发送消息, message: %s", c.Fd(), c.RemoteAddr(), b)
    }

    return
}

// GetClient 获取在线的客户端
func GetClient(devId string) (c *Client, err error) {
    Hub.clients.Range(func(key, value any) bool {
        client, _ := key.(*Client)
        sn, _ := value.(string)
        if sn == devId {
            c = client
            return false
        }
        return true
    })
    if c == nil {
        err = errs.ClientNotFound
    }
    return
}
