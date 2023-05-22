// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-19
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

import (
	"context"
	"time"

	"github.com/auroraride/adapter"
	"github.com/auroraride/adapter/defs/cabdef"
	"github.com/liasica/go-helpers/silk"

	"github.com/auroraride/cabservd/internal/ent"
	"github.com/auroraride/cabservd/internal/ent/bin"
	"github.com/auroraride/cabservd/internal/ent/console"
	"github.com/auroraride/cabservd/internal/types"
)

var (
	binCommand = map[cabdef.Operate]CellCommand{
		cabdef.OperateDoorOpen:   CellOpenDoor,
		cabdef.OperateBinDisable: CellForbid,
		cabdef.OperateBinEnable:  CellUnForbid,
	}
)

// BinOperate 运维操作仓位
func BinOperate(bo *types.Bin, step *types.BinStep) (err error) {
	if step.Step != 1 {
		return
	}
	_, err = fetchCellCommand(&CellCommandRequest{
		Sn:      bo.Serial,
		CellNos: []int{bo.Ordinal},
		Command: binCommand[step.Operate],
	})
	return
}

// BinBusiness 执行仓位业务
func BinBusiness(user *adapter.User, sc *ent.Scan, batterySN string, cb types.StepCallback) (err error) {
	orderNo := *sc.OrderNo

	// 缓冲大小
	size := 2
	// 只有换电操作是4步
	if sc.Business == adapter.BusinessExchange {
		size = 4
	}

	// 监听器
	notifier := make(chan *BusinessNotify, size)
	storeBizNotifier(orderNo, notifier)

	// 退出监听器
	quit := make(chan struct{})

	// 关闭所有channel
	defer func() {
		adapter.ChSafeClose(quit)
		removeBizNotifier(orderNo, notifier)
	}()

	go func() {
		for {
			select {
			case data := <-notifier:
				if data == nil {
					continue
				}

				var stop bool
				var result *cabdef.BinOperateResult
				stop, result, err = bizStep(user, sc, data)

				if result != nil {
					// 回调结果
					go cb(result)
				}

				// 如果是最终步骤或者发生错误
				if stop {
					adapter.ChSafeSend(quit, struct{}{})
					return
				}
			}
		}
	}()

	switch sc.Business {
	case adapter.BusinessExchange:
		_, err = fetchExchange(&bizExchangeRequest{
			Sn:               sc.Serial,
			OrderNo:          orderNo,
			EmptyCellNo:      sc.Data.Empty.Ordinal,
			BatteryCellNo:    sc.Data.Fully.Ordinal,
			BindingBatterySn: batterySN,
		})
	}

	// 请求错误直接退出
	if err != nil {
		adapter.ChSafeSend(quit, struct{}{})
	}

	<-quit
	return
}

func bizStep(user *adapter.User, sc *ent.Scan, data *BusinessNotify) (stop bool, result *cabdef.BinOperateResult, err error) {
	// 时间
	now := time.Now()
	start := time.UnixMilli(data.OperateTime)
	d := now.Sub(start).Seconds()

	// 执行前后仓位状态
	var before *cabdef.Bin

	// 步骤结果
	var status bizStatus
	switch sc.Business {
	case adapter.BusinessExchange:
		status = exchangeStatus(data.OperateStep)
		if status.step() < 2 {
			before = sc.Data.Empty
		} else {
			before = sc.Data.Fully
		}
	default:
		before = sc.Data.BusinessBin
	}

	err = status.error()

	// 是否终止
	// 如果是最后一步 或者 发生错误
	stop = status.last() || err != nil
	step := status.step()

	// 有可能步骤为0
	if step == 0 {
		return
	}

	bs := status.binStep()

	ctx := context.Background()

	// TODO 主动从西六楼查询电池最新状态
	// 查询仓位状态
	after, _ := ent.Database.Bin.Query().Where(bin.ID(before.ID)).First(ctx)

	result = &cabdef.BinOperateResult{
		UUID:     sc.UUID.String(),
		Step:     status.step(),
		Business: sc.Business,
		StartAt:  silk.Pointer(start),
		StopAt:   silk.Pointer(now),
		Success:  err == nil,
		Before:   before.Info(),
		After:    after.Info(),
		Duration: d,
	}

	if err != nil {
		result.Message = err.Error()
	}

	// 存储步骤信息
	go func() {
		creator := ent.Database.Console.Create().
			SetOperate(bs.Operate).
			SetCabinetID(sc.CabinetID).
			SetBinID(before.ID).
			SetSerial(sc.Serial).
			SetUserID(user.ID).
			SetUserType(user.Type).
			SetStartAt(start).
			SetBeforeBin(before.Info()).
			SetStep(step).
			SetBusiness(sc.Business).
			SetUUID(sc.UUID).
			SetDuration(d).
			SetStopAt(now)
		if err != nil {
			creator.SetStatus(console.StatusFailed).SetMessage(err.Error())
		} else {
			creator.SetStatus(console.StatusSuccess)
		}
		_, _ = creator.Save(ctx)
	}()

	return
}
