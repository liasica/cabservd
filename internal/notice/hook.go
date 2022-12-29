// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on cabservd by liasica, magicrolan@qq.com.

package notice

import "sync"

var (
    Postgres *PostgresHook
    Aurservd *aurservdHook
    worker   sync.WaitGroup
)

func Start() {
    Postgres = NewPostgres()
    Aurservd = NewAurservd()

    go Postgres.Start()
    go Aurservd.Start()

    worker.Wait()
}
