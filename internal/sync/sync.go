// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on cabservd by liasica, magicrolan@qq.com.

package sync

import (
    "context"
    "github.com/auroraride/adapter"
    "github.com/auroraride/adapter/defs/batdef"
    "github.com/auroraride/adapter/log"
    "github.com/auroraride/adapter/pqm"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/bin"
    "github.com/auroraride/cabservd/internal/ent/cabinet"
    "github.com/auroraride/cabservd/internal/g"
    "github.com/auroraride/cabservd/internal/types"
    jsoniter "github.com/json-iterator/go"
    "github.com/liasica/go-helpers/silk"
    "go.uber.org/zap"
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
    cabs, _ := ent.Database.Cabinet.Query().WithBins(func(query *ent.BinQuery) { query.Order(ent.Asc(bin.FieldOrdinal)) }).Where(cabinet.Brand(g.Config.Brand)).All(context.Background())

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
        go reign(message.Data, message.Old)
        go SendBin(message.Data.Serial, message.Data)
    })

    // 启动数据数据库监听
    go Cabinet.Listen()
    go Bin.Listen()

    wg.Wait()
}

func reign(data, old *ent.Bin) {
    // 如果电池无变化直接跳过
    if old.BatterySn == data.BatterySn {
        return
    }

    // 从缓存中获取电柜信息
    var result types.CabinetCache
    err := g.Redis.HGet(context.Background(), g.CacheCabinetKey, data.Serial).Scan(&result)
    if err != nil {
        zap.L().Error("从缓存中获取电柜信息失败", zap.Error(err))
    }

    item := &batdef.Reign{
        Serial:  data.Serial,
        Ordinal: silk.Int(data.Ordinal),
        Lng:     result.Lng,
        Lat:     result.Lat,
    }

    if old.BatterySn == "" {
        // 放入 (旧无新有)
        item.Action = batdef.ReignActionIn
        item.SN = data.BatterySn
    } else {
        // 取出 (旧有新无)
        item.SN = old.BatterySn
        item.Action = batdef.ReignActionOut
    }

    // 替换 (旧有新有)
    // 混乱体 (一般是服务器停机阶段有更新导致的)
    if old.BatterySn != "" && data.BatterySn != "" {
        go doReign(item.Clone(data.BatterySn, batdef.ReignActionIn))
    }

    doReign(item)
}

func doReign(data *batdef.Reign) {
    bat := adapter.ParseBatterySN(data.SN)
    url, err := g.Config.GetBmsApiUrl(bat.Brand, "/battery/reign")
    b, _ := jsoniter.Marshal(data)
    if err != nil {
        zap.L().Error("电池在位请求失败", zap.Error(err), log.ResponseBody(b))
        return
    }

    _, _ = adapter.FastRequest[*adapter.Response[batdef.ReignResponse]](url, adapter.RquestMethodPost, b)
}
