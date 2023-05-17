// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-15
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

import (
	"fmt"

	"github.com/auroraride/adapter"
	"github.com/auroraride/adapter/defs/cabdef"
	"github.com/auroraride/adapter/log"
	jsoniter "github.com/json-iterator/go"
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

func (x *xlls) OnMessage(_ *core.Client, b []byte) (serial string, res core.ResponseMessenger, fields []zap.Field, err error) {
	var payload Notifyer

	defer func() {
		if v := recover(); v != nil {
			err = fmt.Errorf("%v", v)
		}
		if payload != nil {
			res = &NotifyResult[any]{
				Code:      0,
				RequestID: payload.GetRequestID(),
				Data:      nil,
			}
		}
	}()

	// 获取path
	path := adapter.ConvertBytes2String(b[:4])

	// {"cellNo":8,"doorStatus":1,"indicatorLightStatus":0}
	// {"voltage":234.4,"current":0.9,"fanStatus":0,"cVersion":"1.18"}
	// /cab {"voltage":228.6,"current":0.87,"temp":28,"fanStatus":0,"lightStatus":0,"power":198.8,"powerFactor":0.15,"activeElectricityEnergy":12.569999694824219,"waterPumpStatus":0,"waterLeachingWarning":0,"humidity":38,"doorStatus":1,"cVersion":"1.18"}

	switch path {
	case pathHardwareOperation:
	case pathBusinesss:
	case pathOfflineExchange:
	case pathCellChange:
		payload = new(CellNotify)
	case pathBatteryChange:
	case pathCabinetChange:
		payload = new(CabinetNotify)
	case pathHardwareFault:
	case pathSelfServiceOpen:
	}

	// 获取数据并解析
	err = jsoniter.Unmarshal(b[4:], payload)

	fields = []zap.Field{
		zap.String("path", path),
		log.Payload(payload),
	}

	return
}

func (x *xlls) SendOperate(_ string, _ cabdef.Operate, _ int, _ int) (err error) {
	return
}

func (x *xlls) GetEmptyDeviation() (voltage, current float64) {
	return
}
