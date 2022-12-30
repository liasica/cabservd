// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-29
// Based on cabservd by liasica, magicrolan@qq.com.

package core

import (
    "context"
    "fmt"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/bin"
    "github.com/auroraride/cabservd/internal/ent/cabinet"
    "github.com/goccy/go-json"
    "github.com/liasica/go-helpers/tools"
    log "github.com/sirupsen/logrus"
)

type Parser interface {
    Bins() ent.BinPointers
    Cabinet() (*ent.CabinetPointer, bool)
}

func UpdateCabinet(brand, serial string, p Parser) {
    ctx := context.Background()

    cab, exists := p.Cabinet()
    if exists {
        SaveCabinetContext(ctx, brand, serial, cab)
    }

    bins := p.Bins()
    SaveBinsContext(ctx, brand, serial, bins)
}

func SaveCabinetContext(ctx context.Context, brand, serial string, item *ent.CabinetPointer) {
    err := ent.Database.Cabinet.Create().
        SetBrand(brand).
        SetSerial(serial).
        OnConflictColumns(cabinet.FieldSerial).
        Update(func(u *ent.CabinetUpsert) {
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
        }).Exec(ctx)
    if err != nil {
        b, _ := json.Marshal(item)
        log.Errorf("电柜保存失败, %s: %v", string(b), err)
    }
}

func SaveBinsContext(ctx context.Context, brand, serial string, items ent.BinPointers) {
    if len(items) == 0 {
        return
    }

    cab, _ := ent.Database.Cabinet.Query().Where(cabinet.Serial(serial)).First(ctx)
    if cab == nil {
        log.Error("仓位保存失败: 未找到电柜信息")
        return
    }

    for _, item := range items {
        uuid := tools.Md5String(fmt.Sprintf("%s_%s_%d", brand, serial, *item.Ordinal))
        name := fmt.Sprintf("%d号仓", *item.Ordinal)
        err := ent.Database.Bin.Create().
            SetUUID(uuid).
            SetBrand(brand).
            SetSerial(serial).
            SetName(name).
            SetCabinetID(cab.ID).
            SetOrdinal(*item.Ordinal).
            OnConflictColumns(bin.FieldUUID).
            Update(func(u *ent.BinUpsert) {
                // 健康状态
                if item.Health != nil {
                    fmt.Printf("%s health :-> %v\n", name, *item.Health)
                    u.SetHealth(*item.Health)
                }

                // 仓门状态
                if item.Open != nil {
                    fmt.Printf("%s open :-> %v\n", name, *item.Open)
                    u.SetOpen(*item.Open)
                }

                // 仓位启用状态
                if item.Enable != nil {
                    fmt.Printf("%s enable :-> %v\n", name, *item.Enable)
                    u.SetEnable(*item.Enable)
                }

                // 电压
                if item.Voltage != nil {
                    fmt.Printf("%s voltage :-> %v\n", name, *item.Voltage)
                    u.SetVoltage(*item.Voltage)
                }

                // 电流
                if item.Current != nil {
                    fmt.Printf("%s current :-> %v\n", name, *item.Current)
                    u.SetCurrent(*item.Current)
                }

                // 电量
                if item.Soc != nil {
                    fmt.Printf("%s soc :-> %v\n", name, *item.Soc)
                    u.SetSoc(*item.Soc)
                }

                // 健康
                if item.Soh != nil {
                    fmt.Printf("%s soh :-> %v\n", name, *item.Soh)
                    u.SetSoh(*item.Soh)
                }

                // 电池编号
                if item.BatterySn != nil {
                    fmt.Printf("%s battery :-> %v\n", name, *item.BatterySn)
                    u.SetBatterySn(*item.BatterySn)
                }

                // 电池在位
                if item.BatteryExists != nil {
                    fmt.Printf("%s battery exists :-> %v\n", name, *item.BatteryExists)
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
            b, _ := json.Marshal(item)
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
