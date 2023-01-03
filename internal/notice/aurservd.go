// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-29
// Based on cabservd by liasica, magicrolan@qq.com.

package notice

import (
    "context"
    "github.com/auroraride/adapter"
    "github.com/auroraride/adapter/codec"
    "github.com/auroraride/adapter/defs/cabdef"
    "github.com/auroraride/adapter/message"
    "github.com/auroraride/adapter/tcp"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/bin"
    "github.com/auroraride/cabservd/internal/g"
    log "github.com/sirupsen/logrus"
)

type aurservd struct {
    *tcp.Client
}

func newAurservd() *aurservd {
    return &aurservd{
        Client: tcp.NewClient(g.Config.Adapter.Aurservd, log.StandardLogger(), &codec.HeaderLength{}),
    }
}

func (h *aurservd) CabinetFullUpdate() {
    cabs, _ := ent.Database.Cabinet.Query().WithBins(func(query *ent.BinQuery) {
        query.Order(ent.Asc(bin.FieldOrdinal))
    }).All(context.Background())
    for _, cab := range cabs {
        h.SendFulldata(cab.Serial, cab, cab.Edges.Bins)
    }
}

func WrapCabinetMessage(full bool, serial string, cab *ent.Cabinet, bins ent.Bins) (message *cabdef.CabinetMessage) {
    // 不符合要求直接返回
    if cab == nil && len(bins) == 0 {
        log.Error("无可同步数据")
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
            Brand:       cab.Brand,
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
            Brand:         b.Brand,
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

func (h *aurservd) SendBattery(sn, serial string) {
    h.SendMessage(&cabdef.BatteryMessage{
        Battery: adapter.ParseBatterySN(sn),
        Cabinet: serial,
    })
}

func (h *aurservd) SendCabinet(serial string, cab *ent.Cabinet) {
    h.SendMessage(WrapCabinetMessage(false, serial, cab, nil))
}

func (h *aurservd) SendBin(serial string, b *ent.Bin) {
    h.SendMessage(WrapCabinetMessage(false, serial, nil, ent.Bins{b}))
}

func (h *aurservd) SendFulldata(serial string, cab *ent.Cabinet, bins ent.Bins) {
    h.SendMessage(WrapCabinetMessage(true, serial, cab, bins))
}

func (h *aurservd) SendMessage(data message.Messenger) {
    h.Sender <- data
}

func (h *aurservd) start() {
    h.Hooks.Start = func() {
        wg.Add(1)
    }

    h.Hooks.Connect = func() {
        h.CabinetFullUpdate()
        wg.Done()
    }

    h.Run()
}
