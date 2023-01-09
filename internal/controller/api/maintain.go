// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-06
// Based on cabservd by liasica, magicrolan@qq.com.

package api

import (
    "github.com/auroraride/adapter/async"
    "github.com/auroraride/adapter/maintain"
    "github.com/auroraride/cabservd/internal/core"
    "github.com/auroraride/cabservd/internal/g"
    jsoniter "github.com/json-iterator/go"
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
            return
        }
    }
    return nil
}

func (*maintainApi) Clients(c echo.Context) (err error) {
    var keys []string
    core.Hub.Clients.Range(func(key, _ any) bool {
        keys = append(keys, key.(string))
        return true
    })

    sort.Slice(keys, func(i, j int) bool {
        return strings.Compare(keys[i], keys[j]) < 0
    })

    b, _ := jsoniter.Marshal(keys)

    return c.HTMLBlob(http.StatusOK, b)
}
