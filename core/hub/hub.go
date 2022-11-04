// Copyright (C) liasica. 2022-present.
//
// Created at 2022-10-31
// Based on cabservd by liasica, magicrolan@qq.com.

package hub

import (
    "bufio"
    "bytes"
    "cabservd/core/bean"
    "fmt"
    "github.com/panjf2000/gnet/v2"
    log "github.com/sirupsen/logrus"
    "sync"
)

type hub struct {
    gnet.BuiltinEventEngine

    // address to listen
    addr string

    // 电柜协议
    bean bean.Bean

    // 在线的客户端
    // key = devId
    // value = Client
    clients sync.Map
}

func (h *hub) OnBoot(_ gnet.Engine) (action gnet.Action) {
    log.Infof("TCP服务器已启动 %s", h.addr)
    return gnet.None
}

func (h *hub) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
    log.Infof("新增客户端连接: %d, address: %s", c.Fd(), c.RemoteAddr())
    return
}

func (h *hub) OnClose(c gnet.Conn, err error) (action gnet.Action) {
    log.Infof("客户端断开连接 %d, error?: %v", c.Fd(), err)
    return
}

func (h *hub) OnTraffic(c gnet.Conn) (action gnet.Action) {
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

    // TODO 解析
    log.Println(buffer.String())

    buffer.Reset()

    return gnet.None
}
