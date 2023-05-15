// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-15
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

import (
	"github.com/auroraride/adapter"
	"github.com/evanphx/wildcat"

	"github.com/auroraride/cabservd/internal/codec"
	"github.com/auroraride/cabservd/internal/core"
)

func New() (h core.Hook, c codec.Codec) {
	conf := adapter.GetKoanf()
	c = &signer{
		appID:     conf.Get("xiliulou.appId").(string),
		appSecret: conf.Get("xiliulou.appSecret").(string),
		parser:    wildcat.NewHTTPParser(),
	}
	h = &handler{
		server:  conf.Get("xiliulou.server").(string),
		version: conf.Get("xiliulou.version").(string),
	}
	return
}
