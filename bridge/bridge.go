// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-26
// Based on cabservd by liasica, magicrolan@qq.com.

package bridge

import "github.com/auroraride/cabservd/internal/ent"

type Bridger struct {
    cabinet *cabinetBridge
}

var (
    hub *Bridger
)

func Start() {
    hub = &Bridger{
        cabinet: newCabinet(),
    }

    go hub.cabinet.run()
}

func SendCabinet(serial string, cab *ent.Cabinet, bins ent.Bins) {
    hub.cabinet.bridger.SendSyncData(hub.cabinet.WrapData(serial, cab, bins))
}
