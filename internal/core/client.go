// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-06
// Based on cabservd by liasica, magicrolan@qq.com.

package core

import (
    "context"
    "github.com/auroraride/adapter"
    "github.com/auroraride/adapter/snag"
    "github.com/auroraride/adapter/zlog"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/cabinet"
    "github.com/google/uuid"
    jsoniter "github.com/json-iterator/go"
    "github.com/panjf2000/gnet/v2"
    "go.uber.org/zap"
    "time"
)

type Client struct {
    ID uuid.UUID

    // gnet 连接
    gnet.Conn

    Hub *hub

    // 电柜编号
    Serial string

    // 消息代理
    receiver chan *MessageProxy

    // 上次接收消息时间
    dead *time.Timer
}

type MessageProxy struct {
    Data   []byte
    Client *Client
}

func NewClient(conn gnet.Conn, h *hub) *Client {
    c := &Client{
        ID:       uuid.New(),
        Conn:     conn,
        Hub:      h,
        receiver: make(chan *MessageProxy),
    }
    c.dead = time.AfterFunc(20*time.Minute, func() {
        c.Offline()
    })
    go c.run()
    return c
}

// 启动客户端任务
func (c *Client) run() {
    for {
        select {
        case message := <-c.receiver:
            if message == nil {
                return
            }
            // 消息代理
            c.Hub.handleMessage(message.Data, message.Client)
        }
    }
}

// SendMessage 向客户端发送消息
// params[0]: 是否记录消息
func (c *Client) SendMessage(message any) (err error) {
    b, _ := jsoniter.Marshal(message)

    data := c.Hub.codec.Encode(b)

    // // TODO DEMO
    // if len(params) > 1 {
    //     x := []byte(fmt.Sprintf(`{"msgType":500,"txnNo":%d,"devId":"CH6004KXHD220728222","paramList":[{"id":"02301001","value":"04","doorId":"7"}]}`, time.Now().UnixMilli()))
    //     data = append(data, c.Hub.codec.Encode(x)...)
    //     fmt.Printf("%x", data)
    // }

    _, err = c.Write(data)
    if err != nil {
        zlog.Error("消息发送失败", zap.Error(err), zap.Int("FD", c.Fd()), zap.String("address", c.RemoteAddr().String()), zap.Binary("payload", data))
    } else {
        zlog.Info("发送消息 ↓", zap.Int("FD", c.Fd()), zap.String("address", c.RemoteAddr().String()), zap.Binary("payload", data))
    }

    return
}

// GetClient 获取在线的客户端
func GetClient(devId string) (c *Client, err error) {
    Hub.Clients.Range(func(key, value any) bool {
        client, _ := value.(*Client)
        sn, _ := key.(string)
        if sn == devId {
            c = client
            return false
        }
        return true
    })
    if c == nil {
        err = adapter.ErrorCabinetClientNotFound
    }
    return
}

// Close 关闭电柜客户端
func (c *Client) Close() {
    snag.WithRecover(func() {
        // 标记电柜为离线
        if c.Serial != "" {
            go c.Offline()
        }

        close(c.receiver)

        // 查找并删除客户端
        c.Hub.Clients.Range(func(k, v any) bool {
            if client, ok := v.(*Client); ok && client.ID == c.ID {
                c.Hub.Clients.Delete(k)
                return false
            }
            return true
        })

    }, zlog.StandardLogger())
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

// Register 保存设备识别码并注册连接
func (c *Client) Register(serial string) {
    if serial != "" {
        c.Serial = serial
        c.Hub.Clients.Store(serial, c)
    }
}
