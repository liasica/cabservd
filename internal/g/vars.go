// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-06
// Based on cabservd by liasica, magicrolan@qq.com.

package g

import "github.com/redis/go-redis/v9"

var (
	Redis *redis.Client

	CacheCabinetKey string
)
