// Copyright (C) liasica. 2022-present.
//
// Created at 2022-10-31
// Based on cabservd by liasica, magicrolan@qq.com.

package core

import (
    "bufio"
    "bytes"
    "fmt"
    "github.com/panjf2000/gnet/v2"
    log "github.com/sirupsen/logrus"
    "sync"
)

type Hub struct {
    gnet.BuiltinEventEngine

    // address to listen
    addr string

    // 电柜类型
    cate string

    // 电柜协议
    bean Hook

    // 在线的客户端
    // *Client => deviceID
    // deviceID 在初次连接的时候为空, 当登录成功后是设备的唯一编码
    clients sync.Map

    // 客户端发起连接
    connect chan *Client

    // 断开客户端连接
    disconnect chan *Client
}

func (h *Hub) OnBoot(_ gnet.Engine) (action gnet.Action) {
    log.Infof("TCP服务器已启动 %s", h.addr)
    return gnet.None
}

func (h *Hub) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
    log.Infof("[FD=%d] 新增客户端连接, address: %s", c.Fd(), c.RemoteAddr())

    // 设置连接上下文信息
    ctx := &Client{
        Conn: c,
        Hub:  h,
    }
    c.SetContext(ctx)

    // 注册连接
    h.connect <- ctx

    return
}

func (h *Hub) OnClose(c gnet.Conn, err error) (action gnet.Action) {
    log.Infof("[FD=%d] 客户端断开连接, address: %s, error?: %v", c.Fd(), c.RemoteAddr(), err)
    return
}

func (h *Hub) OnTraffic(c gnet.Conn) (action gnet.Action) {
    // 获取客户端
    client, ok := c.Context().(*Client)
    if !ok {
        // TODO 关闭连接
        return gnet.Shutdown
    }

    // 读取消息
    reader := bufio.NewReader(c)
    var buffer bytes.Buffer

    for {
        b, prefix, err := reader.ReadLine()

        if err != nil {
            log.Errorf("[Hub] 客户端消息读取失败: %v", err)
        }

        buffer.Write(b)

        // 是否有后续消息
        if prefix {
            fmt.Println("接收后续消息")
            continue
        }

        // 读取成功
        if buffer.Len() > 0 {
            break
        }
    }

    b := buffer.Bytes()
    buffer.Reset()

    // 记录日志
    log.Infof("[FD=%d] 接收到消息, address: %s, message: %s", c.Fd(), c.RemoteAddr(), b)

    // 解析
    // TODO 未知的 Client
    err := h.bean.OnMessage(b, client)
    if err != nil {
        log.Errorf("[FD=%d] 解析失败, address: %s, err: %v, 原始消息: %s", c.Fd(), c.RemoteAddr(), err, b)
    }

    return gnet.None
}
