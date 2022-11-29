// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-21
// Based on cabservd by liasica, magicrolan@qq.com.

package core

import (
    "context"
    "fmt"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/bin"
    jsoniter "github.com/json-iterator/go"
    "github.com/liasica/go-helpers/tools"
    log "github.com/sirupsen/logrus"
)

func SaveBins(brand, sn string, bp BinParser) {
    ctx := context.Background()
    SaveBinsContext(brand, sn, bp, ctx)
}

func SaveBinsContext(brand, sn string, bp BinParser, ctx context.Context) {
    items := bp.Bins()
    for _, item := range items {
        uuid := tools.Md5String(fmt.Sprintf("%s_%s_%d", brand, sn, *item.Index))
        name := fmt.Sprintf("%d号仓", *item.Index+1)
        err := ent.Database.Bin.Create().
            SetUUID(uuid).
            SetBrand(brand).
            SetSn(sn).
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
        Where(bin.Sn(sn)).
        SetBatterySn("").
        SetSoc(0).
        SetSoh(0).
        SetVoltage(0).
        SetCurrent(0).
        // SetEnable(true). // TODO 是否单独设置LOCK
        SetOpen(false).
        Exec(context.Background())
}
