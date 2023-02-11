// Copyright (C) liasica. 2022-present.
//
// Created at 2022-10-31
// Based on cabservd by liasica, magicrolan@qq.com.

package core

import (
    "github.com/auroraride/adapter"
    "github.com/auroraride/adapter/log"
    "github.com/panjf2000/gnet/v2"
    "go.uber.org/zap"
    "strconv"
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
    zap.L().Info("TCP服务器已启动: " + h.addr)
    return gnet.None
}

func (h *hub) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
    zap.L().Info("新增客户端连接 -> " + c.RemoteAddr().String() + ":" + strconv.Itoa(c.Fd()))

    client := NewClient(c, h)

    // 设置连接上下文信息
    c.SetContext(client)

    return
}

func (h *hub) OnClose(c gnet.Conn, err error) (action gnet.Action) {
    zap.L().Info("客户端断开连接 -> "+c.RemoteAddr().String()+":"+strconv.Itoa(c.Fd()), zap.Error(err))
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
            zap.L().Error("消息读取失败 -> "+c.RemoteAddr().String()+":"+strconv.Itoa(c.Fd()), zap.Error(err))
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
    fields := []zap.Field{
        zap.ByteString("decoded", b),
    }
    var (
        err     error
        message any
    )

    // 记录日志
    defer func() {
        lvl := zap.InfoLevel
        if err != nil {
            lvl = zap.ErrorLevel
            fields = append(fields, zap.Error(err))
        }
        if message != nil {
            fields = append(fields, log.Payload(message))
        }
        zap.L().Log(lvl, "收到消息 <- "+c.RemoteAddr().String()+":"+strconv.Itoa(c.Fd()), fields...)
    }()

    // 更新在线状态
    go c.UpdateOnline()

    // 解析数据
    message, err = h.Bean.OnMessage(b, c)
}
