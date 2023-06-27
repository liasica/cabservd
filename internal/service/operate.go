// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-31
// Based on cabservd by liasica, magicrolan@qq.com.

package service

import (
	"net/http"
	"time"

	"github.com/auroraride/adapter"
	"github.com/auroraride/adapter/app"
	"github.com/auroraride/adapter/defs/cabdef"
	"github.com/google/uuid"
	"github.com/liasica/go-helpers/silk"

	"github.com/auroraride/cabservd/internal/ent"
	"github.com/auroraride/cabservd/internal/ent/bin"
	"github.com/auroraride/cabservd/internal/types"
)

type operateService struct {
	*app.BaseService
}

func NewOperate(params ...any) *operateService {
	return &operateService{
		BaseService: app.NewService(params...),
	}
}

func (s *operateService) do(req *cabdef.OperateBinRequest, steps types.BinSteps) (results []*cabdef.BinOperateResult) {
	if !req.Operate.IsCommand() {
		app.Panic(http.StatusBadRequest, adapter.ErrorOperateCommand)
	}

	var binRemark *string

	switch req.Operate {
	case cabdef.OperateBinDisable:
		binRemark = silk.String(req.Remark)
	case cabdef.OperateBinEnable:
		binRemark = silk.String("")
	}

	err := NewBin(s.GetUser()).Operate(&types.Bin{
		Timeout:     120,
		MainOperate: req.Operate,
		Serial:      req.Serial,
		UUID:        uuid.New(),
		Ordinal:     *req.Ordinal,
		Business:    adapter.BusinessOperate,
		Steps:       steps,
		Remark:      req.Remark,
		BinRemark:   binRemark,
		StepCallback: func(result *cabdef.BinOperateResult) {
			results = append(results, result)
		},
	})

	if err != nil {
		app.Panic(http.StatusBadRequest, err)
	}

	return
}

// Bin 单仓位控制
func (s *operateService) Bin(req *cabdef.OperateBinRequest) []*cabdef.BinOperateResult {
	// TODO 是否需要判定仓位状态
	// NewCabinet(s.GetUser()).OperableX(req.Serial)
	return s.do(req, types.OMOperates[req.Operate])
}

// BinOpenAndClose 开仓并等待关闭
func (s *operateService) BinOpenAndClose(req *cabdef.OperateBinRequest) (results []*cabdef.BinOperateResult) {
	if req.Operate != cabdef.OperateDoorOpen {
		app.Panic("指令错误")
	}

	results = s.do(req, types.OpenWaitCloseConfigure)

	// 等待3s查询一次仓位信息
	time.Sleep(3 * time.Second)

	if len(results) == 2 && results[1].After != nil {
		b, _ := ent.Database.Bin.Query().Where(bin.Serial(req.Serial), bin.Ordinal(*req.Ordinal)).First(s.GetContext())
		if b != nil {
			results[1].After = b.Info()
		}
	}
	return
}
