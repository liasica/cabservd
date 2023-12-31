// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-01
// Based on cabservd by liasica, magicrolan@qq.com.

package types

import (
	"fmt"

	"github.com/auroraride/adapter"
	"github.com/auroraride/adapter/defs/cabdef"
	"github.com/google/uuid"

	"github.com/auroraride/cabservd/internal/ent"
)

var (
	// ExchangeConfigure 换电配置
	ExchangeConfigure = []BinSteps{
		{
			{
				Step:    1,
				Operate: cabdef.OperateDoorOpen,
				Door:    cabdef.DetectDoorOpen,
				Battery: cabdef.DetectBatteryIgnore,
				Bin:     cabdef.DetectBinIgnore,
			},
			{
				Step:    2,
				Operate: cabdef.OperateDetect,
				Door:    cabdef.DetectDoorClose,
				Battery: cabdef.DetectBatteryPutin,
				Bin:     cabdef.DetectBinIgnore,
			},
		},
		{
			{
				Step:    3,
				Operate: cabdef.OperateDoorOpen,
				Door:    cabdef.DetectDoorOpen,
				Battery: cabdef.DetectBatteryIgnore,
				Bin:     cabdef.DetectBinIgnore,
			},
			{
				Step:    4,
				Operate: cabdef.OperateDetect,
				Door:    cabdef.DetectDoorClose,
				Battery: cabdef.DetectBatteryPutout,
				Bin:     cabdef.DetectBinIgnore,
			},
		},
	}

	// PutinConfigure 电池放入配置
	PutinConfigure = BinSteps{
		{
			Step:    1,
			Operate: cabdef.OperateDoorOpen,
			Door:    cabdef.DetectDoorOpen,
			Battery: cabdef.DetectBatteryIgnore,
			Bin:     cabdef.DetectBinUsable,
		},
		{
			Step:    2,
			Operate: cabdef.OperateDetect,
			Door:    cabdef.DetectDoorClose,
			Battery: cabdef.DetectBatteryPutin,
			Bin:     cabdef.DetectBinUsable,
		},
	}

	// PutoutConfigure 电池取出配置
	PutoutConfigure = BinSteps{
		{
			Step:    1,
			Operate: cabdef.OperateDoorOpen,
			Door:    cabdef.DetectDoorOpen,
			Battery: cabdef.DetectBatteryIgnore,
			Bin:     cabdef.DetectBinUsable,
		},
		{
			Step:    2,
			Operate: cabdef.OperateDetect,
			Door:    cabdef.DetectDoorClose,
			Battery: cabdef.DetectBatteryPutout,
			Bin:     cabdef.DetectBinUsable,
		},
	}

	// OpenWaitCloseConfigure 开仓并等待关仓
	OpenWaitCloseConfigure = BinSteps{
		{
			Step:    1,
			Operate: cabdef.OperateDoorOpen,
			Door:    cabdef.DetectDoorOpen,
			Battery: cabdef.DetectBatteryIgnore,
			Bin:     cabdef.DetectBinUsable,
		},
		{
			Step:    2,
			Operate: cabdef.OperateDetect,
			Door:    cabdef.DetectDoorClose,
			Battery: cabdef.DetectBatteryIgnore,
			Bin:     cabdef.DetectBinIgnore,
		},
	}

	// OMOpenConfigure 运维开仓
	OMOpenConfigure = BinSteps{
		{
			Step:    1,
			Operate: cabdef.OperateDoorOpen,
			Door:    cabdef.DetectDoorOpen,
			Battery: cabdef.DetectBatteryIgnore,
			Bin:     cabdef.DetectBinIgnore,
		},
	}

	// OMEnableConfigure 运维启用
	OMEnableConfigure = BinSteps{
		{
			Step:    1,
			Operate: cabdef.OperateBinEnable,
			Door:    cabdef.DetectDoorIgnore,
			Battery: cabdef.DetectBatteryIgnore,
			Bin:     cabdef.DetectBinEnable,
		},
	}

	// OMDisableConfigure 运维启用
	OMDisableConfigure = BinSteps{
		{
			Step:    1,
			Operate: cabdef.OperateBinDisable,
			Door:    cabdef.DetectDoorIgnore,
			Battery: cabdef.DetectBatteryIgnore,
			Bin:     cabdef.DetectBinDisable,
		},
	}

	OMOperates = map[cabdef.Operate]BinSteps{
		cabdef.OperateDoorOpen:   OMOpenConfigure,
		cabdef.OperateBinEnable:  OMEnableConfigure,
		cabdef.OperateBinDisable: OMDisableConfigure,
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
	Step    int                  `json:"step"`    // 步骤序号
	Operate cabdef.Operate       `json:"operate"` // 操作指令
	Door    cabdef.DetectDoor    `json:"door"`    // 仓门检测
	Battery cabdef.DetectBattery `json:"battery"` // 电池检测
	Bin     cabdef.DetectBin     `json:"bin"`     // 仓位检测
}

func (b *BinStep) String() string {
	return fmt.Sprintf("第%d步, 指令: %s, 仓门: %s检测, 电池: %s检测", b.Step, b.Operate.Text(), b.Door.Text(), b.Battery.Text())
}

// Bin 仓位操控
type Bin struct {
	index int

	MainOperate  cabdef.Operate   // 主要操作
	Timeout      int64            // 超时时间
	Serial       string           // 电柜编号
	UUID         uuid.UUID        // 任务ID
	Ordinal      int              // 仓位序号
	Business     adapter.Business // 业务类别
	Steps        BinSteps         // 规划步骤
	Battery      string           // 校验放入的电池编号
	StepCallback StepCallback     // 每一步的回调
	Remark       string           // 操作备注
	BinRemark    *string          // 仓位备注
	Scan         *ent.Scan        // 扫码信息
}

// Current 获取当前步骤
func (o *Bin) Current() *BinStep {
	last := len(o.Steps) - 1
	if o.index >= last {
		return o.Steps[last]
	}
	return o.Steps[o.index]
}

// Next 标记当前步骤完成, 开始下一个步骤
func (o *Bin) Next() bool {
	if o.index+1 >= len(o.Steps) {
		return false
	}
	o.index += 1
	return true
}
