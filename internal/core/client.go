// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-06
// Based on cabservd by liasica, magicrolan@qq.com.

package core

import (
    "context"
    "github.com/auroraride/adapter"
    "github.com/auroraride/adapter/log"
    "github.com/auroraride/adapter/snag"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/cabinet"
    "github.com/auroraride/cabservd/internal/g"
    jsoniter "github.com/json-iterator/go"
    "github.com/panjf2000/gnet/v2"
    "go.uber.org/zap"
    "time"
)

type Client struct {
    // gnet 连接
    gnet.Conn

    Hub *hub

    // 电柜编号
    Serial string

    // 上次接收消息时间
    dead *time.Timer
}

type MessageProxy struct {
    Data   []byte
    Client *Client
}

func NewClient(conn gnet.Conn, h *hub) *Client {
    c := &Client{
        Conn: conn,
        Hub:  h,
    }
    c.dead = time.AfterFunc(20*time.Minute, func() {
        c.Offline()
    })
    return c
}

// SendMessage 向客户端发送消息
func (c *Client) SendMessage(message any, savelog bool) (err error) {
    b, _ := jsoniter.Marshal(message)

    data := c.Hub.codec.Encode(b)

    defer func() {
        fields := []zap.Field{
            log.ResponseBody(b),
        }

        if savelog || g.Config.Environment.IsDevelopment() {
            level := zap.InfoLevel
            if err != nil {
                level = zap.ErrorLevel
                fields = append(fields, zap.Error(err), log.Binary(b))
            }
            c.Log(level, "发送消息 ↓ ", fields...)
        }
    }()

    // // TODO DEMO
    // if len(params) > 1 {
    //     x := []byte(fmt.Sprintf(`{"msgType":500,"txnNo":%d,"devId":"CH6004KXHD220728222","paramList":[{"id":"02301001","value":"04","doorId":"7"}]}`, time.Now().UnixMilli()))
    //     data = append(data, c.Hub.codec.Encode(x)...)
    //     fmt.Printf("%x", data)
    // }

    _, err = c.Write(data)

    return
}

// GetClient 获取在线的客户端
func GetClient(serial string) (c *Client, err error) {
    v, exists := Hub.Clients.Load(serial)
    if exists {
        var ok bool
        c, ok = v.(*Client)
        if ok {
            return
        }
    }

    err = adapter.ErrorCabinetClientNotFound

    return
}

// AfterClose 关闭电柜客户端
func (c *Client) AfterClose() {
    snag.WithRecover(func() {
        // 停止计时
        c.dead.Stop()

        // 标记电柜为离线
        if c.Serial != "" {
            go c.Offline()
        }

        // 查找并删除客户端
        c.Hub.Clients.Delete(c.Serial)
    })
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
