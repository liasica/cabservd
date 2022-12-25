// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-03
// Based on cabservd by liasica, magicrolan@qq.com.

package service

import (
    "context"
    "github.com/auroraride/cabservd/internal/ent"
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

func (s *cabinetService) QueryCabinet(serial string) (*ent.Cabinet, error) {
    return s.orm.Query().Where(cabinet.Serial(serial)).First(s.ctx)
}

func (s *cabinetService) QueryAllCabinets() ent.Cabinets {
    items, _ := s.orm.Query().WithBins().All(s.ctx)
    return items
}
