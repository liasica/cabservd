// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-29
// Based on cabservd by liasica, magicrolan@qq.com.

package ent

import (
	"github.com/auroraride/adapter/defs/cabdef"

	"github.com/auroraride/cabservd/internal/ent/console"
)

func (c *Console) OperateResult() (res *cabdef.BinOperateResult) {
	res = &cabdef.BinOperateResult{
		UUID:     c.UUID.String(),
		Step:     c.Step,
		Business: c.Business,
		StartAt:  c.StartAt,
		StopAt:   c.StopAt,
		Success:  c.Status == console.StatusSuccess,
		Before:   c.BeforeBin,
		After:    c.AfterBin,
	}

	if c.Duration != nil {
		res.Duration = *c.Duration
	}

	if c.Message != nil {
		res.Message = *c.Message
	}

	return
}
