// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-04
// Based on cabservd by liasica, magicrolan@qq.com.

package kaixin

import "cabservd/core/bean"

type Hander struct {
    bean.Bean
}

func New() *Hander {
    return &Hander{}
}
