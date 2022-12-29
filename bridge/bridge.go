// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-26
// Based on cabservd by liasica, magicrolan@qq.com.

package bridge

import "sync"

var (
    worker sync.WaitGroup
)

func Start() {
    worker.Add(1)

    go startAurservd()

    worker.Wait()
}
