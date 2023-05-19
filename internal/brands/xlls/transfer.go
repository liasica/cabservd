// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-19
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

import (
	"github.com/auroraride/adapter"
	"github.com/auroraride/adapter/defs/cabdef"

	"github.com/auroraride/cabservd/internal/ent"
)

func BinTransfer(serial string, ordinal int, business adapter.Business, typ cabdef.Operate, notifier chan *ent.Bin, times int) (err error) {
	return
}
