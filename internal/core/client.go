// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-06
// Based on cabservd by liasica, magicrolan@qq.com.

package core

import (
    "github.com/auroraride/cabservd/internal/errs"
    jsoniter "github.com/json-iterator/go"
    "github.com/panjf2000/gnet/v2"
    log "github.com/sirupsen/logrus"
)

type Client struct {
    // gnet 连接
    gnet.Conn

    Hub *hub

    // 消息代理
    receiver chan *MessageProxy
}

type MessageProxy struct {
    Data   []byte
    Client *Client
}

func NewClient(conn gnet.Conn, h *hub) *Client {
    c := &Client{
        Conn:     conn,
        Hub:      h,
        receiver: make(chan *MessageProxy),
    }
    go c.run()
    return c
}

// 启动客户端任务
func (c *Client) run() {
    for {
        select {
        case message := <-c.receiver:
            // 消息代理
            c.Hub.handleMessage(message.Data, message.Client)
        }
    }
}

// SetDeviceID 设置deviceID
func (c *Client) SetDeviceID(id string) {
    c.Hub.clients.Store(c, id)
}

// SendMessage 向客户端发送消息
func (c *Client) SendMessage(message any, params ...any) (err error) {
    // return jsoniter.NewEncoder(c).Encode(c)

    b, _ := jsoniter.Marshal(message)

    var logMessage bool
    if len(params) > 0 {
        logMessage = params[0].(bool)
    }

    data := c.Hub.codec.Encode(b)

    // // TODO DEMO
    // if len(params) > 1 {
    //     x := []byte(fmt.Sprintf(`{"msgType":500,"txnNo":%d,"devId":"CH6004KXHD220728222","paramList":[{"id":"02301001","value":"04","doorId":"7"}]}`, time.Now().UnixMilli()))
    //     data = append(data, c.Hub.codec.Encode(x)...)
    //     fmt.Printf("%x", data)
    // }

    _, err = c.Write(data)
    if err != nil {
        log.Errorf("[FD=%d / %s] 发送失败, message: %s", c.Fd(), c.RemoteAddr(), b)
    } else if logMessage {
        log.Infof("[FD=%d / %s] 发送消息, message: %s", c.Fd(), c.RemoteAddr(), b)
    }

    return
}

// GetClient 获取在线的客户端
func GetClient(devId string) (c *Client, err error) {
    Hub.clients.Range(func(key, value any) bool {
        client, _ := key.(*Client)
        sn, _ := value.(string)
        if sn == devId {
            c = client
            return false
        }
        return true
    })
    if c == nil {
        err = errs.CabinetClientNotFound
    }
    return
}
