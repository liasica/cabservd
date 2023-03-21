// Copyright (C) liasica. 2023-present.
//
// Created at 2023-03-09
// Based on cabservd by liasica, magicrolan@qq.com.

package main

import (
    "github.com/auroraride/cabservd/internal"
    "github.com/auroraride/cabservd/internal/codec"
    "github.com/auroraride/cabservd/internal/core"
    "github.com/auroraride/cabservd/internal/tower"
)

func main() {
    internal.Boot(
        func() (core.Hook, codec.Codec) {
            return tower.New(
                    tower.WithMessageTypeList(&tower.MessageTypeList{
                        LoginRequest:    110,
                        LoginResponse:   111,
                        ReportRequest:   310,
                        ReportResponse:  311,
                        NoticeRequest:   410,
                        NoticeResponse:  411,
                        ControlRequest:  500,
                        ControlResponse: 501,
                    }),
                    tower.WithAutoResetBattery(true),
                    tower.WithCalculateMonVoltage(true),
                ),
                &codec.Linebreak{}
        },
    )
}
