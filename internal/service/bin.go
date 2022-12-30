// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-03
// Based on cabservd by liasica, magicrolan@qq.com.

package service

import (
    "fmt"
    errs "github.com/auroraride/adapter/errors"
    "github.com/auroraride/adapter/model"
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

func (s *binService) OpenDoor(serial string, ordinal int) (err error) {
    err = core.Hub.Bean.SendControl(serial, model.OperatorBinOpen, ordinal)
    message := "成功"
    if err != nil {
        message = fmt.Sprintf("失败, %v", err)
    }
    log.Infof("[BIN] [%s - %d, %s] 开门请求: %s", serial, ordinal, s.User, message)

    return
}

// Operate 控制仓位
func (s *binService) Operate(req *model.OperateRequest) (err error) {
    // TODO 操作的时候验证当前状态, 操作值 = 当前状态时返回报错信息

    // 记录日志
    var (
        ec *ent.Console
        b  *ent.Bin
        ch chan notice.IDSerialGetter
    )
    ec, b, err = NewConsole(s.User).Operate(req)
    if err != nil {
        err = errs.InternalServerError
        return
    }

    // 退出时更新日志
    defer func() {
        NewConsole(s.User).Update(ec, err)
        notice.Postgres.DeleteListener(ch)
    }()

    // 查找电柜
    var cab *ent.Cabinet
    cab, err = NewCabinet(s.User).QuerySerial(req.Serial)
    if err != nil {
        err = errs.CabinetNotFound
        return
    }

    if !cab.Online {
        err = errs.CabinetOffline
        return
    }

    // 监听状态改动
    ch = make(chan notice.IDSerialGetter)
    notice.Postgres.SetListener(notice.PostgresChannelBin, b.ID, ch)

    // 操作仓门
    err = core.Hub.Control(req)
    if err != nil {
        log.Errorf("[OPERATE] (%s) %v, 失败: %v", s.User, req, err)
        return
    }

    // 定义超时时间
    timeout := time.After(60 * time.Second)

    for {
        select {
        case x := <-ch:
            nb := x.(*ent.Bin)
            switch req.Type {
            case model.OperatorBinOpen:
                if nb.Open {
                    return
                }
            case model.OperatorBinEnable:
                if nb.Enable {
                    return
                }
            case model.OperatorBinDisable:
                if !nb.Enable {
                    return
                }
            }

        case <-timeout:
            // 超时
            err = errs.OperateTimeout
            return
        }
    }
}

// DetectUsable 检查仓位是否可用
func (s *binService) DetectUsable(b *ent.Bin) (ok model.Bool, message string) {
    ok = model.Bool(b.IsUsable())
    if ok {
        message = fmt.Sprintf("仓位可用")
    } else {
        message = fmt.Sprintf("仓位不可用(health: %v, enable: %v)", b.Health, b.Enable)
    }
    return
}

// DetectDoor 识别仓门开关状态
func (s *binService) DetectDoor(b *ent.Bin, d model.DetectDoor) (ok model.Bool, err error) {
    switch d {
    case model.DetectDoorOpen:
        ok = model.Bool(b.Open)
    case model.DetectDoorClose:
        ok = model.Bool(!b.Open)
    default:
        // 忽略
        ok = model.True
        return
    }

    usable, message := s.DetectUsable(b)

    if !usable {
        err = errs.CabinetBinNotUsable
    }

    log.Infof("[BIN] [%s - %d, %s] 仓门%s检测: %s, %s", b.Serial, b.Ordinal, s.User, d, ok, message)
    return
}

// DetectBattery 识别电池
func (s *binService) DetectBattery(b *ent.Bin, d model.DetectBattery) (ok model.Bool, err error) {
    fakevoltage, _ := core.Hub.Bean.GetEmptyDeviation()

    switch d {
    case model.DetectBatteryPutin:
        // 检测放入
        ok = model.Bool(b.IsStrictHasBattery(fakevoltage))
    case model.DetectBatteryPutout:
        // 检测取出
        ok = model.Bool(!b.IsStrictHasBattery(fakevoltage))
    default:
        // 忽略
        ok = model.True
        return
    }

    usable, message := s.DetectUsable(b)

    if !usable {
        err = errs.CabinetBinNotUsable
    }

    log.Infof("[BIN] [%s - %d, %s] 电池%s检测: %s, %s", b.Serial, b.Ordinal, s.User, d, ok, message)
    return
}
