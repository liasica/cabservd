// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on cabservd by liasica, magicrolan@qq.com.

package service

import (
    "github.com/auroraride/adapter/model"
    "github.com/auroraride/cabservd/internal/app"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/console"
    "github.com/auroraride/cabservd/internal/ent/scan"
    "github.com/auroraride/cabservd/internal/errs"
    "github.com/google/uuid"
    "net/http"
    "time"
)

type scanService struct {
    *BaseService

    orm *ent.ScanClient
}

func NewScan(params ...any) *scanService {
    return &scanService{
        BaseService: newService(params...),
        orm:         ent.Database.Scan,
    }
}

// Create 新增扫码记录
func (s *scanService) Create(serial string, cab *ent.Cabinet, data *model.ExchangeUsableResponse) *ent.Scan {
    sm, _ := ent.Database.Scan.Create().SetSerial(serial).SetUserID(s.User.ID).SetData(data).SetUserType(s.User.Type).SetCabinet(cab).Save(s.ctx)
    return sm
}

// Query 查询扫码记录
func (s *scanService) Query(str string) (*ent.Scan, error) {
    id, err := uuid.Parse(str)
    if err != nil {
        app.Panic(http.StatusBadRequest, errs.BadRequest)
    }
    return s.orm.Query().Where(scan.ID(id)).First(s.ctx)
}

// CensorX 获取并检查扫码是否有效
func (s *scanService) CensorX(str string, expires time.Duration) (sc *ent.Scan) {
    sc, _ = s.Query(str)
    if sc == nil {
        app.Panic(http.StatusBadRequest, errs.ExchangeTaskNotExist)
    }

    // 后续是否有该电柜扫码记录
    es, _ := s.orm.Query().Where(scan.CreatedAtGT(sc.CreatedAt), scan.CabinetID(sc.CabinetID)).Exist(s.ctx)

    // 后续是否有该电柜操作记录
    ec, _ := ent.Database.Console.Query().Where(console.CabinetID(sc.CabinetID), console.StartAtGT(sc.CreatedAt)).Exist(s.ctx)

    // 超时判定
    if es || ec || time.Now().After(sc.CreatedAt.Add(expires*time.Second)) {
        app.Panic(http.StatusBadRequest, errs.ExchangeExpired)
    }

    return
}
