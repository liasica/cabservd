// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-01
// Based on cabservd by liasica, magicrolan@qq.com.

package types

import (
    "fmt"
    "github.com/auroraride/adapter"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/google/uuid"
)

var (
    // ExchangeConfigure 换电配置
    ExchangeConfigure = []BinSteps{
        {
            {
                Step:    1,
                Operate: adapter.OperateDoorOpen,
                Door:    adapter.DetectDoorOpen,
                Battery: adapter.DetectBatteryIgnore,
            },
            {
                Step:    2,
                Operate: adapter.OperateDetect,
                Door:    adapter.DetectDoorClose,
                Battery: adapter.DetectBatteryPutin,
            },
        },
        {
            {
                Step:    3,
                Operate: adapter.OperateDoorOpen,
                Door:    adapter.DetectDoorOpen,
                Battery: adapter.DetectBatteryIgnore,
            },
            {
                Step:    4,
                Operate: adapter.OperateDetect,
                Door:    adapter.DetectDoorClose,
                Battery: adapter.DetectBatteryPutout,
            },
        },
    }

    // PutinConfigure 电池放入配置
    PutinConfigure = BinSteps{
        {
            Step:    1,
            Operate: adapter.OperateDoorOpen,
            Door:    adapter.DetectDoorOpen,
            Battery: adapter.DetectBatteryIgnore,
        },
        {
            Step:    2,
            Operate: adapter.OperateDetect,
            Door:    adapter.DetectDoorClose,
            Battery: adapter.DetectBatteryPutin,
        },
    }

    // PutoutConfigure 电池取出配置
    PutoutConfigure = BinSteps{
        {
            Step:    1,
            Operate: adapter.OperateDoorOpen,
            Door:    adapter.DetectDoorOpen,
            Battery: adapter.DetectBatteryIgnore,
        },
        {
            Step:    2,
            Operate: adapter.OperateDetect,
            Door:    adapter.DetectDoorClose,
            Battery: adapter.DetectBatteryPutout,
        },
    }
)

type BinResult struct {
    err error
    bin *ent.Bin
}

func (r *BinResult) Result() (*ent.Bin, error) {
    return r.bin, r.err
}

func NewBinResult(eb *ent.Bin, err error) *BinResult {
    return &BinResult{
        err: err,
        bin: eb,
    }
}

type BinSteps []*BinStep

type BinStep struct {
    Step    int                   `json:"step"`    // 步骤序号
    Operate adapter.Operate       `json:"operate"` // 操作指令
    Door    adapter.DetectDoor    `json:"door"`    // 仓门检测
    Battery adapter.DetectBattery `json:"battery"` // 电池检测
}

func (b *BinStep) String() string {
    return fmt.Sprintf("第%d步, 指令: %s, 仓门: %s检测, 电池: %s检测", b.Step, b.Operate.Text(), b.Door.Text(), b.Battery.Text())
}

// Bin 仓位操控
type Bin struct {
    index int

    Timeout      int64            // 超时时间
    Serial       string           // 电柜编号
    UUID         uuid.UUID        // 任务ID
    Ordinal      int              // 仓位序号
    Business     adapter.Business // 业务类别
    Steps        BinSteps         // 规划步骤
    Battery      string           // 校验放入的电池编号
    StepCallback StepCallback     // 每一步的回调
}

// Current 获取当前步骤
func (o *Bin) Current() *BinStep {
    return o.Steps[o.index]
}

// Next 标记当前步骤完成, 开始下一个步骤
func (o *Bin) Next() bool {
    if o.index+1 > len(o.Steps) {
        return false
    }
    o.index += 1
    return true
}
