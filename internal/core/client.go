// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-06
// Based on cabservd by liasica, magicrolan@qq.com.

package core

import (
    "context"
    "github.com/auroraride/adapter/log"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/cabinet"
    jsoniter "github.com/json-iterator/go"
    "github.com/panjf2000/gnet/v2"
    "go.uber.org/zap"
    "time"
)

type Client struct {
    // gnet 连接
    gnet.Conn

    Hub *hub

    // 电柜编号
    Serial string

    // 上次接收消息时间
    dead *time.Timer
}

func NewClient(conn gnet.Conn, h *hub) *Client {
    c := &Client{
        Conn: conn,
        Hub:  h,
    }
    c.dead = time.AfterFunc(20*time.Minute, func() {
        _ = c.Conn.Close()
    })
    return c
}

// SendMessage 向客户端发送消息
func (c *Client) SendMessage(message any) (err error) {
    b, _ := jsoniter.Marshal(message)

    data := c.Hub.codec.Encode(b)

    defer func() {
        fields := []zap.Field{
            log.ResponseBody(b),
        }

        level := zap.InfoLevel
        if err != nil {
            level = zap.ErrorLevel
            fields = append(fields, zap.Error(err), log.Binary(b))
        }
        c.Log(level, "发送消息 ↓ ", fields...)
    }()

    _, err = c.Write(data)

    return
}

// Offline 标记电柜离线
func (c *Client) Offline() {
    if c.Serial == "" {
        return
    }
    // TODO 是否发送消息
    _ = ent.Database.Cabinet.Update().Where(cabinet.Serial(c.Serial)).SetOnline(false).Exec(context.Background())
}

// UpdateOnline 更新电柜离线时间
func (c *Client) UpdateOnline() {
    c.dead.Reset(20 * time.Minute)
}
