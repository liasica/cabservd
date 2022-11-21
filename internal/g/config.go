// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-01
// Based on cabservd by liasica, magicrolan@qq.com.

package g

type config struct {
    Database struct {
        Postgres struct {
            Dsn   string
            Debug bool
        }
    }
}

func LoadConfig() {
    Config = &config{}
}
