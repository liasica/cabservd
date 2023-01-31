// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-01
// Based on cabservd by liasica, magicrolan@qq.com.

package g

import (
    "github.com/auroraride/adapter"
    "github.com/auroraride/adapter/maintain"
    "github.com/auroraride/cabservd/assets"
    "log"
)

const (
    configFile = "config/config.yaml"
)

type config struct {
    adapter.Configure `mapstructure:",squash"`

    Brand    adapter.CabinetBrand
    Maintain maintain.Config
    Debug    bool
    Postgres struct {
        Dsn   string
        Debug bool
    }
    Tcp struct {
        Bind string
    }
    Aurservd struct {
        Api string
    }
    Bmservd map[adapter.BatteryBrand]struct {
        Api string
    }
}

var (
    Config *config
)

func LoadConfig() {
    var err error

    Config = new(config)
    err = adapter.LoadConfigure(Config, configFile, assets.DefaultConfig)
    if err != nil {
        log.Fatal(err)
    }

    // 验证配置
    if Config.Brand == "" {
        log.Fatal("请配置brand参数")
    }

    Config.setKeys()
}

func (c *config) GetBmsApiUrl(brand adapter.BatteryBrand, api string) (url string, err error) {
    cfg, ok := c.Bmservd[brand]
    if !ok || cfg.Api == "" {
        err = adapter.ErrorConfig
        return
    }
    url = cfg.Api + api
    return
}

func (c *config) setKeys() {
    CacheCabinetKey = c.GetCacheKey("CABINET")
}
