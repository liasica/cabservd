// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-01
// Based on cabservd by liasica, magicrolan@qq.com.

package core

import (
    "github.com/auroraride/cabservd/internal/core/kaixin"
    "github.com/auroraride/cabservd/internal/errs"
    "github.com/auroraride/cabservd/types"
)

func (h *hub) Control(req *types.ControlRequest) (err error) {
    if req.Index == nil {
        return errs.CabinetBinIndexRequired
    }

    switch req.Brand {
    case types.BrandKaixin:
        err = kaixin.Control(req.Serial, req.Type, *req.Index)
    }

    return
}
