// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-03
// Based on cabservd by liasica, magicrolan@qq.com.

package service

import (
    "fmt"
    "github.com/auroraride/adapter"
    "github.com/auroraride/adapter/defs/cabdef"
    log "github.com/auroraride/adapter/zlog"
    "github.com/auroraride/cabservd/internal/core"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/bin"
    "github.com/auroraride/cabservd/internal/ent/console"
    "github.com/auroraride/cabservd/internal/task"
    "github.com/auroraride/cabservd/internal/types"
    "github.com/google/uuid"
    "time"
)

type binService struct {
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

// Operate 按步骤操作某个仓位
// TODO 仓位检测
func (s *binService) Operate(bo *types.Bin) (err error) {
    if bo.StepCallback == nil {
        return adapter.ErrorBadRequest
    }

    // 查询仓位
    eb, _ := NewBin(s.User).QuerySerialOrdinal(bo.Serial, bo.Ordinal)
    if eb == nil {
        return adapter.ErrorBinNotFound
    }

    // TODO 是否有必要操作之前验证当前状态, 操作值 等于 当前状态时直接返回成功
    // TODO 其他详细日志

    fakevoltage, _ := core.Hub.Bean.GetEmptyDeviation()

    // 操作超时时间
    timeout := time.After(time.Duration(bo.Timeout) * time.Second)

    // 监听数据库变动
    ch := make(chan *ent.Bin)
    task.Bin.SetListener(eb, ch)

    stepper := make(chan *types.BinResult)

    defer func() {
        // 退出时删除监听
        task.Bin.RemoveListener(ch)
        close(stepper)

        // 判定是否成功以更新备注
        if err == nil && bo.BinRemark != nil {
            _ = s.orm.UpdateOneID(eb.ID).
                SetNillableRemark(bo.BinRemark).
                Exec(s.ctx)
        }
    }()

    // 开启监听
    go func() {
        for {
            select {
            case <-timeout:
                err = adapter.ErrorOperateTimeout
                stepper <- types.NewBinResult(nil, err)
                return
            case x := <-ch:
                // 如果通道关闭直接返回
                if x == nil {
                    return
                }

                // 更新仓位信息
                *eb = *x

                var doorOk, batteryOk, binOk bool

                step := bo.Current()

                switch step.Door {
                case cabdef.DetectDoorIgnore:
                    // 忽略仓门检测
                    doorOk = true
                case cabdef.DetectDoorOpen:
                    // 检测仓门是否开启
                    doorOk = x.Open
                case cabdef.DetectDoorClose:
                    // 检测仓门是否关闭
                    doorOk = !x.Open
                }

                switch step.Battery {
                case cabdef.DetectBatteryIgnore:
                    // 忽略电池检测
                    batteryOk = true
                case cabdef.DetectBatteryPutin:
                    // 严格检测电池是否放入
                    batteryOk = x.IsStrictHasBattery(fakevoltage)
                case cabdef.DetectBatteryPutout:
                    // 检测电池是否取出
                    batteryOk = x.IsLooseNoBattery(fakevoltage)
                }

                switch step.Bin {
                case cabdef.DetectBinIgnore:
                    // 忽略仓位检测
                    binOk = true
                case cabdef.DetectBinUsable:
                    binOk = x.IsUsable()
                    // 如果需要仓位可用但是仓位不可用, 直接发送任务失败并返回
                    if !binOk {
                        stepper <- types.NewBinResult(eb, adapter.ErrorBinNotUsable)
                        return
                    }
                case cabdef.DetectBinEnable:
                    binOk = x.Enable
                case cabdef.DetectBinDisable:
                    binOk = !x.Enable
                }

                if batteryOk && doorOk && binOk {
                    // 检查放入电池编号是否匹配
                    if step.Battery == cabdef.DetectBatteryPutin && bo.Battery != "" && eb.BatterySn != bo.Battery {
                        err = adapter.ErrorBatteryPutin
                    }

                    stepper <- types.NewBinResult(eb, err)

                    // 如果有错误, 终止
                    if err != nil {
                        return
                    }

                    // 尝试开启下次任务, 如果没有下次任务, 终止
                    if !bo.Next() {
                        return
                    }
                }
            }
        }
    }()

    for _, step := range bo.Steps {
        err = s.doOperateStep(bo.UUID, bo.Business, bo.Remark, eb, step, stepper, bo.StepCallback)

        // 遇到错误, 直接返回
        if err != nil {
            return
        }
    }

    return
}

// doOperateStep 按步骤操作
func (s *binService) doOperateStep(uid uuid.UUID, business adapter.Business, remark string, eb *ent.Bin, step *types.BinStep, stepper chan *types.BinResult, scb types.StepCallback) (err error) {
    // 创建记录
    var co *ent.Console
    co, err = ent.Database.Console.Create().
        SetBrand(eb.Brand).
        SetOperate(step.Operate).
        SetCabinetID(eb.CabinetID).
        SetBinID(eb.ID).
        SetSerial(eb.Serial).
        SetUserID(s.User.ID).
        SetUserType(s.User.Type).
        SetStatus(console.StatusRunning).
        SetStartAt(time.Now()).
        SetBeforeBin(eb.Info()).
        SetStep(step.Step).
        SetBusiness(business).
        SetUUID(uid).
        SetRemark(remark).
        Save(s.ctx)
    if err != nil {
        return
    }

    defer func() {
        res := NewConsole(s.User).Update(co, eb, err).OperateResult()
        log.Infof("<%s> [电柜: %s, 仓门: %d] { %s业务%s } 执行%v", s.User, eb.Serial, eb.Ordinal, business.Text(), step, adapter.Or[any](err == nil, "成功", fmt.Errorf("失败: %v", err)))

        // 同步回调结果
        scb(res)
    }()

    if step.Operate.IsCommand() {
        // 电柜控制
        err = core.Hub.Bean.SendControl(eb.Serial, step.Operate, eb.Ordinal)

        // TODO: 开仓失败后是否重复弹开逻辑???
        // TODO: 详细失败日志???
        if err != nil {
            log.Infof("[%s - %d] %s 失败: %v", eb.Serial, eb.Ordinal, step.Operate, err)
            return
        }
    }

    r := <-stepper
    _, err = r.Result()

    return
}

func (s *binService) BinInfo(req *cabdef.BinInfoRequest) (info *cabdef.BinInfo, err error) {
    b, _ := s.QuerySerialOrdinal(req.Serial, *req.Ordinal)
    if b == nil {
        err = adapter.ErrorBinNotFound
        return
    }
    info = b.Info()
    return
}
