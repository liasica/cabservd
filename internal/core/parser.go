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
    jsoniter "github.com/json-iterator/go"
    "github.com/liasica/go-helpers/tools"
    log "github.com/sirupsen/logrus"
)

type Parser interface {
    Bins() ent.BinPointers
    Cabinet() ent.CabinetPointer
}

func UpdateCabinet(brand, serial string, p Parser) {
    SaveCabinet(brand, serial, p)
    SaveBins(brand, serial, p)
}

func SaveCabinet(brand, serial string, p Parser) {
    SaveCabinetContext(brand, serial, p, context.Background())
}

func SaveCabinetContext(brand, serial string, p Parser, ctx context.Context) {
    cab := p.Cabinet()
    err := ent.Database.Cabinet.Create().
        SetBrand(brand).
        SetSerial(serial).
        OnConflictColumns(cabinet.FieldSerial).
        Update(func(u *ent.CabinetUpsert) {
            // 状态
            if cab.Status != nil {
                u.SetStatus(*cab.Status)
            }

            // 经度
            if cab.Lng != nil {
                u.SetLng(*cab.Lng)
            }

            // 纬度
            if cab.Lat != nil {
                u.SetLat(*cab.Lat)
            }

            // GSM
            if cab.Gsm != nil {
                u.SetGsm(*cab.Gsm)
            }

            // 电压
            if cab.Voltage != nil {
                u.SetVoltage(*cab.Voltage)
            }

            // 电流
            if cab.Current != nil {
                u.SetCurrent(*cab.Current)
            }

            // 温度
            if cab.Temperature != nil {
                u.SetTemperature(*cab.Temperature)
            }

            // 启用
            if cab.Enable != nil {
                u.SetEnable(*cab.Enable)
            }

            // 总用电
            if cab.Electricity != nil {
                u.SetElectricity(*cab.Electricity)
            }
        }).Exec(ctx)
    if err != nil {
        b, _ := jsoniter.Marshal(cab)
        log.Errorf("电柜保存失败, %s: %v", string(b), err)
    }
}

func SaveBins(brand, serial string, p Parser) {
    ctx := context.Background()
    SaveBinsContext(brand, serial, p, ctx)
}

func SaveBinsContext(brand, serial string, p Parser, ctx context.Context) {
    items := p.Bins()
    if len(items) == 0 {
        return
    }
    for _, item := range items {
        uuid := tools.Md5String(fmt.Sprintf("%s_%s_%d", brand, serial, *item.Index))
        name := fmt.Sprintf("%d号仓", *item.Index+1)
        err := ent.Database.Bin.Create().
            SetUUID(uuid).
            SetBrand(brand).
            SetSerial(serial).
            SetName(name).
            SetIndex(*item.Index).
            OnConflictColumns(bin.FieldUUID).
            Update(func(u *ent.BinUpsert) {
                // 健康状态
                if item.Health != nil {
                    u.SetHealth(*item.Health)
                }

                // 仓门状态
                if item.Open != nil {
                    fmt.Printf("%s open:->%v\n", name, *item.Open)
                    u.SetOpen(*item.Open)
                }

                // 仓位启用状态
                if item.Enable != nil {
                    u.SetEnable(*item.Enable)
                }

                // 电池编号
                if item.BatterySn != nil {
                    fmt.Printf("%s battery:->%v\n", name, *item.BatterySn)
                    u.SetBatterySn(*item.BatterySn)
                    if *item.BatterySn == "" {
                        u.ResetBattery()
                    }
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
