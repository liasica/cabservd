// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-26
// Based on cabservd by liasica, magicrolan@qq.com.

package notice

import (
    "fmt"
    "github.com/auroraride/adapter/pn"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/g"
    "github.com/lib/pq"
    log "github.com/sirupsen/logrus"
    "sync"
    "time"
)

type PostgresHook struct {
    // 需要监听的电柜信息
    // cabinets sync.Map

    // 需要监听的仓位信息
    // bins sync.Map

    // chan any => key
    listeners sync.Map
}

func NewPostgres() *PostgresHook {
    return &PostgresHook{}
}

func (h *PostgresHook) GetListenerKey(channel pn.Channel, id uint64) string {
    return fmt.Sprintf("%s-%d", channel, id)
}

func (h *PostgresHook) SetListener(channel pn.Channel, id uint64, l chan any) {
    h.listeners.Store(l, h.GetListenerKey(channel, id))
}

func (h *PostgresHook) GetListener(channel pn.Channel, id uint64) (ch chan any) {
    key := h.GetListenerKey(channel, id)
    h.listeners.Range(func(v, k any) bool {
        if k == key {
            ch = v.(chan any)
            return false
        }
        return true
    })
    return
}

func (h *PostgresHook) DeleteListener(ch chan any) {
    h.listeners.Delete(ch)
}

func (*PostgresHook) ParseData(n *pq.Notification) (channel pn.Channel, data any, err error) {
    channel = pn.Channel(n.Channel)
    err = channel.Validate()
    if err != nil {
        return
    }

    b := []byte(n.Extra)

    switch channel {
    case pn.ChannelCabinet:
        var cab *ent.Cabinet
        cab, err = pn.ParseData[*ent.Cabinet](b)
        data = cab
    case pn.ChannelBin:
        var eb *ent.Bin
        eb, err = pn.ParseData[*ent.Bin](b)
        data = eb
    }
    return
}

func (h *PostgresHook) parseNotice(n *pq.Notification) {
    channel, data, err := h.ParseData(n)
    if err != nil {
        log.Errorf("[EVENTS] 消息解析失败: %v", err)
        return
    }

    var id uint64

    switch target := data.(type) {
    case *ent.Cabinet:
        id = target.ID
        go Aurservd.SendCabinet(false, target.Serial, target, nil)
    case *ent.Bin:
        id = target.ID
        go Aurservd.SendCabinet(false, target.Serial, nil, ent.Bins{target})
        if target.BatterySn != "" {
            go Aurservd.SendBattery(target.BatterySn, target.Serial)
        }
    }

    go h.TrySendNotice(id, channel, data)
}

func (h *PostgresHook) TrySendNotice(id uint64, channel pn.Channel, target any) {
    ch := h.GetListener(channel, id)
    if ch != nil {
        ch <- target
    }
}

func (h *PostgresHook) Start() {
    worker.Add(1)

    dsn := g.Config.Postgres.Dsn

    reportProblem := func(ev pq.ListenerEventType, err error) {
        if err != nil {
            fmt.Println(err.Error())
        }
    }

    l := pq.NewListener(dsn, 10*time.Second, time.Minute, reportProblem)
    _ = l.Listen(pn.ChannelBin.String())
    _ = l.Listen(pn.ChannelCabinet.String())

    log.Println("[EVENTS] 开始监听PostgreSQL变化...")

    worker.Done()

    timeout := time.After(90 * time.Second)

    for {
        select {
        case n := <-l.Notify:
            // fmt.Println("[EVENTS] 收到数据库变动 channel [", n.Channel, "] :")
            // var prettyJSON bytes.Buffer
            // _ = json.Indent(&prettyJSON, []byte(n.Extra), "", "  ")
            // fmt.Println(string(prettyJSON.Bytes()))

            go h.parseNotice(n)

        case <-timeout:
            // log.Info("[EVENTS] 超过90s未检测到PostgreSQL变化, 检查连接...")
            go func() {
                _ = l.Ping()
            }()
        }
    }
}
