// Copyright (C) liasica. 2023-present.
//
// Created at 2023-03-01
// Based on cabservd by liasica, magicrolan@qq.com.

package rpc

//
// var (
//     cabinetStatusMap = map[cabinet.Status]pb.CabinetStatus{
//         cabinet.StatusInitializing: pb.CabinetStatus_initializing,
//         cabinet.StatusNormal:       pb.CabinetStatus_normal,
//         cabinet.StatusAbnormal:     pb.CabinetStatus_abnormal,
//     }
// )
//
// type cabinetServer struct {
//     pb.UnimplementedCabinetServer
// }
//
// func (s *cabinetServer) Batch(ctx context.Context, req *pb.CabinetBatchRequest) (res *pb.CabinetBatchResponse, err error) {
//     orm := ent.Database.Cabinet
//
//     var cabs ent.Cabinets
//     cabs, err = orm.Query().Where(cabinet.SerialIn(req.Serial...)).WithBins(func(q *ent.BinQuery) {
//         q.Order(ent.Asc(bin.FieldOrdinal))
//     }).All(ctx)
//     if err != nil {
//         return
//     }
//
//     res = &pb.CabinetBatchResponse{Items: make([]*pb.CabinetItem, len(cabs))}
//     for m, cab := range cabs {
//         res.Items[m] = &pb.CabinetItem{
//             Serial:      cab.Serial,
//             Status:      cabinetStatusMap[cab.Status],
//             Enable:      cab.Enable,
//             Lng:         cab.Lng,
//             Lat:         cab.Lat,
//             Gsm:         cab.Gsm,
//             Voltage:     cab.Voltage,
//             Current:     cab.Current,
//             Temperature: cab.Temperature,
//             Electricity: cab.Electricity,
//             Bins:        make([]*pb.CabinetBinItem, len(cab.Edges.Bins)),
//         }
//
//         for n, b := range cab.Edges.Bins {
//             res.Items[m].Bins[n] = &pb.CabinetBinItem{
//                 Ordinal:       int64(b.Ordinal),
//                 Open:          b.Open,
//                 Enable:        b.Enable,
//                 Health:        b.Health,
//                 BatteryExists: b.BatteryExists,
//                 BatterySn:     b.BatterySn,
//                 Voltage:       b.Voltage,
//                 Current:       b.Current,
//                 Soc:           b.Soc,
//                 Soh:           b.Soh,
//             }
//         }
//     }
//
//     return
// }
//
// func (s *cabinetServer) Sync(ctx context.Context, req *pb.CabinetSyncRequest) (res *pb.CabinetSyncResponse, err error) {
//     orm := ent.Database.Cabinet
//
//     var cabs ent.Cabinets
//     cabs, err = orm.Query().Where(cabinet.SerialIn(req.Serial...)).WithBins().All(ctx)
//     if err != nil {
//         return
//     }
//
//     res = &pb.CabinetSyncResponse{Items: make([]*pb.CabinetSyncItem, len(cabs))}
//     for m, c := range cabs {
//         item := &pb.CabinetSyncItem{
//             Bins: make([]*pb.CabinetSyncBin, len(c.Edges.Bins)),
//         }
//
//         item.Health = pb.CabinetSyncItem_offline
//         if c.Online {
//             item.Health = pb.CabinetSyncItem_online
//         }
//         if c.Status == cabinet.StatusAbnormal {
//             item.Health = pb.CabinetSyncItem_fault
//         }
//
//         for n, b := range c.Edges.Bins {
//             item.Bins[n] = &pb.CabinetSyncBin{
//                 Index:     int64(b.Ordinal - 1),
//                 BatterySn: b.BatterySn,
//                 Battery:   b.BatteryExists,
//                 Remark:    b.Remark,
//                 Soc:       b.Soc,
//                 Current:   b.Current,
//                 Voltage:   b.Voltage,
//
//                 Full:       false,
//                 OpenStatus: false,
//                 DoorHealth: false,
//                 Faults:     nil,
//             }
//
//             if b.BatteryExists {
//                 item.BatteryNum += 1
//                 if b.Soc >= req.FullSoc {
//                     // 满电
//                     item.BatteryFullNum += 1
//                     item.Bins[n].Full = true
//                 } else {
//                     // 充电
//                     item.BatteryChargingNum += 1
//                 }
//             } else {
//                 // 空仓
//                 item.EmptyBinNum += 1
//             }
//
//             //             hasBattery := b.BatteryExists && b.BatterySn != ""
//             //             var (
//             //                 isFull bool
//             //                 remark string
//             //             )
//             //             if b.Remark != nil {
//             //                 remark = *b.Remark
//             //             }
//             //             // 电池数
//             //             if hasBattery {
//             //                 // 如果该仓位有电池
//             //                 // 智能电柜操作放入电池
//             //                 if cab.Intelligent {
//             //                     _, _ = NewBattery().SyncPutin(b.BatterySn, cab.Serial, cab.ID, b.Ordinal, bins)
//             //                 }
//             //                 bn += 1
//             //                 if b.Soc >= model.IntelligentBatteryFullSoc {
//             //                     // 满电
//             //                     bf += 1
//             //                     isFull = true
//             //                 } else {
//             //                     // 充电
//             //                     bc += 1
//             //                 }
//             //             } else {
//             //                 // 如果该仓位无电池
//             //                 NewBattery().SyncPutout(cab.ID, b.Ordinal)
//             //                 // 空仓
//             //                 be += 1
//             //             }
//             //             // 锁仓
//             //             if !b.Enable {
//             //                 bl += 1
//             //             }
//             //
//             //             // 新仓位信息
//             //             nb := &model.CabinetBin{
//             //                 Index:       b.Ordinal - 1,
//             //                 Name:        b.Name,
//             //                 BatterySN:   b.BatterySn,
//             //                 Full:        isFull,
//             //                 Battery:     hasBattery,
//             //                 Electricity: model.NewBatterySoc(b.Soc),
//             //                 OpenStatus:  b.Open,
//             //                 DoorHealth:  b.Health && b.Enable,
//             //                 Current:     b.Current,
//             //                 Voltage:     b.Voltage,
//             //                 Remark:      remark,
//             //             }
//             //
//             //             if data.Full || len(cab.Bin) < len(data.Bins) {
//             //                 bins = append(bins, nb)
//             //             }
//             //
//             //             if !data.Full {
//             //                 for i, xb := range bins {
//             //                     if xb.Index+1 == b.Ordinal {
//             //                         bins[i] = nb
//             //                     }
//             //                 }
//             //             }
//             //         }
//             //
//             //         sort.Slice(bins, func(i, j int) bool {
//             //             return bins[i].Index <= bins[j].Index
//             //         })
//             //
//             //         updater.SetDoors(len(bins)).
//             //             SetBatteryNum(bn).
//             //             SetBatteryFullNum(bf).
//             //             SetBatteryChargingNum(bc).
//             //             SetEmptyBinNum(be).
//             //             SetLockedBinNum(bl).
//             //             SetBin(bins)
//         }
//
//         res.Items[m] = item
//     }
//
//     return
// }
