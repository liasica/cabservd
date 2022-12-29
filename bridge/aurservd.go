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

type aurservd struct {
    *tcp.Client
}

var Aurservd *aurservd

func newAurservd() *aurservd {
    return &aurservd{
        Client: tcp.NewClient(g.Config.Adapter.Aurservd, log.StandardLogger(), &codec.HeaderLength{}),
    }
}

func (c *aurservd) CabinetFullUpdate() {
    cabs := service.NewCabinet().QueryAllCabinetWithBin()
    for _, cab := range cabs {
        c.Sender <- c.CabinetWrapData(cab.Serial, cab, cab.Edges.Bins)
    }
}

func (c *aurservd) CabinetWrapData(serial string, cab *ent.Cabinet, bins ent.Bins) (data *model.Data[*model.CabinetSyncData]) {
    // 不符合要求直接返回
    if cab == nil && len(bins) == 0 {
        log.Error("无可同步数据")
        return
    }

    data = &model.Data[*model.CabinetSyncData]{
        Type: model.DataTypeCabinetSync,
        Value: &model.CabinetSyncData{
            Serial: serial,
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

    for _, bin := range bins {
        data.Value.Bins = append(data.Value.Bins, &model.Bin{
            ID:            bin.ID,
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

func SendCabinet(serial string, cab *ent.Cabinet, bins ent.Bins) {
    Aurservd.Sender <- Aurservd.CabinetWrapData(serial, cab, bins)
}

func startAurservd() {
    Aurservd = newAurservd()
    Aurservd.Hooks.Connect = func() {
        Aurservd.CabinetFullUpdate()
        worker.Done()
    }
    Aurservd.Run()
}
