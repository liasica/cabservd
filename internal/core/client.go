// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-06
// Based on cabservd by liasica, magicrolan@qq.com.

package core

import (
	"context"
	"net"
	"strconv"
	"time"

	"github.com/auroraride/adapter/log"
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

	address net.Addr
}

func NewClient(conn gnet.Conn, h *hub) *Client {
	c := &Client{
		Conn:    conn,
		Hub:     h,
		address: conn.RemoteAddr(),
	}

	if h.Bean.Protocol().Tcp() {
		var dd time.Duration = 20
		if g.Config.DeadDuration > 0 {
			dd = time.Duration(g.Config.DeadDuration)
		}
		c.dead = time.AfterFunc(dd*time.Minute, func() {
			_ = c.Conn.Close()
		})
	}
	return c
}

type ResponseMessenger interface {
	GetMessage(c codec.Codec) ([]byte, []zap.Field)
}

// SendMessage 向客户端发送消息
func (c *Client) SendMessage(messenger ResponseMessenger, times int) (err error) {
	b, fields := messenger.GetMessage(c.Hub.codec)
	var n int
	n, err = c.Write(b)
	lvl := zap.InfoLevel
	if err != nil {
		lvl = zap.ErrorLevel
		fields = append(fields, zap.Error(err))
	}

	// 记录原始消息
	if g.LogBinary {
		fields = append(fields, log.Binary(b))
	}

	msg := "发送消息 ↓ (" + strconv.Itoa(n) + " bytes)"

	if times > 1 {
		// msg = "「第" + strconv.Itoa(times) + "次重试」 " + msg
		fields = append(fields, zap.Int("retry", times))
	}

	c.Log(lvl, msg, fields...)

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

// UpdateDead 更新电柜离线判定
func (c *Client) UpdateDead() {
	c.dead.Reset(20 * time.Minute)
}
