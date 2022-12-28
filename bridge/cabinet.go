// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-25
// Based on cabservd by liasica, magicrolan@qq.com.

package bridge

import (
    "github.com/auroraride/adapter/codec"
    "github.com/auroraride/adapter/model"
    "github.com/auroraride/adapter/tcp"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/g"
    "github.com/auroraride/cabservd/internal/service"
    "github.com/goccy/go-json"
    "github.com/lib/pq"
    log "github.com/sirupsen/logrus"
)

type cabinet struct {
    *tcp.Client
}

var Cabinet *cabinet

func newCabinet() *cabinet {
    return &cabinet{
        Client: tcp.NewClient(g.Config.Adapter.Cabinet, log.StandardLogger(), &codec.HeaderLength{}),
    }
}

func (c *cabinet) FullUpdate() {
    panic("TODO")
    cabs := service.NewCabinet().QueryAllCabinetWithBin()
    for _, cab := range cabs {
        c.Sender <- c.WrapData(cab.Serial, cab, cab.Edges.Bins)
    }
}

func (c *cabinet) WrapData(serial string, cab *ent.Cabinet, bins ent.Bins) (data *model.CabinetSyncRequest) {
    // 不符合要求直接返回
    if cab == nil && len(bins) == 0 {
        log.Error("无可同步数据")
        return
    }

    data = &model.CabinetSyncRequest{
        Serial: serial,
    }

    if cab != nil {
        data.Cabinet = &model.Cabinet{
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

    for _, bin := range bins {
        data.Bins = append(data.Bins, &model.Bin{
            ID:            bin.ID,
            UUID:          bin.UUID,
            CabinetID:     bin.CabinetID,
            Brand:         bin.Brand,
            Serial:        bin.Serial,
            Name:          bin.Name,
            Ordinal:       bin.Ordinal,
            Open:          bin.Open,
            Enable:        bin.Enable,
            Health:        bin.Health,
            BatteryExists: bin.BatteryExists,
            BatterySn:     bin.BatterySn,
            Voltage:       bin.Voltage,
            Current:       bin.Current,
            Soc:           bin.Soc,
            Soh:           bin.Soh,
            Remark:        bin.Remark,
        })
    }

    return
}

func SendCabinetSyncData(n *pq.Notification) {
    type notificationData interface {
        *ent.Cabinet | *ent.Bin
    }

    type data[T notificationData] struct {
        Table  string `json:"table"`
        Action string `json:"action"`
        Data   T      `json:"data"`
    }

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

    SendCabinet(serial, cab, bins)
}

func SendCabinet(serial string, cab *ent.Cabinet, bins ent.Bins) {
    Cabinet.Sender <- Cabinet.WrapData(serial, cab, bins)
}

func startCabinet() {
    Cabinet = newCabinet()
    Cabinet.Hooks.Connect = func() {
        // go Cabinet.FullUpdate()
        worker.Done()
    }
    Cabinet.Run()
}
