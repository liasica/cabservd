// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on cabservd by liasica, magicrolan@qq.com.

package hook

import "sync"

var (
    Postgres = NewPostgresHook()
    worker   sync.WaitGroup
)

func Start() {
    worker.Add(1)

    go Postgres.Start()

    worker.Wait()
}
