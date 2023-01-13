// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-29
// Based on cabservd by liasica, magicrolan@qq.com.

package core

import (
    "context"
    "fmt"
    "github.com/auroraride/adapter/defs/cabdef"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/bin"
    "github.com/auroraride/cabservd/internal/ent/cabinet"
    jsoniter "github.com/json-iterator/go"
    "github.com/liasica/go-helpers/tools"
    log "github.com/sirupsen/logrus"
    "time"
)

type ReportParser interface {
    GetSerial() (string, bool)
    GetCabinet() (*ent.CabinetPointer, bool)
    GetBins() ent.BinPointers
}

func UpdateCabinet(brand cabdef.Brand, p ReportParser) {
    ctx := context.Background()

    serial, ok := p.GetSerial()
    if !ok {
        return
    }

    cab, exists := p.GetCabinet()
    if exists {
        SaveCabinet(ctx, brand, serial, cab)
    }

    bins := p.GetBins()
    SaveBins(ctx, brand, serial, bins)
}

func LoadOrStoreCabinet(ctx context.Context, brand cabdef.Brand, serial string) (cab *ent.Cabinet) {
    client := ent.Database.Cabinet
    cab, _ = client.Query().Where(cabinet.Serial(serial)).First(ctx)
    if cab != nil {
        return
    }
    var err error
    cab, err = client.Create().SetSerial(serial).SetBrand(brand).Save(ctx)
    if err != nil {
        log.Errorf("电柜保存失败: %v", err)
    }
    return
}

func SaveCabinet(ctx context.Context, brand cabdef.Brand, serial string, item *ent.CabinetPointer) {
    log.Info(item)
    err := ent.Database.Cabinet.Create().
        SetBrand(brand).
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
        b, _ := jsoniter.Marshal(item)
        log.Errorf("电柜保存失败, %s: %v", string(b), err)
    }
}

func SaveBins(ctx context.Context, brand cabdef.Brand, serial string, items ent.BinPointers) {
    if len(items) == 0 {
        return
    }

    cab := LoadOrStoreCabinet(ctx, brand, serial)
    if cab == nil {
        log.Error("仓位保存失败: 未找到电柜信息")
        return
    }

    for _, item := range items {
        log.Info(item)
        uuid := tools.Md5String(fmt.Sprintf("%s_%s_%d", brand, serial, *item.Ordinal))
        err := ent.Database.Bin.Create().
            SetUUID(uuid).
            SetBrand(brand).
            SetSerial(serial).
            SetName(*item.Name).
            SetCabinetID(cab.ID).
            SetOrdinal(*item.Ordinal).
            OnConflictColumns(bin.FieldUUID).
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
            UpdateUUID().
            Exec(ctx)
        if err != nil {
            b, _ := jsoniter.Marshal(item)
            log.Errorf("仓位保存失败, %s: %v", string(b), err)
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
