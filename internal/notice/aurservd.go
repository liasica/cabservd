// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-29
// Based on cabservd by liasica, magicrolan@qq.com.

package notice

import (
    "context"
    "github.com/auroraride/adapter"
    "github.com/auroraride/adapter/codec"
    "github.com/auroraride/adapter/tcp"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/bin"
    "github.com/auroraride/cabservd/internal/g"
    log "github.com/sirupsen/logrus"
)

type aurservdHook struct {
    *tcp.Client
}

func NewAurservd() *aurservdHook {
    return &aurservdHook{
        Client: tcp.NewClient(g.Config.Adapter.Aurservd, log.StandardLogger(), &codec.HeaderLength{}),
    }
}

func (*aurservdHook) CabinetFullUpdate() {
    cabs, _ := ent.Database.Cabinet.Query().WithBins(func(query *ent.BinQuery) {
        query.Order(ent.Asc(bin.FieldOrdinal))
    }).All(context.Background())
    for _, cab := range cabs {
        Aurservd.SendCabinet(true, cab.Serial, cab, cab.Edges.Bins)
    }
}

func (*aurservdHook) CabinetWrapData(full bool, serial string, cab *ent.Cabinet, bins ent.Bins) (message *adapter.CabinetMessage) {
    // 不符合要求直接返回
    if cab == nil && len(bins) == 0 {
        log.Error("无可同步数据")
        return
    }

    message = &adapter.CabinetMessage{
        Serial: serial,
        Full:   full,
    }

    if cab != nil {
        message.Cabinet = &adapter.Cabinet{
            ID:          cab.ID,
            Online:      cab.Online,
            Brand:       cab.Brand,
            Serial:      cab.Serial,
            Status:      adapter.CabinetStatus(cab.Status),
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
        message.Bins = append(message.Bins, &adapter.Bin{
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

func (h *aurservdHook) SendBattery(sn, serial string) {
    h.Sender <- &adapter.BatteryMessage{
        Battery: adapter.ParseBatterySN(sn),
        Cabinet: serial,
    }
}

func (h *aurservdHook) SendCabinet(full bool, serial string, cab *ent.Cabinet, bins ent.Bins) {
    h.Sender <- Aurservd.CabinetWrapData(full, serial, cab, bins)
}

func (h *aurservdHook) SendData(data adapter.Messenger) {
    h.Sender <- data
}

func (h *aurservdHook) Start() {
    h.Hooks.Start = func() {
        worker.Add(1)
    }

    h.Hooks.Connect = func() {
        h.CabinetFullUpdate()
        worker.Done()
    }

    h.Run()
}
