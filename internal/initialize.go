// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-06
// Based on cabservd by liasica, magicrolan@qq.com.

package internal

import (
    "context"
    "fmt"
    "github.com/auroraride/adapter/pkg/logger"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/cabinet"
    "github.com/auroraride/cabservd/internal/ent/console"
    "github.com/auroraride/cabservd/internal/g"
    "os"
    "time"
)

func initialize() {
    ctx := context.Background()

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

    // 标记所有电柜为离线和空闲
    _ = ent.Database.Cabinet.Update().Where(cabinet.Brand(g.Config.Brand)).SetOnline(false).SetStatus(cabinet.StatusIdle).Exec(ctx)

    // 标记所有正在进行的任务为失败
    _, _ = ent.Database.Console.ExecContext(ctx, fmt.Sprintf(
        `UPDATE %s SET %s = '%s', MESSAGE = '重启终止', %s = NOW(), %s = EXTRACT(EPOCH FROM (NOW() - start_at)) WHERE %s = %s AND %s = %s`,
        console.Table,
        console.FieldStatus,
        console.StatusFailed,
        console.FieldStopAt,
        console.FieldDuration,
        console.FieldBrand,
        g.Config.Brand,
        console.FieldStatus,
        console.StatusRunning,
    ))
}
