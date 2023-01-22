// Copyright (C) liasica. 2022-present.
//
// Created at 2022-10-31
// Based on cabservd by liasica, magicrolan@qq.com.

package core

import (
    "github.com/auroraride/adapter"
    "github.com/auroraride/adapter/zlog"
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

    // 电柜品牌
    brand adapter.CabinetBrand

    // 电柜协议
    Bean Hook

    // 编码协议
    codec Codec

    // 在线的客户端
    // serial => *Client
    // serial 在初次连接的时候为空, 当登录成功后是设备的唯一编码
    Clients *sync.Map
}

func (h *hub) OnBoot(_ gnet.Engine) (action gnet.Action) {
    zlog.Info("TCP服务器已启动: " + h.addr)
    return gnet.None
}

func (h *hub) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
    zlog.Info("新增客户端连接", zap.Int("FD", c.Fd()), zap.String("address", c.RemoteAddr().String()))

    client := NewClient(c, h)

    // 设置连接上下文信息
    c.SetContext(client)

    return
}

func (h *hub) OnClose(c gnet.Conn, err error) (action gnet.Action) {
    zlog.Info("客户端断开连接", zap.Error(err), zap.Int("FD", c.Fd()), zap.String("address", c.RemoteAddr().String()))
    // 获取客户端
    client, ok := c.Context().(*Client)
    // 关闭客户端
    if ok {
        go client.Close()
    }
    return
}

func (h *hub) OnTraffic(c gnet.Conn) (action gnet.Action) {
    // 获取客户端
    client, ok := c.Context().(*Client)
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

        if err == adapter.ErrorIncompletePacket {
            break
        }

        if err != nil {
            zlog.Error("消息读取失败", zap.Error(err), zap.Int("FD", c.Fd()), zap.String("address", c.RemoteAddr().String()))
            return
        }

        // 使用channel处理消息体
        client.receiver <- &MessageProxy{
            Data:   b,
            Client: client,
        }
    }

    return gnet.None
}

func (h *hub) handleMessage(b []byte, c *Client) {
    // 记录日志
    zlog.Info("收到消息 ↑", zap.Int("FD", c.Fd()), zap.String("address", c.RemoteAddr().String()), zap.ByteString("payload", b))

    // 更新在线状态
    go c.UpdateOnline()

    // 解析数据
    err := h.Bean.OnMessage(b, c)
    if err != nil {
        zlog.Error("消息解析失败", zap.Error(err), zap.Int("FD", c.Fd()), zap.String("address", c.RemoteAddr().String()), zap.ByteString("payload", b))
    }
}
