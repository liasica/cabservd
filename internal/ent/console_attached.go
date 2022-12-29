// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-29
// Based on cabservd by liasica, magicrolan@qq.com.

package ent

import (
    "github.com/auroraride/adapter/model"
    "github.com/auroraride/cabservd/internal/ent/console"
)

func (c *Console) StepResult() *model.ExchangeStepResult {
    return &model.ExchangeStepResult{
        StartAt: c.StartAt,
        StopAt:  c.StopAt,
        Success: c.Status == console.StatusSuccess,
        Step:    c.Step,
        Before:  c.BeforeBin,
        After:   c.AfterBin,
    }
}
