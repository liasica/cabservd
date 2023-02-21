// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-29
// Based on cabservd by liasica, magicrolan@qq.com.

package core

import (
    "context"
    "github.com/auroraride/adapter/log"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/bin"
    "github.com/auroraride/cabservd/internal/ent/cabinet"
    "go.uber.org/zap"
    "time"
)

type CabinetUpdater interface {
    GetSerial() (string, bool)
    GetCabinet() (*ent.CabinetPointer, bool)
    GetBins() ent.BinPointers
}

func UpdateCabinet(p CabinetUpdater) {
    ctx := context.Background()

    serial, ok := p.GetSerial()
    if !ok {
        return
    }

    cab, exists := p.GetCabinet()
    if exists {
        SaveCabinet(ctx, serial, cab)
    }

    bins := p.GetBins()
    SaveBins(ctx, serial, bins)
}

func LoadOrStoreCabinet(ctx context.Context, serial string) (cab *ent.Cabinet) {
    orm := ent.Database.Cabinet
    cab, _ = orm.Query().Where(cabinet.Serial(serial)).First(ctx)
    if cab != nil {
        return
    }
    var err error
    cab, err = orm.Create().SetSerial(serial).Save(ctx)
    if err != nil {
        zap.L().Error("电柜保存失败", zap.Error(err))
    }
    return
}

func SaveCabinet(ctx context.Context, serial string, item *ent.CabinetPointer) {
    err := ent.Database.Cabinet.Create().
        SetSerial(serial).
        OnConflictColumns(cabinet.FieldSerial).
        Update(func(u *ent.CabinetUpsert) {
            u.SetUpdatedAt(time.Now())
            // 在线
            if item.Online != nil {
                u.SetOnline(*item.Online)
            }

            // 状态
            if item.Status != nil {
                u.SetStatus(*item.Status)
            }

            // 经度
            if item.Lng != nil {
                u.SetLng(*item.Lng)
            }

            // 纬度
            if item.Lat != nil {
                u.SetLat(*item.Lat)
            }

            // GSM
            if item.Gsm != nil {
                u.SetGsm(*item.Gsm)
            }

            // 电压
            if item.Voltage != nil {
                u.SetVoltage(*item.Voltage)
            }

            // 电流
            if item.Current != nil {
                u.SetCurrent(*item.Current)
            }

            // 温度
            if item.Temperature != nil {
                u.SetTemperature(*item.Temperature)
            }

            // 启用
            if item.Enable != nil {
                u.SetEnable(*item.Enable)
            }

            // 总用电
            if item.Electricity != nil {
                u.SetElectricity(*item.Electricity)
            }
        }).
        Exec(ctx)
    if err != nil {
        zap.L().Error("电柜保存失败", zap.Error(err), log.Payload(item))
    }
}

func SaveBins(ctx context.Context, serial string, items ent.BinPointers) {
    if len(items) == 0 {
        return
    }

    cab := LoadOrStoreCabinet(ctx, serial)
    if cab == nil {
        zap.L().Error("仓位保存失败: 未找到电柜信息")
        return
    }

    for _, item := range items {
        err := ent.Database.Bin.Create().
            SetSerial(serial).
            SetName(*item.Name).
            SetCabinetID(cab.ID).
            SetOrdinal(*item.Ordinal).
            OnConflictColumns(bin.FieldSerial, bin.FieldOrdinal).
            Update(func(u *ent.BinUpsert) {
                // 更新时间和电柜ID
                u.SetUpdatedAt(time.Now()).SetCabinetID(cab.ID)
                // 健康状态
                if item.Health != nil {
                    u.SetHealth(*item.Health)
                }

                // 仓门状态
                if item.Open != nil {
                    u.SetOpen(*item.Open)
                }

                // 仓位启用状态
                if item.Enable != nil {
                    u.SetEnable(*item.Enable)
                }

                // 电压
                if item.Voltage != nil {
                    u.SetVoltage(*item.Voltage)
                }

                // 电流
                if item.Current != nil {
                    u.SetCurrent(*item.Current)
                }

                // 电量
                if item.Soc != nil {
                    u.SetSoc(*item.Soc)
                }

                // 健康
                if item.Soh != nil {
                    u.SetSoh(*item.Soh)
                }

                // 电池编号
                if item.BatterySn != nil {
                    u.SetBatterySn(*item.BatterySn)
                }

                // 电池在位
                if item.BatteryExists != nil {
                    u.SetBatteryExists(*item.BatteryExists)
                    // TODO 电池不在位是否清除电池信息
                    // if !*item.BatteryExists {
                    //     u.ResetBattery()
                    // }
                }
            }).
            Exec(ctx)
        if err != nil {
            zap.L().Error("仓位保存失败", zap.Error(err), log.Payload(item))
        }
    }
}

// ResetBins 重置电柜仓位信息
func ResetBins(sn string) error {
    return ent.Database.Bin.Update().
        Where(bin.Serial(sn)).
        SetBatterySn("").
        SetSoc(0).
        SetSoh(0).
        SetVoltage(0).
        SetCurrent(0).
        // SetEnable(true). // TODO 是否单独设置LOCK
        SetOpen(false).
        Exec(context.Background())
}
