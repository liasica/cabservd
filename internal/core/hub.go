// Copyright (C) liasica. 2022-present.
//
// Created at 2022-10-31
// Based on cabservd by liasica, magicrolan@qq.com.

package core

import (
    "github.com/auroraride/adapter"
    "github.com/auroraride/cabservd/internal/codec"
    "github.com/panjf2000/gnet/v2"
    "go.uber.org/zap"
    "sync"
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
    codec codec.Codec

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
    c.Info("新增客户端连接")

    // 设置连接上下文信息
    conn.SetContext(c)
    return
}

func (h *hub) OnClose(conn gnet.Conn, err error) (action gnet.Action) {
    // 获取客户端
    c, ok := conn.Context().(*Client)
    // 关闭客户端
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
        b, err = h.codec.Decode(conn)

        if err == adapter.ErrorIncompletePacket {
            break
        }

        if err != nil {
            c.Error("消息读取失败", zap.Error(err))
            return
        }

        // 处理消息
        go h.handleMessage(c, b)
    }

    return gnet.None
}

func (h *hub) handleMessage(c *Client, b []byte) {
    // 更新在线状态
    go c.UpdateOnline()

    // 解析数据
    err := h.Bean.OnMessage(b, c)

    if err != nil {
        c.Error("消息解析失败", zap.Error(err))
    }
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
