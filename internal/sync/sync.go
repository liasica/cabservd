// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on cabservd by liasica, magicrolan@qq.com.

package sync

import (
    "context"
    "github.com/auroraride/adapter/pqm"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/bin"
    "github.com/auroraride/cabservd/internal/g"
    "github.com/auroraride/cabservd/internal/types"
    "sync"
)

var (
    Cabinet *pqm.Monitor[*ent.Cabinet]
    Bin     *pqm.Monitor[*ent.Bin]

    wg sync.WaitGroup
)

func Start() {
    // 创建同步客户端
    createSync()

    // 获取所有电柜
    cabs, _ := ent.Database.Cabinet.Query().WithBins(func(query *ent.BinQuery) { query.Order(ent.Asc(bin.FieldOrdinal)) }).All(context.Background())

    // 缓存电柜信息
    for _, cab := range cabs {
        go SendCabinetFull(cab.Serial, cab, cab.Edges.Bins)
        g.Redis.HSet(context.Background(), g.CacheCabinetKey, cab.Serial, &types.CabinetCache{
            Lng: cab.Lng,
            Lat: cab.Lat,
        })
    }

    dsn := g.Config.Postgres.Dsn

    Cabinet = pqm.NewMonitor(dsn, &ent.Cabinet{}, func(message *pqm.Message[*ent.Cabinet]) {
        g.Redis.HSet(context.Background(), g.CacheCabinetKey, message.Data.Serial, &types.CabinetCache{
            Lng: message.Data.Lng,
            Lat: message.Data.Lat,
        })
        go SendCabinet(message.Data.Serial, message.Data)
    })

    Bin = pqm.NewMonitor(dsn, &ent.Bin{}, func(message *pqm.Message[*ent.Bin]) {
        go SendBin(message.Data.Serial, message.Data)
    })

    // 启动数据数据库监听
    go Cabinet.Listen()
    go Bin.Listen()

    wg.Wait()
}
