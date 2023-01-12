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
    "github.com/auroraride/adapter/pkg/loki"
    "github.com/auroraride/cabservd/internal/ent/migrate"
    _ "github.com/auroraride/cabservd/internal/ent/runtime"
    _ "github.com/jackc/pgx/v4/stdlib"
)

var Database *Client

func OpenDatabase(dsn string, debug bool) *Client {
    pgx, err := sql.Open("pgx", dsn)
    if err != nil {
        loki.Fatalf("数据库打开失败: %v", err)
    }

    // 从db变量中构造一个ent.Driver对象。
    drv := entsql.OpenDB(dialect.Postgres, pgx)
    c := NewClient(Driver(drv))
    if debug {
        c = c.Debug()
    }

    autoMigrate(c)

    createFunction(c)

    return c
}

func createFunction(c *Client) {
    raw := `CREATE OR REPLACE FUNCTION notify_event() RETURNS TRIGGER AS
$$

DECLARE
    data         JSON;
    notification JSON;

BEGIN
    -- Convert the old or new row to JSON, based on the kind of action.
    -- Action = DELETE?             -> OLD row
    -- Action = INSERT or UPDATE?   -> NEW row
    IF (TG_OP = 'DELETE') THEN
        data = ROW_TO_JSON(OLD);
    ELSE
        data = ROW_TO_JSON(NEW);
    END IF;

    -- Contruct the notification as a JSON string.
    notification = JSON_BUILD_OBJECT(
            'table', TG_TABLE_NAME,
            'action', TG_OP,
            'data', data);

    -- Execute pg_notify(channel, notification)
    PERFORM pg_notify(TG_TABLE_NAME, notification::TEXT);

    -- Result is ignored since this is an AFTER trigger
    RETURN NULL;
END;

$$ LANGUAGE plpgsql;

DO
$$
    BEGIN
        IF NOT EXISTS(SELECT 1 FROM pg_trigger WHERE tgname = 'cabinet_notify_event') THEN
            CREATE TRIGGER cabinet_notify_event
                AFTER INSERT OR UPDATE OR DELETE
                ON cabinet
                FOR EACH ROW
            EXECUTE PROCEDURE notify_event();
        END IF;

        IF NOT EXISTS(SELECT 1 FROM pg_trigger WHERE tgname = 'bin_notify_event') THEN
            CREATE TRIGGER bin_notify_event
                AFTER INSERT OR UPDATE OR DELETE
                ON bin
                FOR EACH ROW
            EXECUTE PROCEDURE notify_event();
        END IF;
    END
$$;`
    _, err := c.ExecContext(context.Background(), raw)
    if err != nil {
        loki.Fatal(err)
    }
}

func autoMigrate(c *Client) {
    ctx := context.Background()

    if err := c.Schema.Create(
        ctx,
        migrate.WithDropIndex(true),
        migrate.WithDropColumn(true),
        migrate.WithForeignKeys(false),
    ); err != nil {
        loki.Fatalf("数据库迁移失败: %v", err)
    }
}
