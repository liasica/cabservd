// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on cabservd by liasica, magicrolan@qq.com.

package service

import (
    errs "github.com/auroraride/adapter/errors"
    "github.com/auroraride/adapter/model"
    "github.com/auroraride/cabservd/internal/app"
    "github.com/auroraride/cabservd/internal/core"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/cabinet"
    "github.com/auroraride/cabservd/internal/ent/console"
    "github.com/auroraride/cabservd/internal/ent/scan"
    "github.com/auroraride/cabservd/internal/hook"
    "github.com/goccy/go-json"
    "github.com/jinzhu/copier"
    log "github.com/sirupsen/logrus"
    "net/http"
    "time"
)

type exchangeService struct {
    *BaseService
}

func NewExchange(params ...any) *exchangeService {
    return &exchangeService{
        BaseService: newService(params...),
    }
}

func (s *exchangeService) DetectCabinet(cab *ent.Cabinet) error {
    if cab == nil {
        return errs.CabinetNotFound
    }

    if cab.Status != cabinet.StatusIdle {
        return errs.CabinetBusy
    }

    if !cab.Online {
        return errs.CabinetOffline
    }

    if len(cab.Edges.Bins) < 2 {
        return errs.CabinetNoEmpty
    }

    return nil
}

// Usable 获取电柜待换电信息
func (s *exchangeService) Usable(req *model.ExchangeUsableRequest) (res *model.ExchangeUsableResponse) {
    res = &model.ExchangeUsableResponse{}

    defer func() {
        b, _ := json.Marshal(res)
        log.Infof("[SCAN] %s", b)
    }()

    // 获取电柜状态
    cab, _ := NewCabinet().QueryCabinetWithBin(req.Serial)

    // 检查电柜是否可换电
    err := s.DetectCabinet(cab)
    if err != nil {
        app.Panic(http.StatusBadRequest, err)
    }

    // 查询限定时间内其他扫码用户
    exists, _ := ent.Database.Scan.Query().Where(
        scan.CabinetID(cab.ID),
        scan.UserIDNEQ(s.User.ID),
        scan.CreatedAtGT(time.Now().Add(-req.Lock*time.Second)),
    ).Exist(s.ctx)
    if exists {
        app.Panic(http.StatusBadRequest, errs.CabinetBusy)
    }

    // TODO 查询是否有正在执行的任务?

    // 获取仓位
    var (
        fully, empty *ent.Bin
    )

    fakevoltage, fakecurrent := core.Hub.Bean.GetEmptyDeviation()
    for _, item := range cab.Edges.Bins {
        // 如果仓位未启用或仓位不健康直接跳过
        if !item.Enable || !item.Health {
            continue
        } else if item.Open {
            // 有正常未关闭的仓门直接报错
            app.Panic(http.StatusBadRequest, errs.CabinetDoorOpened)
        }
        // 获取满电仓位
        if fully == nil || fully.Soc < item.Soc {
            // 若该仓位无电池, 继续循环
            if !item.IsStrictHasBattery(fakevoltage) {
                // TODO 该仓位是否出错
                continue
            }
            // 该仓位电量小于最小电量
            if item.Soc < req.Minsoc {
                continue
            }
            // 标定满仓
            fully = item
        }
        if empty == nil {
            // 若该仓位无电池, 标记为空仓
            if !item.IsLooseHasBattery(fakevoltage, fakecurrent) {
                empty = item
            }
        }
    }

    // 如果无满电
    if fully == nil {
        app.Panic(http.StatusBadRequest, errs.CabinetNoFully)
    }

    // 如果无空仓
    if empty == nil {
        app.Panic(http.StatusBadRequest, errs.CabinetNoEmpty)
    }

    // 拷贝属性
    _ = copier.Copy(&res.Cabinet, cab)
    _ = copier.Copy(&res.Fully, fully)
    _ = copier.Copy(&res.Empty, empty)

    // 存储扫码记录
    sm := NewScan(s.User).Create(req.Serial, cab, res)
    res.UUID = sm.ID.String()

    return
}

func (s *exchangeService) Do(req *model.ExchangeRequest) (res *model.ExchangeResponse) {
    // 查询扫码记录
    sc := NewScan(s.User).CensorX(req)

    // 开始同步换电流程
    results, err := s.start(req, sc)

    return &model.ExchangeResponse{
        Results: results,
        Error:   err.Error(),
    }
}

func (s *exchangeService) start(req *model.ExchangeRequest, sc *ent.Scan) (results []*model.ExchangeStepResult, err error) {
    var (
        // bin service
        bs = NewBin(s.User)

        // 仓位
        eb *ent.Bin

        // 记录
        ec *ent.Console
    )

    cab, _ := NewCabinet().Query(sc.CabinetID)

    // 检查电柜是否可换电
    err = s.DetectCabinet(cab)
    if err != nil {
        return
    }

    // 标记电柜为换电中
    _ = cab.Update().SetStatus(cabinet.StatusExchange).Exec(s.ctx)

    defer func() {
        // 删除监听
        hook.Postgres.DeleteBinListener(eb.ID)

        // 标记电柜为空闲
        _ = cab.Update().SetStatus(cabinet.StatusIdle).Exec(s.ctx)

        // TODO 任务标记

        // 更新记录
        cr := ec.Update().SetStopAt(time.Now()).SetAfterBin(eb.Info())
        if err != nil {
            cr.SetStatus(console.StatusFailed).SetMessage(err.Error())
        } else {
            cr.SetStatus(console.StatusSuccess)
        }

        ec, _ = cr.Save(s.ctx)

        results = append(results, ec.StepResult())
    }()

    for _, conf := range model.ExchangeStepConfigures {

        // 创建换电步骤
        ec, eb, err = NewConsole(s.User).StartExchangeStep(sc, conf.Step)
        if err != nil {
            return
        }

        ch := make(chan *ent.Bin)
        hook.Postgres.SetBinListener(eb.ID, ch)

        // 如果需要开仓
        if conf.Door == model.DetectDoorOpen {
            // 开仓
            err = bs.OpenDoor(sc.Serial, eb.Ordinal)
            if err != nil {
                return
            }
        }

        // TODO: 开仓失败后是否重复弹开逻辑???

        var batteryOk, doorOk model.Bool

        for {
            select {
            case b := <-ch:
                // 检查仓位是否满足条件
                doorOk, err = bs.DetectDoor(b, conf.Door)
                if err != nil {
                    return
                }

                // 检查电池是否满足条件
                batteryOk, err = bs.DetectBattery(b, conf.Battery)
                if err != nil {
                    return
                }

                if doorOk && batteryOk {
                    // 更新仓位信息
                    *eb = *b
                    return
                }

            case <-time.After(req.TimeOut * time.Second):
                // 超时
                err = errs.ExchangeTimeOut
                return
            }
        }
    }

    return
}
