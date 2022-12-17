// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-01
// Based on cabservd by liasica, magicrolan@qq.com.

package core

import (
    "github.com/auroraride/cabservd/internal/errs"
    "github.com/auroraride/cabservd/internal/g"
    "github.com/auroraride/cabservd/internal/types"
)

func (h *hub) Control(req *types.ControlRequest) (err error) {
    if req.Serial == "" {
        return errs.CabinetSerialRequired
    }

    if req.Type == "" {
        return errs.CabinetControlParamError
    }

    if req.Ordinal == nil {
        return errs.CabinetBinOrdinalRequired
    }

    switch g.Config.Brand {
    case types.BrandKaixin:
        err = h.Bean.SendControl(req.Serial, req.Type, *req.Ordinal)
    }

    return
}
