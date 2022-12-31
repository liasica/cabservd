// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-03
// Based on cabservd by liasica, magicrolan@qq.com.

package service

import (
    "fmt"
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

func (s *binService) OpenDoor(serial string, ordinal int) (err error) {
    err = core.Hub.Bean.SendControl(serial, adapter.OperatorBinOpen, ordinal)
    message := "成功"
    if err != nil {
        message = fmt.Sprintf("失败, %v", err)
    }
    log.Infof("[BIN] [%s - %d, %s] 开门请求: %s", serial, ordinal, s.User, message)

    return
}

// Operate 控制仓位
func (s *binService) Operate(req *adapter.OperateRequest) (err error) {
    // TODO 操作的时候验证当前状态, 操作值 = 当前状态时返回报错信息

    // 记录日志
    var (
        ec *ent.Console
        b  *ent.Bin
        ch chan any
    )
    ec, b, err = NewConsole(s.User).Operate(req)
    if err != nil {
        err = adapter.ErrorInternalServer
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
        err = adapter.ErrorCabinetNotFound
        return
    }

    if !cab.Online {
        err = adapter.ErrorCabinetOffline
        return
    }

    // 监听状态改动
    ch = make(chan any)
    notice.Postgres.SetListener(pn.ChannelBin, b.ID, ch)

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
            case adapter.OperatorBinOpen:
                if nb.Open {
                    return
                }
            case adapter.OperatorBinEnable:
                if nb.Enable {
                    return
                }
            case adapter.OperatorBinDisable:
                if !nb.Enable {
                    return
                }
            }

        case <-timeout:
            // 超时
            err = adapter.ErrorOperateTimeout
            return
        }
    }
}

// DetectUsable 检查仓位是否可用
func (s *binService) DetectUsable(b *ent.Bin) (ok adapter.Bool, message string) {
    ok = adapter.Bool(b.IsUsable())
    if ok {
        message = fmt.Sprintf("仓位可用")
    } else {
        message = fmt.Sprintf("仓位不可用(health: %v, enable: %v)", b.Health, b.Enable)
    }
    return
}

// DetectDoor 识别仓门开关状态
func (s *binService) DetectDoor(b *ent.Bin, d adapter.DetectDoor) (ok adapter.Bool, err error) {
    switch d {
    case adapter.DetectDoorOpen:
        ok = adapter.Bool(b.Open)
    case adapter.DetectDoorClose:
        ok = adapter.Bool(!b.Open)
    default:
        // 忽略
        ok = adapter.True
        return
    }

    usable, message := s.DetectUsable(b)

    if !usable {
        err = adapter.ErrorCabinetBinNotUsable
    }

    log.Infof("[BIN] [%s - %d, %s] 仓门%s检测: %s, %s", b.Serial, b.Ordinal, s.User, d, ok, message)
    return
}

// DetectBattery 识别电池
func (s *binService) DetectBattery(b *ent.Bin, d adapter.DetectBattery) (ok adapter.Bool, err error) {
    fakevoltage, _ := core.Hub.Bean.GetEmptyDeviation()

    switch d {
    case adapter.DetectBatteryPutin:
        // 检测放入
        ok = adapter.Bool(b.IsStrictHasBattery(fakevoltage))
    case adapter.DetectBatteryPutout:
        // 检测取出
        ok = adapter.Bool(!b.IsStrictHasBattery(fakevoltage))
    default:
        // 忽略
        ok = adapter.True
        return
    }

    usable, message := s.DetectUsable(b)

    if !usable {
        err = adapter.ErrorCabinetBinNotUsable
    }

    log.Infof("[BIN] [%s - %d, %s] 电池%s检测: %s, %s", b.Serial, b.Ordinal, s.User, d, ok, message)
    return
}
