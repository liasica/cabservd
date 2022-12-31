// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-01
// Based on cabservd by liasica, magicrolan@qq.com.

package core

import (
    "github.com/auroraride/adapter"
    "github.com/auroraride/cabservd/internal/g"
)

func (h *hub) Control(req *adapter.OperateRequest) (err error) {
    if req.Serial == "" {
        return adapter.CabinetSerialRequired
    }

    if req.Type == "" {
        return adapter.CabinetControlParamError
    }

    if req.Ordinal == nil {
        return adapter.CabinetBinOrdinalRequired
    }

    switch g.Config.Brand {
    case adapter.BrandKaixin:
        err = h.Bean.SendControl(req.Serial, req.Type, *req.Ordinal)
    }

    return
}
