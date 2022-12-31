// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-26
// Based on cabservd by liasica, magicrolan@qq.com.

package notice

import (
    "fmt"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/g"
    "github.com/goccy/go-json"
    "github.com/lib/pq"
    log "github.com/sirupsen/logrus"
    "sync"
    "time"
)

const (
    PostgresChannelCabinet PostgresChannel = "cabinet"
    PostgresChannelBin     PostgresChannel = "bin"
)

type PostgresChannel string

func (c PostgresChannel) String() string {
    return string(c)
}

type PostgresHook struct {
    // 需要监听的电柜信息
    // cabinets sync.Map

    // 需要监听的仓位信息
    // bins sync.Map

    // chan any => key
    listeners sync.Map
}

func (h *PostgresHook) GetListenerKey(channel PostgresChannel, id uint64) string {
    return fmt.Sprintf("%s-%d", channel, id)
}

func (h *PostgresHook) SetListener(channel PostgresChannel, id uint64, l chan IDSerialGetter) {
    h.listeners.Store(l, h.GetListenerKey(channel, id))
}

func (h *PostgresHook) GetListener(channel PostgresChannel, id uint64) (ch chan IDSerialGetter) {
    key := h.GetListenerKey(channel, id)
    h.listeners.Range(func(v, k any) bool {
        if k == key {
            ch = v.(chan IDSerialGetter)
            return false
        }
        return true
    })
    return
}

func (h *PostgresHook) DeleteListener(ch chan IDSerialGetter) {
    h.listeners.Delete(ch)
}

// func (h *PostgresHook) SetCabinetListener(id uint64, ch chan *ent.Cabinet) {
//     h.cabinets.Store(id, ch)
// }

// func (h *PostgresHook) DeleteCabinetListener(id uint64) {
//     v, ok := h.cabinets.LoadAndDelete(id)
//     if ok {
//         close(v.(chan *ent.Cabinet))
//     }
// }

// func (h *PostgresHook) SetBinListener(id uint64, ch chan *ent.Bin) {
//     h.bins.Store(id, ch)
// }

// func (h *PostgresHook) DeleteBinListener(id uint64) {
//     v, ok := h.bins.LoadAndDelete(id)
//     if ok {
//         close(v.(chan *ent.Bin))
//     }
// }

func NewPostgres() *PostgresHook {
    return &PostgresHook{}
}

type IDSerialGetter interface {
    GetID() uint64
    GetSerial() string
}

type PostgresNoticeData[T IDSerialGetter] struct {
    Table  string `json:"table"`
    Action string `json:"action"`
    Data   T      `json:"data"`
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
    _ = l.Listen("bin")
    _ = l.Listen("cabinet")

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

            var (
                cab    *ent.Cabinet
                bins   ent.Bins
                cn     = PostgresChannel(n.Channel)
                target IDSerialGetter
            )

            switch cn {
            case PostgresChannelCabinet:
                v := new(PostgresNoticeData[*ent.Cabinet])
                err := json.Unmarshal([]byte(n.Extra), v)
                if err != nil {
                    log.Errorf("[EVENTS] 消息解析失败: %v", err)
                    continue
                }

                target = v.Data
                cab = v.Data

            case PostgresChannelBin:
                v := new(PostgresNoticeData[*ent.Bin])
                err := json.Unmarshal([]byte(n.Extra), v)
                if err != nil {
                    log.Errorf("[EVENTS] 消息解析失败: %v", err)
                    continue
                }

                target = v.Data
                bins = ent.Bins{v.Data}

            default:
                return
            }

            go h.SendNotice(cn, target)
            go Aurservd.SendCabinet(false, target.GetSerial(), cab, bins)

        case <-timeout:
            // log.Info("[EVENTS] 超过90s未检测到PostgreSQL变化, 检查连接...")
            go func() {
                _ = l.Ping()
            }()
        }
    }
}

func (h *PostgresHook) SendNotice(channel PostgresChannel, target IDSerialGetter) {
    id := target.GetID()
    ch := h.GetListener(channel, id)
    if ch != nil {
        ch <- target
    }
}
