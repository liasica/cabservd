// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-01
// Based on cabservd by liasica, magicrolan@qq.com.

package internal

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/auroraride/adapter/app"
	"github.com/auroraride/adapter/log"
	"github.com/auroraride/adapter/maintain"
	"github.com/go-redis/redis/v9"
	"github.com/labstack/echo/v4"

	"github.com/auroraride/cabservd/assets"
	"github.com/auroraride/cabservd/internal/demo"
	"github.com/auroraride/cabservd/internal/ent"
	"github.com/auroraride/cabservd/internal/ent/cabinet"
	"github.com/auroraride/cabservd/internal/ent/console"
	"github.com/auroraride/cabservd/internal/g"
	"github.com/auroraride/cabservd/internal/router"
	"github.com/auroraride/cabservd/internal/rpc"
	"github.com/auroraride/cabservd/internal/sync"
)

func Boot(starter func()) {
	ctx := context.Background()

	// 设置全局时区
	tz := "Asia/Shanghai"
	_ = os.Setenv("TZ", tz)
	loc, _ := time.LoadLocation(tz)
	time.Local = loc

	// 加载配置
	g.LoadConfig(assets.DefaultConfig)

	// 加载redis
	g.Redis = redis.NewClient(&redis.Options{
		Addr: g.Config.Redis.Address,
		DB:   0,
	})

	// 初始化日志
	logcfg := &log.Config{
		FormatJson: !g.Config.LoggerDebug,
		Stdout:     g.Config.LoggerDebug,
		LoggerName: g.Config.LoggerName,
		NoCaller:   true,
	}
	if !g.Config.LoggerDebug {
		logcfg.Writers = append(logcfg.Writers, log.NewRedisWriter(g.Redis))
	}
	log.New(logcfg)

	// 加载模板
	assets.LoadTemplates()

	// 加载数据库
	ent.Database = ent.OpenDatabase(g.Config.Postgres.Dsn, g.Config.Postgres.Debug)

	// 标记所有电柜为正常和空闲
	_ = ent.Database.Cabinet.Update().SetOnline(false).SetStatus(cabinet.StatusNormal).Exec(ctx)

	// 标记所有正在进行的任务为失败
	_, _ = ent.Database.Console.ExecContext(ctx, fmt.Sprintf(
		`UPDATE %s SET %s = '%s', MESSAGE = '异常重启终止', %s = NOW(), %s = EXTRACT(EPOCH FROM (NOW() - start_at)) WHERE %s = '%s'`,
		console.Table,
		console.FieldStatus,
		console.StatusFailed,
		console.FieldStopAt,
		console.FieldDuration,
		console.FieldStatus,
		console.StatusRunning,
	))

	// 加载hooks
	sync.Start()

	// 启动rpc server
	go rpc.Run()

	// 启动 http server
	userSkipper := map[string]bool{
		"/maintain/update/:token": true,
		"/maintain/clients":       true,
		"/xllscab/battery":        true,
	}
	e := app.NewEcho(&app.EchoConfig{
		AuthSkipper: func(c echo.Context) bool {
			return userSkipper[c.Path()]
		},
		Maintain: g.Config.Maintain,
		DumpSkipper: func(c echo.Context) bool {
			return c.Path() == "/maintain/clients"
		},
	})
	go router.Start(e, g.Config.Brand)

	go starter()

	// maintain
	if maintain.Exists() {
		_ = maintain.Remove()
	}

	// debug
	go demo.Debug()

	select {
	case <-app.Quit:
		_ = e.Shutdown(context.Background())
	}
}
