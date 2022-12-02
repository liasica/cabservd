// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-11
// Based on cabservd by liasica, magicrolan@qq.com.

package kaixin

import (
    "fmt"
    "github.com/auroraride/cabservd/internal/core"
    "github.com/auroraride/cabservd/internal/errs"
    "github.com/auroraride/cabservd/internal/types"
    "time"
)

type ControlValue string

const (
    ControlCabinetDisable ControlValue = "00" // 设置换电柜不可用
    ControlExchange       ControlValue = "01" // 换电
    ControlPutIn          ControlValue = "02" // 放电
    ControlPutOut         ControlValue = "03" // 取电
    ControlOpenDoor       ControlValue = "04" // 开启柜门
    ControlBinDisable     ControlValue = "06" // 设置柜门不可用
    ControlBinEnable      ControlValue = "07" // 设置柜门可用
    ControlBatteryBind    ControlValue = "08" // 柜门绑定电池序列号
    ControlBatteryUnbind  ControlValue = "09" // 柜门解绑电池序列号
    ControlCabinetEnable  ControlValue = "10" // 设置换电柜可用
    ControlBatteryRent    ControlValue = "11" // 租用电池(首放)
    ControlBatteryTenancy ControlValue = "12" // 退还电池
)

var (
    controlValueMap = map[types.ControlType]ControlValue{
        types.ControlTypeOpenDoor: ControlOpenDoor,
        types.BinDisable:          ControlBinDisable,
        types.BinEnable:           ControlBinEnable,
    }
)

func (h *Hander) OnControl(serial string, typ types.ControlType, index int) error {
    v, ok := controlValueMap[typ]
    if !ok {
        return errs.CabinetControlParamError
    }

    req := ControlRequest{
        ParamList: []ControlParam{{
            SignalData: SignalData{
                ID:    SignalCabinetControl,
                Value: v,
            },
            DoorID: fmt.Sprintf("%d", index+1),
        }},
    }
    return SendControl(serial, req)
}

func SendControl(serial string, req ControlRequest) (err error) {
    msg := &Request{
        Message: Message{
            MsgType: MessageTypeControlRequest,
            TxnNo:   time.Now().UnixMilli(),
            DevID:   serial,
        },
        ControlRequest: req,
    }

    var c *core.Client
    c, err = core.GetClient(serial)
    if err != nil {
        return
    }

    return c.SendMessage(msg, true)
}
