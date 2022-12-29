// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-03
// Based on cabservd by liasica, magicrolan@qq.com.

package service

import (
    "context"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/bin"
    "github.com/auroraride/cabservd/internal/ent/cabinet"
)

type cabinetService struct {
    ctx context.Context
    orm *ent.CabinetClient
}

func NewCabinet() *cabinetService {
    return &cabinetService{
        ctx: context.Background(),
        orm: ent.Database.Cabinet,
    }
}

func (s *cabinetService) Query(id uint64) (*ent.Cabinet, error) {
    return s.orm.Query().Where(cabinet.ID(id)).First(s.ctx)
}

func (s *cabinetService) QueryCabinet(serial string) (*ent.Cabinet, error) {
    return s.orm.Query().Where(cabinet.Serial(serial)).First(s.ctx)
}

func (s *cabinetService) QueryCabinetWithBin(serial string) (*ent.Cabinet, error) {
    return s.orm.Query().Where(cabinet.Serial(serial)).WithBins(func(query *ent.BinQuery) {
        query.Order(ent.Asc(bin.FieldOrdinal))
    }).First(s.ctx)
}

func (s *cabinetService) QueryAllCabinetWithBin() ent.Cabinets {
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
