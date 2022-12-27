// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-26
// Based on cabservd by liasica, magicrolan@qq.com.

package hook

import (
    "bytes"
    "fmt"
    "github.com/auroraride/cabservd/bridge"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/g"
    "github.com/goccy/go-json"
    "github.com/lib/pq"
    log "github.com/sirupsen/logrus"
    "time"
)

func waitForNotification(l *pq.Listener) {
    type notificationData interface {
        *ent.Cabinet | *ent.Bin
    }

    type data[T notificationData] struct {
        Table  string `json:"table"`
        Action string `json:"action"`
        Data   T      `json:"data"`
    }

    for {
        select {
        case n := <-l.Notify:
            fmt.Println("[EVENTS] 收到数据库变动 channel [", n.Channel, "] :")
            var prettyJSON bytes.Buffer
            _ = json.Indent(&prettyJSON, []byte(n.Extra), "", "  ")
            fmt.Println(string(prettyJSON.Bytes()))

            var (
                serial string
                cab    *ent.Cabinet
                bins   ent.Bins
            )

            switch n.Channel {
            case "bin":
                var d data[*ent.Bin]
                _ = json.Unmarshal([]byte(n.Extra), &d)
                serial = d.Data.Serial
                bins = ent.Bins{d.Data}
            case "cabinet":
                var d data[*ent.Cabinet]
                _ = json.Unmarshal([]byte(n.Extra), &d)
                cab = d.Data
                serial = d.Data.Serial
            }

            bridge.SendCabinet(serial, cab, bins)

            return
        case <-time.After(90 * time.Second):
            log.Info("[EVENTS] 超过90s未检测到PostgreSQL变化, 检查连接...")
            go func() {
                _ = l.Ping()
            }()
            return
        }
    }
}

func ListenPqEvents() {
    dsn := g.Config.Postgres.Dsn

    reportProblem := func(ev pq.ListenerEventType, err error) {
        if err != nil {
            fmt.Println(err.Error())
        }
    }

    listener := pq.NewListener(dsn, 10*time.Second, time.Minute, reportProblem)
    _ = listener.Listen("bin")
    _ = listener.Listen("cabinet")

    log.Println("[EVENTS] 开始监听PostgreSQL变化...")
    for {
        waitForNotification(listener)
    }
}
