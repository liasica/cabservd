// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-15
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

import (
	"github.com/auroraride/adapter"
	"github.com/evanphx/wildcat"

	"github.com/auroraride/cabservd/internal/core"
)

var (
	appID     string
	appSecret []byte
	baseURL   string
	version   string
)

func New() (h core.Hook, c core.Codec) {
	conf := adapter.GetKoanf()

	appID = conf.Get("xiliulou.appId").(string)
	appSecret = []byte(conf.Get("xiliulou.appSecret").(string))
	baseURL = conf.Get("xiliulou.server").(string)
	version = conf.Get("xiliulou.version").(string)

	c = &signer{
		parser: wildcat.NewHTTPParser(),
	}

	h = &xlls{}
	return
}