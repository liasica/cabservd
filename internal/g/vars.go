// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-06
// Based on cabservd by liasica, magicrolan@qq.com.

package g

import "github.com/go-redis/redis/v9"

var (
    Quit  chan bool
    Redis *redis.Client
)

func init() {
    Quit = make(chan bool)
}
