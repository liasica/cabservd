// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on cabservd by liasica, magicrolan@qq.com.

package hook

var (
    Postgres = NewPostgresHook()
)

func Start() {
    go Postgres.Start()
}
