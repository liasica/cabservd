// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on cabservd by liasica, magicrolan@qq.com.

package service

import (
	"time"

	"github.com/auroraride/adapter/app"

	"github.com/auroraride/cabservd/internal/ent"
	"github.com/auroraride/cabservd/internal/ent/console"
	"github.com/auroraride/cabservd/internal/types"
)

type consoleService struct {
	*app.BaseService

	orm *ent.ConsoleClient
}

func NewConsole(params ...any) *consoleService {
	return &consoleService{
		BaseService: app.NewService(params...),
		orm:         ent.Database.Console,
	}
}

// Create 创建操作记录
func (s *consoleService) Create(bo *types.Bin, step *types.BinStep, eb *ent.Bin) (*ent.Console, error) {
	return ent.Database.Console.Create().
		SetOperate(step.Operate).
		SetCabinetID(eb.CabinetID).
		SetBinID(eb.ID).
		SetSerial(eb.Serial).
		SetUserID(s.GetUser().ID).
		SetUserType(s.GetUser().Type).
		SetStatus(console.StatusRunning).
		SetStartAt(time.Now()).
		SetBeforeBin(eb.Info()).
		SetStep(step.Step).
		SetBusiness(bo.Business).
		SetUUID(bo.UUID).
		SetRemark(bo.Remark).
		Save(s.GetContext())
}

// Update 更新记录
func (s *consoleService) Update(ec *ent.Console, b *ent.Bin, times int, err error) *ent.Console {
	now := time.Now()
	cr := ec.Update().SetStopAt(now).SetCommandRetryTimes(times)
	if ec.StartAt != nil {
		cr.SetDuration(now.Sub(*ec.StartAt).Seconds())
	}

	// 仓位信息
	if b != nil {
		cr.SetAfterBin(b.Info())
	}

	if err != nil {
		cr.SetStatus(console.StatusFailed).SetMessage(err.Error())
	} else {
		cr.SetStatus(console.StatusSuccess)
	}

	ec, _ = cr.Save(s.GetContext())
	return ec
}

// InJob 查询是否有正在执行的任务
func (s *consoleService) InJob(serial string) bool {
	exists, _ := s.orm.Query().Where(console.Serial(serial), console.Or(console.StatusIn(console.StatusRunning))).Exist(s.GetContext())
	return exists
}
