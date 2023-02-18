// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-29
// Based on cabservd by liasica, magicrolan@qq.com.

package sync

import (
    "github.com/auroraride/adapter/defs/cabdef"
    "github.com/auroraride/adapter/sync"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/g"
)

var (
    syncCabinet  *sync.Sync[cabdef.CabinetMessage]
    syncExchange *sync.Sync[cabdef.ExchangeStepMessage]
)

func createSync() {
    syncCabinet = sync.New[cabdef.CabinetMessage](
        g.Redis,
        g.Config.Environment,
        sync.StreamCabinet,
        nil,
    )

    syncExchange = sync.New[cabdef.ExchangeStepMessage](
        g.Redis,
        g.Config.Environment,
        sync.StreamExchange,
        nil,
    )
}

func WrapCabinetMessage(full bool, serial string, cab *ent.Cabinet, bins ent.Bins) (message *cabdef.CabinetMessage) {
    // 不符合要求直接返回
    if cab == nil && len(bins) == 0 {
        return
    }

    message = &cabdef.CabinetMessage{
        Serial: serial,
        Full:   full,
    }

    if cab != nil {
        message.Cabinet = &cabdef.Cabinet{
            ID:          cab.ID,
            Online:      cab.Online,
            Brand:       g.Config.Brand,
            Serial:      cab.Serial,
            Status:      cabdef.CabinetStatus(cab.Status),
            Enable:      cab.Enable,
            Lng:         cab.Lng,
            Lat:         cab.Lat,
            Gsm:         cab.Gsm,
            Voltage:     cab.Voltage,
            Current:     cab.Current,
            Temperature: cab.Temperature,
            Electricity: cab.Electricity,
        }
    }

    for _, b := range bins {
        message.Bins = append(message.Bins, &cabdef.Bin{
            ID:            b.ID,
            Brand:         g.Config.Brand,
            Serial:        b.Serial,
            Name:          b.Name,
            Ordinal:       b.Ordinal,
            Open:          b.Open,
            Enable:        b.Enable,
            Health:        b.Health,
            BatteryExists: b.BatteryExists,
            BatterySn:     b.BatterySn,
            Voltage:       b.Voltage,
            Current:       b.Current,
            Soc:           b.Soc,
            Soh:           b.Soh,
            Remark:        b.Remark,
        })
    }

    return
}

func SendCabinet(serial string, cab *ent.Cabinet) {
    SendMessage(WrapCabinetMessage(false, serial, cab, nil))
}

func SendBin(serial string, b *ent.Bin) {
    SendMessage(WrapCabinetMessage(false, serial, nil, ent.Bins{b}))
}

func SendCabinetFull(serial string, cab *ent.Cabinet, bins ent.Bins) {
    SendMessage(WrapCabinetMessage(true, serial, cab, bins))
}

func SendMessage(data any) {
    switch message := data.(type) {
    case *cabdef.CabinetMessage:
        syncCabinet.Push(message)
    case *cabdef.ExchangeStepMessage:
        syncExchange.Push(message)
    }
}
