// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-15
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

import (
	"github.com/auroraride/adapter"
)

var (
	appID     string
	appSecret []byte
	baseURL   string
	version   string
)

func New() {
	conf := adapter.GetKoanf()

	appID = conf.Get("xiliulou.appId").(string)
	appSecret = []byte(conf.Get("xiliulou.appSecret").(string))
	baseURL = conf.Get("xiliulou.server").(string)
	version = conf.Get("xiliulou.version").(string)
}
