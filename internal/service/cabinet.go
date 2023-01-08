// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-30
// Based on cabservd by liasica, magicrolan@qq.com.

package service

import (
    "github.com/auroraride/adapter"
    "github.com/auroraride/cabservd/internal/core"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/bin"
    "github.com/auroraride/cabservd/internal/ent/cabinet"
    "strings"
)

type cabinetService struct {
    *BaseService

    orm *ent.CabinetClient
}

func NewCabinet(params ...any) *cabinetService {
    return &cabinetService{
        BaseService: newService(params...),
        orm:         ent.Database.Cabinet,
    }
}

func (s *cabinetService) Query(id uint64) (*ent.Cabinet, error) {
    return s.orm.Query().Where(cabinet.ID(id)).First(s.ctx)
}

func (s *cabinetService) QueryWithBin(id uint64) (*ent.Cabinet, error) {
    return s.orm.Query().Where(cabinet.ID(id)).WithBins(func(query *ent.BinQuery) {
        query.Order(ent.Asc(bin.FieldOrdinal))
    }).First(s.ctx)
}

func (s *cabinetService) QuerySerial(serial string) (*ent.Cabinet, error) {
    return s.orm.Query().Where(cabinet.Serial(serial)).First(s.ctx)
}

func (s *cabinetService) QuerySerialWithBin(serial string) (*ent.Cabinet, error) {
    return s.orm.Query().Where(cabinet.Serial(serial)).WithBins(func(query *ent.BinQuery) {
        query.Order(ent.Asc(bin.FieldOrdinal))
    }).First(s.ctx)
}

func (s *cabinetService) QuerySerialWithBinAll() ent.Cabinets {
    items, _ := s.orm.Query().WithBins(func(query *ent.BinQuery) {
        query.Order(ent.Asc(bin.FieldOrdinal))
    }).All(s.ctx)
    return items
}

func (s *cabinetService) QueryAllCabinet() ent.Cabinets {
    items, _ := s.orm.Query().All(s.ctx)
    return items
}

// UpdateStatus 更新电柜状态
func (s *cabinetService) UpdateStatus(serial string, status cabinet.Status) error {
    return s.orm.Update().Where(cabinet.Serial(serial)).SetStatus(status).Exec(s.ctx)
}

// DetectCabinet 验证电柜是否满足基本业务需求
func (s *cabinetService) DetectCabinet(cab *ent.Cabinet) error {
    if cab == nil {
        return adapter.ErrorCabinetNotFound
    }

    // 电柜需要空闲
    if cab.Status != cabinet.StatusIdle {
        return adapter.ErrorCabinetBusy
    }

    // 电柜需要在线
    if !cab.Online {
        return adapter.ErrorCabinetOffline
    }

    // 可办理业务的仓位至少有两个
    if len(cab.Edges.Bins) < 2 {
        return adapter.ErrorBinNotEnough
    }

    return nil
}

// BusinessInfo 获取业务仓位信息
// minsoc 指定最小电量 TODO 是否需要判定最小电量?
// minfull 指定最小满电仓位
// minempty 指定最小空仓位
func (s *cabinetService) BusinessInfo(bm string, cab *ent.Cabinet, minsoc float64, minbattery, minempty int) (fully, empty *ent.Bin, err error) {
    fakevoltage, fakecurrent := core.Hub.Bean.GetEmptyDeviation()

    var batteries, emptynum int

    for _, item := range cab.Edges.Bins {
        // 如果仓位未启用或仓位不健康直接跳过
        if !item.IsUsable() {
            continue
        }

        // 有正常未关闭的仓门直接报错
        if item.Open {
            err = adapter.ErrorCabinetDoorOpened
            return
        }

        // 判断电池型号
        if item.BatterySn != "" {
            bat := adapter.ParseBatterySN(item.BatterySn)
            if strings.ToUpper(bat.Model) != bm {
                continue
            }
        }

        // 判定是否可以满足业务
        switch true {
        case item.IsStrictHasBattery(fakevoltage):
            // 严格判定是否有电池
            batteries += 1
            // 若有电池
            // 获取满电仓位
            if fully == nil || fully.Soc < item.Soc {
                // 该仓位电量小于最小电量
                if item.Soc < minsoc {
                    continue
                }
                // 标定满仓
                fully = item
            }
        case item.IsStrictNoBattery(fakevoltage, fakecurrent):
            // 严格判定是否无电池
            emptynum += 1
            // 若无电池
            if empty == nil {
                empty = item
            }
        }
    }

    if batteries < minbattery {
        err = adapter.ErrorBatteryNotEnough
        return
    }

    if emptynum < minempty {
        err = adapter.ErrorBinNotEnough
        return
    }

    return
}
