// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-06
// Based on cabservd by liasica, magicrolan@qq.com.

package api

import (
    "github.com/auroraride/adapter/async"
    "github.com/auroraride/adapter/maintain"
    "github.com/auroraride/adapter/zlog"
    "github.com/auroraride/cabservd/internal/core"
    "github.com/auroraride/cabservd/internal/g"
    "github.com/labstack/echo/v4"
    "net/http"
    "sort"
    "strings"
    "time"
)

type maintainApi struct{}

var Maintain = new(maintainApi)

func (*maintainApi) Update(echo.Context) (err error) {
    ticker := time.NewTicker(time.Second)
    defer ticker.Stop()

    _ = maintain.Create()

    for ; true; <-ticker.C {
        // 是否有进行中的异步业务
        if async.IsDone() {
            g.Quit <- true
            zlog.Sync()
            return
        }
    }
    return nil
}

func (*maintainApi) Clients(c echo.Context) (err error) {
    var clients []*core.Client
    core.Hub.Clients.Range(func(k, v any) bool {
        clients = append(clients, v.(*core.Client))
        return true
    })

    sort.Slice(clients, func(i, j int) bool {
        return strings.Compare(clients[i].Serial, clients[j].Serial) < 0
    })

    return c.Render(http.StatusOK, "maintain/clients.go.html", clients)
}
