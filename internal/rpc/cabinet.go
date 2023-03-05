// Copyright (C) liasica. 2023-present.
//
// Created at 2023-03-01
// Based on cabservd by liasica, magicrolan@qq.com.

package rpc

import (
    "context"
    "github.com/auroraride/adapter/rpc/pb"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/bin"
    "github.com/auroraride/cabservd/internal/ent/cabinet"
)

var (
    cabinetStatusMap = map[cabinet.Status]pb.CabinetStatus{
        cabinet.StatusInitializing: pb.CabinetStatus_initializing,
        cabinet.StatusNormal:       pb.CabinetStatus_normal,
        cabinet.StatusAbnormal:     pb.CabinetStatus_abnormal,
    }
)

type cabinetServer struct {
    pb.UnimplementedCabinetServer
}

func (s *cabinetServer) Batch(ctx context.Context, req *pb.CabinetBatchRequest) (res *pb.CabinetBatchResponse, err error) {
    orm := ent.Database.Cabinet

    var cabs ent.Cabinets
    cabs, err = orm.Query().Where(cabinet.SerialIn(req.Serial...)).WithBins(func(q *ent.BinQuery) {
        q.Order(ent.Asc(bin.FieldOrdinal))
    }).All(ctx)
    if err != nil {
        return
    }

    res = &pb.CabinetBatchResponse{Items: make([]*pb.CabinetItem, len(cabs))}
    for m, cab := range cabs {
        res.Items[m] = &pb.CabinetItem{
            Serial:      cab.Serial,
            Status:      cabinetStatusMap[cab.Status],
            Enable:      cab.Enable,
            Lng:         cab.Lng,
            Lat:         cab.Lat,
            Gsm:         cab.Gsm,
            Voltage:     cab.Voltage,
            Current:     cab.Current,
            Temperature: cab.Temperature,
            Electricity: cab.Electricity,
            Bins:        make([]*pb.CabinetBinItem, len(cab.Edges.Bins)),
        }

        for n, b := range cab.Edges.Bins {
            res.Items[m].Bins[n] = &pb.CabinetBinItem{
                Ordinal:       int64(b.Ordinal),
                Open:          b.Open,
                Enable:        b.Enable,
                Health:        b.Health,
                BatteryExists: b.BatteryExists,
                BatterySn:     b.BatterySn,
                Voltage:       b.Voltage,
                Current:       b.Current,
                Soc:           b.Soc,
                Soh:           b.Soh,
            }
        }
    }

    return
}

func (s *cabinetServer) Sync(ctx context.Context, req *pb.CabinetSyncRequest) (res *pb.CabinetSyncResponse, err error) {
    orm := ent.Database.Cabinet

    var cabs ent.Cabinets
    cabs, err = orm.Query().Where(cabinet.SerialIn(req.Serial...)).WithBins(func(bq *ent.BinQuery) { bq.Order(ent.Asc(bin.FieldOrdinal)) }).All(ctx)
    if err != nil {
        return
    }

    res = &pb.CabinetSyncResponse{Items: make([]*pb.CabinetSyncItem, len(cabs))}
    for m, c := range cabs {
        item := &pb.CabinetSyncItem{
            Serial: c.Serial,
            Bins:   make([]*pb.CabinetSyncBin, len(c.Edges.Bins)),
        }

        item.Health = pb.CabinetSyncItem_offline
        if c.Online {
            item.Health = pb.CabinetSyncItem_online
        }
        if c.Status == cabinet.StatusAbnormal {
            item.Health = pb.CabinetSyncItem_fault
        }

        for n, b := range c.Edges.Bins {

            // 新仓位信息
            nb := &pb.CabinetSyncBin{
                Index:     int64(b.Ordinal - 1),
                BatterySn: b.BatterySn,
                Battery:   b.BatteryExists,
                Soc:       b.Soc,
                Current:   b.Current,
                Voltage:   b.Voltage,

                OpenStatus: b.Open,
                DoorHealth: b.Health && b.Enable,
                Faults:     make([]string, 0), // TODO 需要保存记录仓位故障
            }

            if b.Remark != nil {
                nb.Remark = *b.Remark
            }

            if b.BatteryExists {
                // 有电池仓位
                item.BatteryNum += 1
                if b.Soc >= req.FullSoc {
                    nb.Full = true
                    // 满电
                    item.BatteryFullNum += 1
                    nb.Full = true
                } else {
                    // 充电
                    item.BatteryChargingNum += 1
                }
            } else {
                // 空仓
                item.EmptyBinNum += 1
            }

            // 锁仓
            if !b.Enable {
                item.LockedBinNum += 1
            }

            item.Bins[n] = nb
        }

        res.Items[m] = item
    }

    return
}
