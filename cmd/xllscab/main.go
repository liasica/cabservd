// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-09
// Based on cabservd by liasica, magicrolan@qq.com.

package main

import (
	"github.com/auroraride/cabservd/internal"
	"github.com/auroraride/cabservd/internal/brands/xlls"
	"github.com/auroraride/cabservd/internal/g"
)

func main() {
	g.UseHttp = true

	internal.Boot(func() {
		xlls.Start()
	})
}
