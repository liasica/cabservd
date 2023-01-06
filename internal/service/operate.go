// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-31
// Based on cabservd by liasica, magicrolan@qq.com.

package service

import (
    "github.com/auroraride/adapter"
    "github.com/auroraride/adapter/defs/cabdef"
    "github.com/auroraride/cabservd/internal/app"
    "github.com/auroraride/cabservd/internal/types"
    "github.com/google/uuid"
    "github.com/liasica/go-helpers/silk"
    "net/http"
)

type operateService struct {
    *BaseService
}

func NewOperate(params ...any) *operateService {
    return &operateService{
        BaseService: newService(params...),
    }
}

// Bin 单仓位控制
func (s *operateService) Bin(req *cabdef.OperateBinRequest) (results []*cabdef.BinOperateResult) {
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

    err := NewBin(s.User).Operate(&types.Bin{
        Timeout:   120,
        Serial:    req.Serial,
        UUID:      uuid.New(),
        Ordinal:   *req.Ordinal,
        Business:  adapter.BusinessOperate,
        Steps:     types.OMOperates[req.Operate],
        Remark:    req.Remark,
        BinRemark: binRemark,
        StepCallback: func(result *cabdef.BinOperateResult) {
            results = append(results, result)
        },
    })

    if err != nil {
        app.Panic(http.StatusBadRequest, err)
    }

    return
}
