// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-19
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

import (
	"github.com/auroraride/adapter"
	"github.com/auroraride/adapter/defs/cabdef"

	"github.com/auroraride/cabservd/internal/types"
)

var binCommand = map[cabdef.Operate]CellCommand{
	cabdef.OperateDoorOpen:   CellOpenDoor,
	cabdef.OperateBinDisable: CellForbid,
	cabdef.OperateBinEnable:  CellUnForbid,
}

func BinTransfer(bo *types.Bin, step *types.BinStep) (err error) {
	if step.Step != 1 {
		return
	}

	switch bo.Business {
	case adapter.BusinessOperate:
		// 运维操作
		_, err = fetchCellCommand(&CellCommandRequest{
			Sn:      bo.Serial,
			CellNos: []int{bo.Ordinal},
			Command: binCommand[step.Operate],
		})
		return
	case adapter.BusinessExchange:
		_, err = fetchExchange(&BusinessExchangeRequest{
			Sn:               bo.Serial,
			OrderNo:          *bo.Scan.OrderNo,
			EmptyCellNo:      bo.Scan.Data.Empty.Ordinal,
			BatteryCellNo:    bo.Scan.Data.Fully.Ordinal,
			BindingBatterySn: bo.Battery,
		})
	}
	return
}
