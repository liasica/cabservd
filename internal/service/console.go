// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on cabservd by liasica, magicrolan@qq.com.

package service

import (
    errs "github.com/auroraride/adapter/errors"
    "github.com/auroraride/adapter/model"
    "github.com/auroraride/cabservd/internal/app"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/console"
    "github.com/google/uuid"
    "github.com/liasica/go-helpers/silk"
    "time"
)

type consoleService struct {
    *BaseService

    orm *ent.ConsoleClient
}

func NewConsole(params ...any) *consoleService {
    return &consoleService{
        BaseService: newService(params...),
        orm:         ent.Database.Console,
    }
}

// StartExchangeStep 开始换电步骤
func (s *consoleService) StartExchangeStep(sc *ent.Scan, step model.ExchangeStep, open bool) (ec *ent.Console, b *ent.Bin, err error) {
    var (
        bid uint64
        op  *model.Operator
    )

    switch step {
    case model.ExchangeStepFirst, model.ExchangeStepSecond:
        bid = sc.Data.Empty.ID
    case model.ExchangeStepThird, model.ExchangeStepFourth:
        bid = sc.Data.Fully.ID
    }

    // 查询仓位信息
    b, err = NewBin(s.User).Query(bid)
    if err != nil {
        return
    }

    if open {
        op = silk.Pointer(model.OperatorBinOpen)
    }

    ec, err = s.orm.Create().
        SetNillableOperate(op).
        SetCabinetID(sc.CabinetID).
        SetSerial(sc.Serial).
        SetBinID(b.ID).
        SetUUID(sc.ID).
        SetType(console.TypeExchange).
        SetUserID(sc.UserID).
        SetUserType(sc.UserType).
        SetStep(step).
        SetStatus(console.StatusRunning).
        SetBeforeBin(b.Info()).
        SetStartAt(time.Now()).
        Save(s.ctx)

    return
}

func (s *consoleService) Operate(req *model.OperateRequest) (ec *ent.Console, b *ent.Bin, err error) {
    if req.Ordinal == nil {
        app.Panic(errs.CabinetBinOrdinalRequired)
    }

    // 查询仓位信息
    b, err = NewBin(s.User).QuerySerialOrdinal(req.Serial, *req.Ordinal)
    if err != nil {
        err = errs.CabinetBinNotFound
        return
    }

    ec, err = s.orm.Create().
        SetOperate(req.Type).
        SetCabinetID(b.CabinetID).
        SetBinID(b.ID).
        SetSerial(b.Serial).
        SetType(console.TypeOperate).
        SetUUID(uuid.New()).
        SetUserID(s.User.ID).
        SetUserType(s.User.Type).
        SetStatus(console.StatusRunning).
        SetStartAt(time.Now()).
        SetBeforeBin(b.Info()).Save(s.ctx)
    return
}

// Update 更新记录
func (s *consoleService) Update(ec *ent.Console, err error) (*ent.Console, *ent.Bin) {
    now := time.Now()
    cr := ec.Update().SetStopAt(now)
    if ec.StartAt != nil {
        cr.SetDuration(now.Sub(*ec.StartAt).Seconds())
    }

    // 查询最新仓位信息
    b, _ := NewBin(s.User).Query(ec.BinID)
    if b != nil {
        cr.SetAfterBin(b.Info())
    }

    if err != nil {
        cr.SetStatus(console.StatusFailed).SetMessage(err.Error())
    } else {
        cr.SetStatus(console.StatusSuccess)
    }

    ec, _ = cr.Save(s.ctx)
    return ec, b
}
