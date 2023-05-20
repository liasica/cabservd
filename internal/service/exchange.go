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
	"github.com/jinzhu/copier"
	"github.com/liasica/go-helpers/silk"

	"github.com/auroraride/cabservd/internal/ent"
	"github.com/auroraride/cabservd/internal/ent/scan"
	"github.com/auroraride/cabservd/internal/g"
	"github.com/auroraride/cabservd/internal/sync"
	"github.com/auroraride/cabservd/internal/types"
)

type exchangeService struct {
	*app.BaseService
}

func NewExchange(params ...any) *exchangeService {
	return &exchangeService{
		BaseService: app.NewService(params...),
	}
}

// Usable 获取电柜待换电信息
func (s *exchangeService) Usable(req *cabdef.ExchangeUsableRequest) (res *cabdef.CabinetBinUsableResponse) {
	res = &cabdef.CabinetBinUsableResponse{
		Cabinet: new(cabdef.Cabinet),
		Fully:   new(cabdef.Bin),
		Empty:   new(cabdef.Bin),
	}

	// 检查电柜是否可换电
	cab := NewCabinet(app.PermissionNotRequired).OperableX(req.Serial)

	// 查询限定时间内其他扫码用户
	if exists, _ := ent.Database.Scan.Query().Where(
		scan.CabinetID(cab.ID),
		scan.UserIDNEQ(s.GetUser().ID),
		scan.CreatedAtGT(time.Now().Add(-time.Duration(req.Lock)*time.Second)),
	).Exist(s.GetContext()); exists {
		app.Panic(http.StatusBadRequest, adapter.ErrorCabinetBusy)
	}

	// 获取空仓和满电仓位
	fully, empty, err := NewCabinet(s.GetUser()).BusinessInfo(req.Model, cab, req.Minsoc, 1, 1)

	if err != nil {
		app.Panic(http.StatusBadRequest, err)
	}

	// // 如果无满电
	// if fully == nil {
	//     app.Panic(http.StatusBadRequest, adapter.ErrorCabinetNoFully)
	// }
	//
	// // 如果无空仓
	// if empty == nil {
	//     app.Panic(http.StatusBadRequest, adapter.ErrorCabinetNoEmpty)
	// }

	// 拷贝属性
	_ = copier.Copy(res.Cabinet, cab)
	_ = copier.Copy(res.Fully, fully)
	_ = copier.Copy(res.Empty, empty)

	// 存储扫码记录
	sm := NewScan(s.GetUser()).Create(adapter.BusinessExchange, req.Serial, cab, res)
	res.UUID = sm.UUID.String()

	return
}

func (s *exchangeService) Do(req *cabdef.ExchangeRequest) (res *cabdef.ExchangeResponse) {
	// 校验参数
	if req.Battery == "" && !g.Config.NonBms {
		res.Error = adapter.ErrorBatteryNotFound.Error()
		return
	}

	// 查询扫码记录
	sc := NewScan(s.GetUser()).CensorX(req.UUID, req.Timeout, req.Minsoc)

	// 检查电柜是否可换电
	_ = NewCabinet(s.GetUser()).OperableX(sc.Serial)

	// 开始同步换电流程
	results, err := s.start(req, sc)

	res = &cabdef.ExchangeResponse{
		Results: results,
	}

	for _, result := range results {
		res.Success = result.Step == 4 && result.Success

		// 取出的电池 (第三步成功即视为取走电池)
		if result.Step == 3 && result.Success {
			res.PutoutBattery = sc.Data.Fully.BatterySn
		}

		// 放入的电池, 第二步成功且有电池视为放入电池
		if result.Step <= 2 && result.BatterySN != "" {
			res.PutinBattery = result.BatterySN
		}
	}

	if err != nil {
		res.Error = err.Error()
	}

	return
}

func (s *exchangeService) start(req *cabdef.ExchangeRequest, sc *ent.Scan) (res []*cabdef.ExchangeStepMessage, err error) {
	defer func() {
		// 标记扫码失效
		_ = sc.Update().SetEfficient(false).Exec(s.GetContext())

		// TODO 任务标记???
	}()

	bins := []*cabdef.Bin{
		sc.Data.Empty,
		sc.Data.Fully,
	}

	cb := func(r *cabdef.BinOperateResult) {
		data := silk.Pointer(cabdef.ExchangeStepMessage(*r))
		res = append(res, data)
		// 异步发送结果
		go sync.SendMessage(data)
	}

	for i, conf := range types.ExchangeConfigure {
		b := bins[i]

		err = NewBin(s.GetUser()).Operate(&types.Bin{
			Timeout:      req.Timeout,
			MainOperate:  cabdef.OperateDoorOpen,
			Serial:       sc.Serial,
			UUID:         req.UUID,
			Ordinal:      b.Ordinal,
			Business:     adapter.BusinessExchange,
			Steps:        conf,
			Battery:      req.Battery,
			StepCallback: cb,
			Scan:         sc,
		})

		if err != nil {
			return
		}

	}

	return
}
