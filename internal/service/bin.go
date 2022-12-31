// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-03
// Based on cabservd by liasica, magicrolan@qq.com.

package service

import (
    "github.com/auroraride/adapter"
    "github.com/auroraride/adapter/pn"
    "github.com/auroraride/cabservd/internal/core"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/bin"
    "github.com/auroraride/cabservd/internal/notice"
    log "github.com/sirupsen/logrus"
    "time"
)

type binService struct {
    // ordinal int
    // bin     *ent.Bin
    // cabinet *ent.Cabinet
    *BaseService
    orm *ent.BinClient
}

func NewBin(params ...any) *binService {
    return &binService{
        BaseService: newService(params...),
        orm:         ent.Database.Bin,
    }
}

func (s *binService) QueryAllBin() ent.Bins {
    items, _ := s.orm.Query().All(s.ctx)
    return items
}

func (s *binService) Query(id uint64) (*ent.Bin, error) {
    return s.orm.Query().Where(bin.ID(id)).First(s.ctx)
}

func (s *binService) QuerySerialOrdinal(serial string, ordinal int) (*ent.Bin, error) {
    return s.orm.Query().Where(bin.Serial(serial), bin.Ordinal(ordinal)).First(s.ctx)
}

func (s *binService) Operate(req *adapter.OperateRequest) (ec *ent.Console, err error) {
    if req.Ordinal == nil {
        err = adapter.ErrorCabinetBinOrdinalRequired
        return
    }

    var (
        eb *ent.Bin
        cs = NewConsole(s.User)
    )

    // 创建记录
    ec, eb, err = cs.Create(req)
    if err != nil {
        return
    }

    ch := make(chan any)
    notice.Postgres.SetListener(pn.ChannelBin, eb.ID, ch)

    defer func() {
        // 删除监听
        notice.Postgres.DeleteListener(ch)

        // 更新记录
        ec = cs.Update(ec, eb, err)
    }()

    // 是否跳过发送仓控指令, 例如: 检查电池是否放入
    var skipSend bool

    // 操作之前验证当前状态, 操作值 等于 当前状态时直接返回成功
    // TODO 是否有必要?
    // TODO 其他详细日志
    switch req.Operate {
    default:
        err = adapter.ErrorOperate
        skipSend = true
        return
    case adapter.OperateDoorOpen:
        if eb.Open {
            log.Info(adapter.ErrorBinOpened)
            return
        }
    case adapter.OperateBinDisable:
        if !eb.Enable {
            log.Info(adapter.ErrorBinDisabled)
            return
        }
    case adapter.OperateBinEnable:
        if eb.Enable {
            log.Info(adapter.ErrorBinEnabled)
            return
        }
    }

    if !skipSend {
        err = core.Hub.Bean.SendControl(req.Serial, req.Operate, *req.Ordinal)
        // TODO: 开仓失败后是否重复弹开逻辑???
    }

    fakevoltage, _ := core.Hub.Bean.GetEmptyDeviation()

    // 定义超时时间
    timeout := time.After(time.Duration(req.Timeout) * time.Second)

    for {
        select {
        case x := <-ch:
            // 更新仓位信息
            *eb = *x.(*ent.Bin)

            var (
                ok    adapter.Bool
                check bool // 是否检查仓位是否可用
            )

            // 检查是否成功
            switch req.Operate {
            case adapter.OperateDoorOpen:
                // 检查仓门是否开启
                ok = adapter.Bool(eb.Open)
            case adapter.OperateBinEnable:
                // 检查仓位是否启用, 不进行后续仓位健康检查
                ok = adapter.Bool(eb.Enable)
                check = false
            case adapter.OperateBinDisable:
                // 检查仓位是否禁用, 不进行后续仓位健康检查
                ok = adapter.Bool(!eb.Enable)
                ok = adapter.Bool(eb.Enable)
            case adapter.OperatePutin:
                // 检测放入, 仓内有电池并且仓门关闭
                ok = adapter.Bool(eb.IsStrictHasBattery(fakevoltage) && !eb.Open)
            case adapter.OperatePutout:
                // 检测取出, 仓内无电池且仓门关闭
                ok = adapter.Bool(eb.IsLooseNoBattery(fakevoltage) && !eb.Open)
            }

            log.Infof("{ %s } 结果: %v", req.String(), ok)

            // 检查仓位可用状态
            if check && !eb.IsUsable() {
                err = adapter.ErrorCabinetBinNotUsable
                return
            }

            if ok {
                // 检查放入电池是否匹配
                if req.Operate == adapter.OperatePutin && eb.BatterySn != req.VerifyPutinBattery {
                    err = adapter.ErrorBatteryPutin
                }
                return
            }

        case <-timeout:
            // 超时
            err = adapter.ErrorExchangeTimeOut
            return
        }
    }
}
