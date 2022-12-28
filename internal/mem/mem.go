// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on cabservd by liasica, magicrolan@qq.com.

package mem

import (
    "fmt"
    "github.com/auroraride/cabservd/internal/ent"
    "sync"
)

var (
    // Cabinets 电柜缓存 id => *ent.Cabinet
    Cabinets sync.Map

    // Bins 仓位缓存 serial+id => *ent.Bin
    Bins sync.Map
)

func SetCabinet(cab *ent.Cabinet) {
    Cabinets.Store(cab.ID, cab)
}

func GetCabinet(id uint64) *ent.Cabinet {
    cab, ok := Cabinets.Load(id)
    if ok {
        return cab.(*ent.Cabinet)
    }
    return nil
}

func SetBin(b *ent.Bin) {
    Cabinets.Store(fmt.Sprintf("%s-%d", b.Serial, b.ID), b)
}

func GetBin(serial string, id uint64) *ent.Bin {
    b, ok := Bins.Load(fmt.Sprintf("%s-%d", serial, id))
    if ok {
        return b.(*ent.Bin)
    }
    return nil
}
