// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-29
// Based on cabservd by liasica, magicrolan@qq.com.

package sync

import (
    "github.com/auroraride/adapter/defs/batdef"
    "github.com/auroraride/adapter/defs/cabdef"
    "github.com/auroraride/adapter/sync"
    "github.com/auroraride/cabservd/internal/g"
)

var (
    syncBatteryFlow *sync.Sync[batdef.BatteryFlow]
    syncExchange    *sync.Sync[cabdef.ExchangeStepMessage]
)

func createSync() {
    syncBatteryFlow = sync.New[batdef.BatteryFlow](
        g.Redis,
        g.Config.Environment,
        sync.StreamBatteryFlow,
        nil,
    )

    syncExchange = sync.New[cabdef.ExchangeStepMessage](
        g.Redis,
        g.Config.Environment,
        sync.StreamExchange,
        nil,
    )
}

func SendMessage(data any) {
    switch message := data.(type) {
    case *batdef.BatteryFlow:
        syncBatteryFlow.Push(message)
    case *cabdef.ExchangeStepMessage:
        syncExchange.Push(message)
    }
}
