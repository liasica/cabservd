// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-01
// Based on cabservd by liasica, magicrolan@qq.com.

package internal

import (
    "github.com/auroraride/adapter/logger"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/g"
    "os"
    "time"
)

func Boot() {
    // 设置全局时区
    tz := "Asia/Shanghai"
    _ = os.Setenv("TZ", tz)
    loc, _ := time.LoadLocation(tz)
    time.Local = loc

    // 日志
    logger.LoadWithConfig(logger.Config{
        Color:  true,
        Level:  "info",
        Age:    8192,
        Caller: true,
    })

    // 加载配置
    g.LoadConfig()

    // 加载数据库
    ent.Database = ent.OpenDatabase(g.Config.Postgres.Dsn, g.Config.Postgres.Debug)
}
