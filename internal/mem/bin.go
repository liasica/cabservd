// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on cabservd by liasica, magicrolan@qq.com.

package mem

import (
    "github.com/auroraride/adapter/defs/cabdef"
    "strconv"
    "sync"
)

var (
    // 操作中的仓位 string => cabdef.Operate
    binOperating sync.Map // operation
)

func binKey(serial string, ordinal int) string {
    return serial + "-" + strconv.Itoa(ordinal)
}

// BinInOperation 获取当前进行中的操作
func BinInOperation(serial string, ordinal int) cabdef.Operate {
    v, ok := binOperating.Load(binKey(serial, ordinal))
    if ok {
        return v.(cabdef.Operate)
    }
    return ""
}

func BinOperationFinished(serial string, ordinal int) {
    binOperating.Delete(binKey(serial, ordinal))
}

func BinOperate(serial string, ordinal int, o cabdef.Operate) {
    binOperating.Store(binKey(serial, ordinal), o)
}
