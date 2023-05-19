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

var binCommand = map[cabdef.Operate]CellCommand{
	cabdef.OperateDoorOpen:   CellOpenDoor,
	cabdef.OperateBinDisable: CellForbid,
	cabdef.OperateBinEnable:  CellUnForbid,
}

func BinTransfer(serial string, ordinal int, business adapter.Business, operate cabdef.Operate, notifier chan *ent.Bin, times int) (err error) {
	switch business {
	case adapter.BusinessOperate:
		// 运维操作
		_, err = FetchCellCommand(&CellCommandRequest{
			Sn:      serial,
			CellNos: []int{ordinal},
			Command: binCommand[operate],
		})
		return
	}
	return
}
