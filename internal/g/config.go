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
    Brand       adapter.CabinetBrand
    Maintain    maintain.Config
    Application string
    Debug       bool
    Postgres    struct {
        Dsn   string
        Debug bool
    }
    Tcp struct {
        Bind string
    }
    Api struct {
        Bind string
    }
    Adapter struct {
        Aurservd string
    }
    Loki struct {
        Job string
        Url string
    }
    Redis struct {
        Logkey  string
        Address string
    }
}

var (
    Config *config
)

func LoadConfig() {
    var err error

    Config, err = adapter.LoadConfigure[config](configFile, assets.DefaultConfig)
    if err != nil {
        log.Fatal(err)
    }

    // 验证配置
    if Config.Brand == "" {
        log.Fatal("请配置brand参数")
    }
}
