// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-01
// Based on cabservd by liasica, magicrolan@qq.com.

package g

import (
	"log"
	"os"

	"github.com/auroraride/adapter"
	"github.com/auroraride/adapter/maintain"
)

type config struct {
	adapter.Configure `mapstructure:",squash"`

	Brand        adapter.CabinetBrand
	Maintain     maintain.Config
	Debug        bool
	DeadDuration int  `koanf:"dead-duration"` // 离线判定时间
	NonBms       bool `koanf:"non-bms"`       // 是否不包含bms通讯
	Postgres     struct {
		Dsn   string
		Debug bool
	}
	Tcp struct {
		Bind string
	}
	Rpc struct {
		Bind string
	}
}

var (
	Config *config
)

func LoadConfig(defaultConfig []byte) {
	var err error
	configFile := os.Getenv("CONFIG_FILE")
	if configFile == "" {
		configFile = "config/config.yaml"
	}

	Config = new(config)
	err = adapter.LoadConfigure(Config, configFile, defaultConfig)
	if err != nil {
		log.Fatal(err)
	}

	// 验证配置
	if Config.Brand == "" {
		log.Fatal("请配置brand参数")
	}

	Config.setKeys()
}

func (c *config) setKeys() {
	CacheCabinetKey = c.GetCacheKey("CABINET")
}
