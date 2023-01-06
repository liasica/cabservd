// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-06
// Based on cabservd by liasica, magicrolan@qq.com.

package g

var (
    HttpStop chan bool
)

func init() {
    HttpStop = make(chan bool)
}
