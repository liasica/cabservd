// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-26
// Based on cabservd by liasica, magicrolan@qq.com.

package hook

import (
    "fmt"
    "github.com/auroraride/cabservd/bridge"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/g"
    jsoniter "github.com/json-iterator/go"
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
            // fmt.Println("Received data from channel [", n.Channel, "] :")
            // // Prepare notification payload for pretty print
            // var prettyJSON bytes.Buffer
            // err := json.Indent(&prettyJSON, []byte(n.Extra), "", "  ")
            // if err != nil {
            //     fmt.Println("Error processing JSON: ", err)
            //     return
            // }
            // fmt.Println(string(prettyJSON.Bytes()))

            var (
                serial string
                cab    *ent.Cabinet
                bins   ent.Bins
            )

            switch n.Channel {
            case "bin":
                var d data[*ent.Bin]
                _ = jsoniter.Unmarshal([]byte(n.Extra), &d)
                serial = d.Data.Serial
                bins = ent.Bins{d.Data}
            case "cabinet":
                var d data[*ent.Cabinet]
                _ = jsoniter.Unmarshal([]byte(n.Extra), &d)
                cab = d.Data
                serial = d.Data.Serial
            }

            bridge.SendCabinet(serial, cab, bins)

            return
        case <-time.After(90 * time.Second):
            // Received no events for 90 seconds, checking connection
            log.Info("[EVENTS] Received no events for 90 seconds, checking connection")
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

    log.Println("[EVENTS] Start monitoring PostgreSQL...")
    for {
        waitForNotification(listener)
    }
}
