// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-31
// Based on cabservd by liasica, magicrolan@qq.com.

package service

import (
    "github.com/auroraride/adapter"
)

type operateService struct {
    *BaseService
}

func NewOperate(params ...any) *operateService {
    return &operateService{
        BaseService: newService(params...),
    }
}

func (s *operateService) Do(req *adapter.OperateRequest) (res []*adapter.OperateStepResult) {
    // var (
    //     ec  *ent.Console
    //     err error
    // )

    // switch req.Operate {
    // default:
    //     app.Panic(adapter.ErrorOperateCommand)
    // case adapter.OperateDoorOpen, adapter.OperateBinEnable, adapter.OperateBinDisable:
    //     // 非业务操作
    //     ec, err = NewBin(s.ctx, s.User).Operate(req)
    //     if err != nil {
    //         app.Panic(err)
    //     }
    //
    //     res = append(res, ec.OperateResult())
    // case adapter.OperatePutin, adapter.OperatePutout:
    //     // 业务操作
    //     res = s.Business(req)
    // }

    return
}

func (s *operateService) Business(req *adapter.OperateRequest) (res []*adapter.OperateStepResult) {

    // NewBin(s.User).Operate(&adapter.OperateRequest{
    //     Serial:             req.Serial,
    //     Operate:            req.Operate,
    //     Timeout:            req.Timeout,
    //     Ordinal:            ordinal,
    //     VerifyPutinBattery: vb,
    // })

    return
}
