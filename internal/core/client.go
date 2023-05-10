// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-06
// Based on cabservd by liasica, magicrolan@qq.com.

package core

import (
	"context"
	"time"

	"github.com/panjf2000/gnet/v2"
	"go.uber.org/zap"

	"github.com/auroraride/cabservd/internal/codec"
	"github.com/auroraride/cabservd/internal/ent"
	"github.com/auroraride/cabservd/internal/ent/cabinet"
	"github.com/auroraride/cabservd/internal/g"
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
	var dd time.Duration = 20
	if g.Config.DeadDuration > 0 {
		dd = time.Duration(g.Config.DeadDuration)
	}
	c.dead = time.AfterFunc(dd*time.Minute, func() {
		_ = c.Conn.Close()
	})
	return c
}

type ResponseMessenger interface {
	GetMessage(c codec.Codec) ([]byte, []zap.Field)
}

// SendMessage 向客户端发送消息
func (c *Client) SendMessage(messenger ResponseMessenger) (err error) {
	b, fields := messenger.GetMessage(c.Hub.codec)
	_, err = c.Write(b)
	lvl := zap.InfoLevel
	if err != nil {
		lvl = zap.ErrorLevel
		fields = append(fields, zap.Error(err))
	}
	c.Log(lvl, "发送消息 ↓ ", fields...)

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
