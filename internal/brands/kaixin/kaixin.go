// Copyright (C) liasica. 2023-present.
//
// Created at 2023-03-09
// Based on cabservd by liasica, magicrolan@qq.com.

package kaixin

import (
    "github.com/auroraride/cabservd/internal/codec"
    "github.com/auroraride/cabservd/internal/core"
    "github.com/auroraride/cabservd/internal/tower"
)

func New() (core.Hook, codec.Codec) {
    return tower.New(
            &tower.MessageTypeList{
                LoginRequest:    100,
                LoginResponse:   101,
                ReportRequest:   300,
                ReportResponse:  301,
                NoticeRequest:   400,
                NoticeResponse:  401,
                ControlRequest:  500,
                ControlResponse: 501,
            },
            binSignals,
        ),
        &codec.HeaderLength{}
}
