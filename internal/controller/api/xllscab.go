// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-22
// Based on cabservd by liasica, magicrolan@qq.com.

package api

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/auroraride/adapter/app"
	"github.com/labstack/echo/v4"

	"github.com/auroraride/cabservd/internal/brands/xlls"
)

type xllscab struct{}

var Xllscab = new(xllscab)

func (*xllscab) Battery(c echo.Context) (err error) {
	ctx := app.Context(c)
	var f *os.File
	f, err = os.OpenFile("runtime/xllscab-battery.list", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return
	}

	var b []byte
	b, err = io.ReadAll(f)
	if err != nil {
		return
	}
	_ = f.Close()

	s := string(b)

	var body []byte
	var notice map[string]string

	if c.Request().Method == "POST" {
		go func() {
			f, _ = os.OpenFile("runtime/xllscab-battery.list", os.O_WRONLY|os.O_APPEND, 0644)
			body, err = io.ReadAll(c.Request().Body)
			if err != nil {
				return
			}

			posts := strings.Split(string(body), "\n")

			req := &xlls.BatteryCreateRequest{}
			var reqs []*xlls.BatteryCreateRequest
			for _, sn := range posts {
				if len(req.BatteryList) == 50 {
					reqs = append(reqs, req)
					req = &xlls.BatteryCreateRequest{}
				}
				req.BatteryList = append(req.BatteryList, xlls.BatterySnData{BatterySn: sn})
			}
			if len(req.BatteryList) < 50 {
				reqs = append(reqs, req)
			}

			var sb strings.Builder
			sb.WriteRune('\n')

			for _, data := range reqs {
				_, err = xlls.FetchBatteryCreate(data)
				if err == nil {
					sb.WriteRune('\n')

					for _, sd := range data.BatteryList {
						sb.WriteString(sd.BatterySn)
						sb.WriteRune('\n')
					}
					_, _ = f.WriteString(sb.String())
				}
			}
		}()
	}

	return ctx.Render(http.StatusOK, "xllscab/battery.go.html", map[string]any{
		"items":  s,
		"notice": notice,
	})
}
