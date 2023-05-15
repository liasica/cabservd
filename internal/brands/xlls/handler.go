// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-15
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

import (
	"github.com/auroraride/adapter/defs/cabdef"
	"go.uber.org/zap"

	"github.com/auroraride/cabservd/internal/core"
)

type xlls struct {
}

func (x *xlls) Protocol() core.Protocol {
	return core.ProtocolHttp
}

func (x *xlls) OnConnect(c *core.Client) {

}

func (x *xlls) OnMessage(_ *core.Client, b []byte) (serial string, _ core.ResponseMessenger, fields []zap.Field, err error) {
	return
}

func (x *xlls) SendOperate(_ string, _ cabdef.Operate, _ int, _ int) (err error) {
	return
}

func (x *xlls) GetEmptyDeviation() (voltage, current float64) {
	return
}
