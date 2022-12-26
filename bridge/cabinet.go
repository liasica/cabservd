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
    log "github.com/sirupsen/logrus"
)

type cabinet struct {
    *tcp.Client
}

var Cabinet *cabinet

func newCabinet() *cabinet {
    return &cabinet{
        tcp.NewClient(g.Config.Adapter.Cabinet, log.StandardLogger(), &codec.HeaderLength{}),
    }
}

func (c *cabinet) FullUpdate() {
    cabs := service.NewCabinet().QueryAllCabinets()
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
            Online:      cab.Online,
            Status:      cab.Status.String(),
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

func SendCabinet(serial string, cab *ent.Cabinet, bins ent.Bins) {
    Cabinet.Sender <- Cabinet.WrapData(serial, cab, bins)
}

func startCabinet() {
    Cabinet = newCabinet()
    Cabinet.Hooks.Connect = func() {
        go Cabinet.FullUpdate()
    }
    Cabinet.Run()
}
