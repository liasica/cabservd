// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on cabservd by liasica, magicrolan@qq.com.

package service

import (
	"net/http"
	"time"

	"github.com/auroraride/adapter"
	"github.com/auroraride/adapter/app"
	"github.com/auroraride/adapter/defs/cabdef"
	"github.com/oklog/ulid/v2"

	"github.com/auroraride/cabservd/internal/ent"
	"github.com/auroraride/cabservd/internal/ent/bin"
	"github.com/auroraride/cabservd/internal/ent/console"
	"github.com/auroraride/cabservd/internal/ent/scan"
	"github.com/auroraride/cabservd/internal/g"

	"github.com/google/uuid"
)

type scanService struct {
	*app.BaseService

	orm *ent.ScanClient
}

func NewScan(params ...any) *scanService {
	return &scanService{
		BaseService: app.NewService(params...),
		orm:         ent.Database.Scan,
	}
}

// Create 新增扫码记录
func (s *scanService) Create(ab adapter.Business, serial string, cab *ent.Cabinet, data *cabdef.CabinetBinUsableResponse) *ent.Scan {
	sm, _ := ent.Database.Scan.Create().
		SetSerial(serial).
		SetUserID(s.GetUser().ID).
		SetData(data).
		SetUserType(s.GetUser().Type).
		SetCabinet(cab).
		SetBusiness(ab).
		SetOrderNo(ulid.Make().String()).
		Save(s.GetContext())
	return sm
}

// QueryUUID 查询扫码记录
func (s *scanService) QueryUUID(uid uuid.UUID) (*ent.Scan, error) {
	// id, err := uuid.Parse(str)
	// if err != nil {
	//     app.Panic(http.StatusBadRequest, adapter.ErrorBadRequest)
	// }
	return s.orm.Query().Where(scan.UUID(uid)).First(s.GetContext())
}

// CensorX 获取并检查扫码是否有效
func (s *scanService) CensorX(uid uuid.UUID, expires int64, minsoc float64) (sc *ent.Scan) {
	sc, _ = s.QueryUUID(uid)
	if sc == nil || sc.Data == nil {
		app.Panic(http.StatusBadRequest, adapter.ErrorScanNotExist)
	}

	// 后续是否有该电柜扫码记录
	es, _ := s.orm.Query().Where(scan.CreatedAtGT(sc.CreatedAt), scan.CabinetID(sc.CabinetID)).Exist(s.GetContext())

	// 后续是否有该电柜操作记录
	ec, _ := ent.Database.Console.Query().Where(console.CabinetID(sc.CabinetID), console.StartAtGT(sc.CreatedAt)).Exist(s.GetContext())

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
	bins, _ := ent.Database.Bin.Query().Where(bin.IDIn(data.Fully.ID, data.Empty.ID)).All(s.GetContext())
	if len(bins) != 2 {
		app.Panic(http.StatusBadRequest, adapter.ErrorScanExpired)
	}

	// 验证是否可以换电
	for _, b := range bins {
		if !b.BusinessPossible(b.ID == data.Fully.ID, g.Fakevoltage, g.Fakecurrent, minsoc) {
			app.Panic(http.StatusBadRequest, adapter.ErrorExchangeCannot)
		}
	}
}
