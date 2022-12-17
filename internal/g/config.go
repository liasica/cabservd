// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-01
// Based on cabservd by liasica, magicrolan@qq.com.

package g

import (
    "github.com/auroraride/cabservd/assets"
    log "github.com/sirupsen/logrus"
    "github.com/spf13/viper"
    "os"
    "path/filepath"
)

type config struct {
    Brand    string
    Postgres struct {
        Dsn   string
        Debug bool
    }
    Api struct {
        Bind string
    }
    Tcp struct {
        Bind string
    }
}

var (
    Config *config
)

func LoadConfig() {
    dir := "config"
    if _, err := os.Stat(dir); os.IsNotExist(err) {
        err = os.MkdirAll(dir, 0755)
        if err != nil {
            log.Fatal(err)
        }
    }

    cf := filepath.Join(dir, "config.yaml")
    if _, err := os.Stat(cf); os.IsNotExist(err) {
        _ = os.WriteFile(cf, assets.DefaultConfig, 0755)
    }

    viper.SetConfigFile(cf)
    viper.AutomaticEnv()
    // 读取配置
    err := viper.ReadInConfig()
    if err != nil {
        log.Fatal(err)
    }

    Config = &config{}

    // 解析配置
    err = viper.Unmarshal(Config)
    if err != nil {
        log.Fatal(err)
    }

    // 验证配置
    if Config.Brand == "" {
        log.Fatal("请配置brand参数")
    }
}
