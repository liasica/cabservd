// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-19
// Based on cabservd by liasica, magicrolan@qq.com.

package kernel

import (
	"github.com/auroraride/adapter"
	"github.com/google/uuid"

	"github.com/auroraride/cabservd/internal/ent"
	"github.com/auroraride/cabservd/internal/types"
)

type BinOperateOptions struct {
	Uid          uuid.UUID
	Business     adapter.Business
	Remark       string
	Bin          *ent.Bin
	Stepper      chan *types.BinResult
	StepCallback types.StepCallback
}
