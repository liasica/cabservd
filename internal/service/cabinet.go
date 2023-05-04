// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-30
// Based on cabservd by liasica, magicrolan@qq.com.

package service

import (
	"net/http"
	"strings"

	"github.com/auroraride/adapter"
	"github.com/auroraride/adapter/app"
	"go.uber.org/zap"

	"github.com/auroraride/cabservd/internal/core"
	"github.com/auroraride/cabservd/internal/ent"
	"github.com/auroraride/cabservd/internal/ent/bin"
	"github.com/auroraride/cabservd/internal/ent/cabinet"
	"github.com/auroraride/cabservd/internal/ent/console"
	"github.com/auroraride/cabservd/internal/g"
)

type cabinetService struct {
	*app.BaseService

	orm *ent.CabinetClient
}

func NewCabinet(params ...any) *cabinetService {
	return &cabinetService{
		BaseService: app.NewService(params...),
		orm:         ent.Database.Cabinet,
	}
}

func (s *cabinetService) Query(id uint64) (*ent.Cabinet, error) {
	return s.orm.Query().Where(cabinet.ID(id)).First(s.GetContext())
}

func (s *cabinetService) QueryWithBin(id uint64) (*ent.Cabinet, error) {
	return s.orm.Query().Where(cabinet.ID(id)).WithBins(func(query *ent.BinQuery) {
		query.Order(ent.Asc(bin.FieldOrdinal))
	}).First(s.GetContext())
}

func (s *cabinetService) QuerySerial(serial string) (*ent.Cabinet, error) {
	return s.orm.Query().Where(cabinet.Serial(serial)).First(s.GetContext())
}

func (s *cabinetService) QuerySerialWithBin(serial string) (*ent.Cabinet, error) {
	return s.orm.Query().Where(cabinet.Serial(serial)).WithBins(func(query *ent.BinQuery) {
		query.Order(ent.Asc(bin.FieldOrdinal))
	}).First(s.GetContext())
}

func (s *cabinetService) QuerySerialWithBinAll() ent.Cabinets {
	items, _ := s.orm.Query().WithBins(func(query *ent.BinQuery) {
		query.Order(ent.Asc(bin.FieldOrdinal))
	}).All(s.GetContext())
	return items
}

func (s *cabinetService) QueryAllCabinet() ent.Cabinets {
	items, _ := s.orm.Query().All(s.GetContext())
	return items
}

// UpdateStatus 更新电柜状态
func (s *cabinetService) UpdateStatus(serial string, status cabinet.Status) error {
	return s.orm.Update().Where(cabinet.Serial(serial)).SetStatus(status).Exec(s.GetContext())
}

// Operable 验证电柜是否满足基本业务操作需求
func (s *cabinetService) Operable(serial string) (cab *ent.Cabinet, err error) {
	// 查找电柜和仓位
	cab, _ = NewCabinet(app.PermissionNotRequired).QuerySerialWithBin(serial)

	if cab == nil {
		return nil, adapter.ErrorCabinetNotFound
	}

	// 电柜非正常
	if cab.Status != cabinet.StatusNormal {
		return nil, adapter.ErrorCabinetAbnormal
	}

	// 电柜需要在线
	if !cab.Online {
		return nil, adapter.ErrorCabinetOffline
	}

	// 可办理业务的仓位至少有两个
	if len(cab.Edges.Bins) < 2 {
		return nil, adapter.ErrorBinNotEnough
	}

	// 查询是否有正在执行的任务
	if exists, _ := ent.Database.Console.Query().Where(console.Serial(serial), console.Or(console.StatusIn(console.StatusRunning))).Exist(s.GetContext()); exists {
		return nil, adapter.ErrorCabinetBusy
	}

	return
}

func (s *cabinetService) OperableX(serial string) *ent.Cabinet {
	cab, err := NewCabinet(app.PermissionNotRequired).Operable(serial)
	if err != nil {
		app.Panic(http.StatusBadRequest, err)
	}
	return cab
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

		// 如果有bms通讯, 判断电池编码和型号
		if item.BatterySn != "" && !g.Config.NonBms {
			var bat adapter.Battery
			bat, err = adapter.ParseBatterySN(item.BatterySn)
			if err != nil {
				zap.L().Error("电池编码错误: "+item.BatterySn, zap.Error(err))
				continue
			}
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
			// 标定满仓
			switch {
			case fully == nil, fully.Soc <= item.Soc:
				// 如果满电标定为空 或 满电标定电量小于该仓位电量
				fully = item
			case fully != nil && g.Config.NonBms && fully.Voltage < item.Voltage:
				// 非智能柜独有逻辑: 如果满电标定不为空但满电标定电压小于该仓位电压 -- 2023年04月23日15:04:41 /  2023年05月04日14:52:22 曹博文提出
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
