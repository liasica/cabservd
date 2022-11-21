// Copyright (C) liasica. 2022-present.
//
// Created at 2022-06-30
// Based on aurservd by liasica, magicrolan@qq.com.

package ent

import (
    "context"
    "database/sql"
    "entgo.io/ent/dialect"
    entsql "entgo.io/ent/dialect/sql"
    "github.com/auroraride/cabservd/internal/ent/migrate"
    log "github.com/sirupsen/logrus"

    _ "github.com/auroraride/cabservd/internal/ent/runtime"
    _ "github.com/jackc/pgx/v4/stdlib"
)

var Database *Client

func OpenDatabase(dsn string, debug bool) *Client {
    pgx, err := sql.Open("pgx", dsn)
    if err != nil {
        log.Fatalf("数据库打开失败: %v", err)
    }

    // 从db变量中构造一个ent.Driver对象。
    drv := entsql.OpenDB(dialect.Postgres, pgx)
    c := NewClient(Driver(drv))
    if debug {
        c = c.Debug()
    }

    autoMigrate(c)

    return c
}

func autoMigrate(c *Client) {
    ctx := context.Background()
    if err := c.Schema.Create(
        ctx,
        migrate.WithDropIndex(true),
        migrate.WithDropColumn(true),
        migrate.WithForeignKeys(false),
    ); err != nil {
        log.Fatalf("数据库迁移失败: %v", err)
    }
}
