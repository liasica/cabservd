// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on cabservd by liasica, magicrolan@qq.com.

package task

import (
    "github.com/auroraride/adapter/pkg/loki"
    "github.com/auroraride/adapter/pqm"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/g"
    "sync"
)

var (
    Aurservd *aurservd

    Cabinet *pqm.Monitor[*ent.Cabinet]
    Bin     *pqm.Monitor[*ent.Bin]

    wg sync.WaitGroup
)

func Start() {
    dsn := g.Config.Postgres.Dsn

    // 启动同步任务
    Aurservd = newAurservd()
    go Aurservd.start()

    // TODO 同步消息删除
    Cabinet = pqm.NewMonitor(dsn, loki.StandardLogger(), &ent.Cabinet{}, func(message *pqm.Message[*ent.Cabinet]) {
        go Aurservd.SendCabinet(message.Data.Serial, message.Data)
    })

    Bin = pqm.NewMonitor(dsn, loki.StandardLogger(), &ent.Bin{}, func(message *pqm.Message[*ent.Bin]) {
        go Aurservd.SendBin(message.Data.Serial, message.Data)
    })

    // 启动数据数据库监听
    go Cabinet.Listen()
    go Bin.Listen()

    wg.Wait()
}
