// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-25
// Based on cabservd by liasica, magicrolan@qq.com.

package bridge

import (
    "github.com/auroraride/bridge"
    "github.com/auroraride/bridge/pb"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/cabinet"
    "github.com/auroraride/cabservd/internal/g"
    "github.com/auroraride/cabservd/internal/service"
    log "github.com/sirupsen/logrus"
)

type cabinetBridge struct {
    bridger *bridge.Cabinet

    sender chan *pb.CabinetSyncRequest
}

var (
    CabinetStatus = map[cabinet.Status]pb.CabinetStatus{
        cabinet.StatusInitializing: pb.CabinetStatus_INITIALIZING,
        cabinet.StatusIdle:         pb.CabinetStatus_IDLE,
        cabinet.StatusBusy:         pb.CabinetStatus_BUSY,
        cabinet.StatusExchange:     pb.CabinetStatus_EXCHANGE,
        cabinet.StatusAbnormal:     pb.CabinetStatus_ABNORMAL,
    }
)

func newCabinet() *cabinetBridge {
    return &cabinetBridge{
        bridger: bridge.NewCabinet(log.StandardLogger()),
        sender:  make(chan *pb.CabinetSyncRequest),
    }
}

func (c *cabinetBridge) FullUpdate() {
    cabs := service.NewCabinet().QueryAllCabinets()
    for _, cab := range cabs {
        c.sender <- c.WrapData(cab.Serial, cab, cab.Edges.Bins)
    }
}

func (c *cabinetBridge) WrapData(serial string, cab *ent.Cabinet, bins ent.Bins) (data *pb.CabinetSyncRequest) {
    // 不符合要求直接返回
    if cab == nil && len(bins) == 0 {
        log.Error("无可同步数据")
        return
    }

    data = &pb.CabinetSyncRequest{
        Serial: serial,
    }

    if cab != nil {
        data.Cabinet = &pb.CabinetData{
            Online:      cab.Online,
            Status:      CabinetStatus[cab.Status],
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
        data.Bins = append(data.Bins, &pb.BinData{
            Ordinal:       int32(bin.Ordinal),
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

func (c *cabinetBridge) run() {
    go c.bridger.RunClient(g.Config.Bridge.Address, func() {
        c.FullUpdate()
    })

    for {
        select {
        case data := <-c.sender:
            c.bridger.SendSyncData(data)
        }
    }
}
