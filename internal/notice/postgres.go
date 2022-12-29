// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-26
// Based on cabservd by liasica, magicrolan@qq.com.

package notice

import (
    "bytes"
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
    PostgresChannelCabinet = "cabinet"
    PostgresChannelBin     = "bin"
)

type PostgresHook struct {
    // 需要监听的电柜信息
    cabinets sync.Map

    // 需要监听的仓位信息
    bins sync.Map
}

func (h *PostgresHook) SetCabinetListener(id uint64, ch chan *ent.Cabinet) {
    h.cabinets.Store(id, ch)
}

func (h *PostgresHook) DeleteCabinetListener(id uint64) {
    v, ok := h.cabinets.LoadAndDelete(id)
    if ok {
        close(v.(chan *ent.Cabinet))
    }
}

func (h *PostgresHook) SetBinListener(id uint64, ch chan *ent.Bin) {
    h.bins.Store(id, ch)
}

func (h *PostgresHook) DeleteBinListener(id uint64) {
    v, ok := h.bins.LoadAndDelete(id)
    if ok {
        close(v.(chan *ent.Bin))
    }
}

func NewPostgres() *PostgresHook {
    return &PostgresHook{}
}

type PostgresNoticeData[T any] struct {
    Table  string `json:"table"`
    Action string `json:"action"`
    Data   *T     `json:"data"`
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

    for {
        select {
        case n := <-l.Notify:
            fmt.Println("[EVENTS] 收到数据库变动 channel [", n.Channel, "] :")
            var prettyJSON bytes.Buffer
            _ = json.Indent(&prettyJSON, []byte(n.Extra), "", "  ")
            fmt.Println(string(prettyJSON.Bytes()))

            switch n.Channel {
            case PostgresChannelCabinet:
                v := new(PostgresNoticeData[ent.Cabinet])
                err := json.Unmarshal([]byte(n.Extra), v)
                if err != nil {
                    log.Errorf("[EVENTS] 消息解析失败: %v", err)
                    continue
                }

                cab := v.Data

                // 发送aurservd同步数据
                go SendCabinet(cab.Serial, cab, nil)

                // 发送监听数据
                if v, ok := h.cabinets.Load(cab.ID); ok {
                    go func() { v.(chan *ent.Cabinet) <- cab }()
                }
            case PostgresChannelBin:
                v := new(PostgresNoticeData[ent.Bin])
                err := json.Unmarshal([]byte(n.Extra), v)
                if err != nil {
                    log.Errorf("[EVENTS] 消息解析失败: %v", err)
                    continue
                }

                b := v.Data

                // 发送aurservd同步数据
                go SendCabinet(b.Serial, nil, ent.Bins{b})

                // 发送监听数据
                if v, ok := h.bins.Load(b.ID); ok {
                    go func() { v.(chan *ent.Bin) <- b }()
                }
            }

        case <-time.After(90 * time.Second):
            // log.Info("[EVENTS] 超过90s未检测到PostgreSQL变化, 检查连接...")
            go func() {
                _ = l.Ping()
            }()
        }
    }
}
