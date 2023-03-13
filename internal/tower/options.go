// Copyright (C) liasica. 2023-present.
//
// Created at 2023-03-10
// Based on cabservd by liasica, magicrolan@qq.com.

package tower

type Option interface {
    apply(h *Handler)
}

type optionFunc func(h *Handler)

func (f optionFunc) apply(h *Handler) {
    f(h)
}

func WithFakeVoltage(v float64) Option {
    return optionFunc(func(h *Handler) {
        h.fakeVoltage = v
    })
}

func WithFakeCurrent(v float64) Option {
    return optionFunc(func(h *Handler) {
        h.fakeCurrent = v
    })
}

func WithMessageTypeList(v *MessageTypeList) Option {
    return optionFunc(func(h *Handler) {
        h.mtl = v
    })
}

func WithCabinetSignals(v map[Signal]CabinetSignalFunc) Option {
    return optionFunc(func(_ *Handler) {
        cabinetSignals = v
        for k := range v {
            cabinetSignalMap[k] = struct{}{}
        }
    })
}

func WithBinSignals(v map[Signal]BinSignalFunc) Option {
    return optionFunc(func(_ *Handler) {
        binSignals = v
    })
}

func WithAutoResetBattery(v bool) Option {
    return optionFunc(func(h *Handler) {
        h.device.AutoResetWithoutBatterySN = v
    })
}

func WithBatteryReign(v bool) Option {
    return optionFunc(func(h *Handler) {
        h.device.BatteryReign = v
    })
}
