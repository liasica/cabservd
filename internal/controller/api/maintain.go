// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-06
// Based on cabservd by liasica, magicrolan@qq.com.

package api

import (
	"net/http"
	"sort"
	"strings"

	"github.com/auroraride/cabservd/internal/core"
	"github.com/labstack/echo/v4"
)

type maintainApi struct{}

var Maintain = new(maintainApi)

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
