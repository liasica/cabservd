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

type handler struct {
	server  string
	version string
}

func (h *handler) Protocol() core.Protocol {
	return core.ProtocolHttp
}

func (h *handler) OnConnect(c *core.Client) {

}

func (h *handler) OnMessage(_ *core.Client, b []byte) (serial string, _ core.ResponseMessenger, fields []zap.Field, err error) {
	return
}

func (h *handler) SendOperate(_ string, _ cabdef.Operate, _ int, _ int) (err error) {
	return
}

func (h *handler) GetEmptyDeviation() (voltage, current float64) {
	return
}
