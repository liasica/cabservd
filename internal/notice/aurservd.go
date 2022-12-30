// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-29
// Based on cabservd by liasica, magicrolan@qq.com.

package notice

import (
    "context"
    "github.com/auroraride/adapter/codec"
    "github.com/auroraride/adapter/model"
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

func (c *aurservdHook) CabinetFullUpdate() {
    cabs, _ := ent.Database.Cabinet.Query().WithBins(func(query *ent.BinQuery) {
        query.Order(ent.Asc(bin.FieldOrdinal))
    }).All(context.Background())
    for _, cab := range cabs {
        SendCabinet(true, cab.Serial, cab, cab.Edges.Bins)
    }
}

func (c *aurservdHook) CabinetWrapData(full bool, serial string, cab *ent.Cabinet, bins ent.Bins) (data *model.Data[*model.CabinetSyncData]) {
    // 不符合要求直接返回
    if cab == nil && len(bins) == 0 {
        log.Error("无可同步数据")
        return
    }

    data = &model.Data[*model.CabinetSyncData]{
        Type: model.DataTypeCabinetSync,
        Value: &model.CabinetSyncData{
            Serial: serial,
            Full:   full,
        },
    }

    if cab != nil {
        data.Value.Cabinet = &model.Cabinet{
            ID:          cab.ID,
            Online:      cab.Online,
            Brand:       cab.Brand,
            Serial:      cab.Serial,
            Status:      model.CabinetStatus(cab.Status),
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
        data.Value.Bins = append(data.Value.Bins, &model.Bin{
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

func SendCabinet(full bool, serial string, cab *ent.Cabinet, bins ent.Bins) {
    Aurservd.Sender <- Aurservd.CabinetWrapData(full, serial, cab, bins)
}

func (*aurservdHook) Start() {
    Aurservd.Hooks.Start = func() {
        worker.Add(1)
    }

    Aurservd.Hooks.Connect = func() {
        Aurservd.CabinetFullUpdate()
        worker.Done()
    }

    Aurservd.Run()
}
