// Copyright (C) liasica. 2022-present.
//
// Created at 2022-10-31
// Based on cabservd by liasica, magicrolan@qq.com.

package core

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/auroraride/adapter"
	"github.com/auroraride/adapter/log"

	"github.com/auroraride/cabservd/internal/g"

	"github.com/panjf2000/gnet/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Hub *hub
)

type hub struct {
	gnet.BuiltinEventEngine

	// address to listen
	addr string

	// 电柜协议
	Bean Hook

	// 编码协议
	codec Codec

	// 在线的客户端
	// serial => *Client
	// serial 在初次连接的时候为空, 当登录成功后是设备的唯一编码
	Clients sync.Map
}

func (h *hub) OnBoot(_ gnet.Engine) (action gnet.Action) {
	zap.L().Info("TCP服务器已启动: " + h.addr)
	return gnet.None
}

func (h *hub) OnOpen(conn gnet.Conn) (out []byte, action gnet.Action) {
	c := NewClient(conn, h)

	if h.Bean.Protocol().Tcp() {
		c.Info("新增客户端连接")
	}

	h.Bean.OnConnect(c)

	// 设置连接上下文信息
	conn.SetContext(c)
	return
}

func (h *hub) OnClose(conn gnet.Conn, err error) (action gnet.Action) {
	// http服务端直接返回关闭, 无须进行客户端在线维护
	if h.Bean.Protocol().Http() {
		return gnet.Close
	}

	// 获取客户端
	c, ok := conn.Context().(*Client)
	// 手动关闭客户端并标记离线
	if ok {
		c.Info("客户端断开连接", zap.Error(err))
		// 停止计时
		c.dead.Stop()

		// 标记电柜为离线
		if c.Serial != "" {
			go c.Offline()

			// 查找并删除客户端
			h.Clients.Delete(c.Serial)
		}
	}
	return
}

func (h *hub) OnTraffic(conn gnet.Conn) (action gnet.Action) {
	// 获取客户端
	c, ok := conn.Context().(*Client)
	if !ok {
		// TODO 关闭连接
		return gnet.Shutdown
	}

	var (
		b   []byte
		err error
	)

	for {
		b, err = h.codec.Decode(c)

		if len(b) > 0 {
			if os.Getenv("LOCAL_DEV") == "true" {
				fmt.Printf("%d\t%s\n", len(b), b)
			}
		}

		if err == adapter.ErrorIncompletePacket {
			break
		}

		if err != nil {
			c.Error("消息读取失败", zap.Error(err))
			return
		}

		// 处理消息
		if len(b) > 0 {
			go h.handleMessage(c, b)
		}
	}

	return
}

func (h *hub) handleMessage(c *Client, b []byte) {
	// TCP需要更新电柜离线判定
	if h.Bean.Protocol().Tcp() {
		go c.UpdateDead()
	}

	// 解析数据
	serial, res, fields, err := h.Bean.OnMessage(c, b)
	lvl := zapcore.InfoLevel

	if err != nil {
		lvl = zapcore.ErrorLevel
		fields = append(fields, zap.Error(err))
	}

	if g.LogBinary {
		fields = append(fields, log.Binary(b))
	}

	c.Log(lvl, "收到消息 ↑ ("+strconv.Itoa(len(b))+" bytes)", fields...)

	// 注册电柜客户端
	if serial != "" {
		c.Serial = serial
		h.register(c)
	}

	// 如果需要发送消息
	if res != nil {
		_ = c.SendMessage(res, 1)
	}

	// 如果是http链接, 处理完成后需要关闭客户端
	if h.Bean.Protocol().Http() {
		_ = c.Conn.Flush()
		_ = c.Conn.Close()
	}
}

// Register 保存设备识别码并注册连接
func (h *hub) register(c *Client) {
	if c.Serial == "" {
		return
	}

	h.Clients.Store(c.Serial, c)
}

// GetClient 获取在线的客户端
func GetClient(serial string) (c *Client, err error) {
	v, exists := Hub.Clients.Load(serial)
	if exists {
		var ok bool
		if c, ok = v.(*Client); ok {
			return
		}
	}

	err = adapter.ErrorCabinetClientNotFound
	return
}
