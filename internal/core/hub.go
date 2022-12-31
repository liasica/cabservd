// Copyright (C) liasica. 2022-present.
//
// Created at 2022-10-31
// Based on cabservd by liasica, magicrolan@qq.com.

package core

import (
    "github.com/auroraride/adapter"
    "github.com/panjf2000/gnet/v2"
    log "github.com/sirupsen/logrus"
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
    brand adapter.Brand

    // 电柜协议
    Bean Hook

    // 编码协议
    codec Codec

    // 在线的客户端
    // *Client => serial
    // serial 在初次连接的时候为空, 当登录成功后是设备的唯一编码
    clients sync.Map

    // 客户端发起连接
    connect chan *Client

    // 断开客户端连接
    disconnect chan *Client
}

func (h *hub) OnBoot(_ gnet.Engine) (action gnet.Action) {
    log.Infof("TCP服务器已启动 %s", h.addr)
    return gnet.None
}

func (h *hub) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
    log.Infof("[FD=%d / %s] 新增客户端连接", c.Fd(), c.RemoteAddr())

    // 设置连接上下文信息
    ctx := NewClient(c, h)
    c.SetContext(ctx)

    // 注册连接
    h.connect <- ctx

    return
}

func (h *hub) OnClose(c gnet.Conn, err error) (action gnet.Action) {
    log.Infof("[FD=%d / %s] 客户端断开连接, error?: %v", c.Fd(), c.RemoteAddr(), err)
    // 获取客户端
    client, ok := c.Context().(*Client)
    // 删除客户端
    if ok {
        client.Close()
        h.clients.Delete(client)
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

        if err == adapter.IncompletePacket {
            break
        }

        if err != nil {
            log.Errorf("[FD=%d / %s] 消息读取失败, err: %v", c.Fd(), c.RemoteAddr(), err)
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

func (h *hub) handleMessage(b []byte, client *Client) {
    // 记录日志
    log.Infof("[FD=%d / %s] 接收到消息, message: %s", client.Fd(), client.RemoteAddr(), b)

    // 解析
    // TODO 未知的 Client
    err := h.Bean.OnMessage(b, client)
    if err != nil {
        log.Errorf("[FD=%d / %s] 解析失败, err: %v, 原始消息: %s", client.Fd(), client.RemoteAddr(), err, b)
    }
}
