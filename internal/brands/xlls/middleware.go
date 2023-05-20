// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-18
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

import (
	"bufio"
	"bytes"
	"io"
	"net"
	"net/http"

	"github.com/auroraride/adapter/log"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type DumpResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w *DumpResponseWriter) WriteHeader(code int) {
	w.ResponseWriter.WriteHeader(code)
}

func (w *DumpResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func (w *DumpResponseWriter) Flush() {
	w.ResponseWriter.(http.Flusher).Flush()
}

func (w *DumpResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return w.ResponseWriter.(http.Hijacker).Hijack()
}

func dump() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			var reqBody []byte
			if c.Request().Body != nil {
				reqBody, _ = io.ReadAll(c.Request().Body)
			}
			c.Request().Body = io.NopCloser(bytes.NewBuffer(reqBody)) // Reset body

			// Response
			resBody := new(bytes.Buffer)
			mw := io.MultiWriter(c.Response().Writer, resBody)
			writer := &DumpResponseWriter{Writer: mw, ResponseWriter: c.Response().Writer}
			c.Response().Writer = writer

			err = next(c)

			// 记录
			zap.L().Info(
				"收到消息↑↓",
				zap.String("path", c.Path()),
				log.Payload(c.Get("payload")),
				zap.ByteString("body", reqBody),
				log.ResponseBody(bytes.TrimRight(resBody.Bytes(), "\n")),
			)

			return
		}
	}
}
