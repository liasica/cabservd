// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on cabservd by liasica, magicrolan@qq.com.

package service

import (
    "github.com/auroraride/adapter"
    "github.com/auroraride/cabservd/internal/app"
    "github.com/auroraride/cabservd/internal/core"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/bin"
    "github.com/auroraride/cabservd/internal/ent/console"
    "github.com/auroraride/cabservd/internal/ent/scan"
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
func (s *scanService) Create(ab adapter.Business, serial string, cab *ent.Cabinet, data *adapter.CabinetBinUsableResponse) *ent.Scan {
    sm, _ := ent.Database.Scan.Create().SetSerial(serial).SetUserID(s.User.ID).SetData(data).SetUserType(s.User.Type).SetCabinet(cab).SetBusiness(ab).Save(s.ctx)
    return sm
}

// QueryUUID 查询扫码记录
func (s *scanService) QueryUUID(uid uuid.UUID) (*ent.Scan, error) {
    // id, err := uuid.Parse(str)
    // if err != nil {
    //     app.Panic(http.StatusBadRequest, adapter.ErrorBadRequest)
    // }
    return s.orm.Query().Where(scan.UUID(uid)).First(s.ctx)
}

// CensorX 获取并检查扫码是否有效
func (s *scanService) CensorX(uid uuid.UUID, expires int64, minsoc float64) (sc *ent.Scan) {
    sc, _ = s.QueryUUID(uid)
    if sc == nil || sc.Data == nil {
        app.Panic(http.StatusBadRequest, adapter.ErrorScanNotExist)
    }

    // 后续是否有该电柜扫码记录
    es, _ := s.orm.Query().Where(scan.CreatedAtGT(sc.CreatedAt), scan.CabinetID(sc.CabinetID)).Exist(s.ctx)

    // 后续是否有该电柜操作记录
    ec, _ := ent.Database.Console.Query().Where(console.CabinetID(sc.CabinetID), console.StartAtGT(sc.CreatedAt)).Exist(s.ctx)

    // 超时判定
    if es || ec || !sc.Efficient || time.Now().After(sc.CreatedAt.Add(time.Duration(expires)*time.Second)) {
        app.Panic(http.StatusBadRequest, adapter.ErrorScanExpired)
    }

    // 如果业务是换电业务, 后续验证
    if sc.Business == adapter.BusinessExchange {
        s.ExchangeAbleX(sc, minsoc)
    }

    // TODO 验证其他业务

    return
}

// ExchangeAbleX 验证扫码结果是否可以换电
func (s *scanService) ExchangeAbleX(sc *ent.Scan, minsoc float64) {
    // 再次检查仓位是否正确
    data := sc.Data
    bins, _ := ent.Database.Bin.Query().Where(bin.IDIn(data.Fully.ID, data.Empty.ID)).All(s.ctx)
    if len(bins) != 2 {
        app.Panic(http.StatusBadRequest, adapter.ErrorScanExpired)
    }

    // 验证是否可以换电
    fakevoltage, fakecurrent := core.Hub.Bean.GetEmptyDeviation()
    for _, b := range bins {
        if !b.ExchangePossible(b.ID == data.Fully.ID, fakevoltage, fakecurrent, minsoc) {
            app.Panic(http.StatusBadRequest, adapter.ErrorExchangeCannot)
        }
    }
}
