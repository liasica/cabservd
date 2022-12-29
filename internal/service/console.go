// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on cabservd by liasica, magicrolan@qq.com.

package service

import (
    "github.com/auroraride/adapter/model"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/console"
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

// CreateExchangeSteps 创建换电步骤
// func (s *consoleService) CreateExchangeSteps(sc *ent.Scan) (cos ent.Consoles, err error) {
//     var bid uint64
//     bulks := make([]*ent.ConsoleCreate, 4)
//     for _, step := range model.ExchangeSteps {
//         switch step {
//         case model.ExchangeStepFirst, model.ExchangeStepSecond:
//             bid = sc.Data.Empty.ID
//         case model.ExchangeStepThird, model.ExchangeStepFourth:
//             bid = sc.Data.Fully.ID
//         }
//         v := s.orm.Create().
//             SetCabinetID(sc.CabinetID).
//             SetUUID(sc.ID).
//             SetType(console.TypeExchange).
//             SetUserID(sc.UserID).
//             SetUserType(sc.UserType).
//             SetStep(step).
//             SetStatus(console.StatusPending).
//             SetBinID(bid)
//         bulks[step.Index()] = v
//     }
//
//     return s.orm.CreateBulk(bulks...).Save(s.ctx)
// }

// StartExchangeStep 开始换电步骤
func (s *consoleService) StartExchangeStep(sc *ent.Scan, step model.ExchangeStep) (ec *ent.Console, b *ent.Bin, err error) {
    var (
        bid uint64
    )

    switch step {
    case model.ExchangeStepFirst, model.ExchangeStepSecond:
        bid = sc.Data.Empty.ID
    case model.ExchangeStepThird, model.ExchangeStepFourth:
        bid = sc.Data.Fully.ID
    }

    // 查询最新仓位信息
    b, err = NewBin().Query(bid)
    if err != nil {
        return
    }

    ec, err = s.orm.Create().
        SetCabinetID(sc.CabinetID).
        SetSerial(sc.Serial).
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
