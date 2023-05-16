// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-16
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

import (
	"net/http"
	"strconv"
	"time"

	"github.com/auroraride/adapter"
)

type Response struct {
	Status int
	Data   any
}

var resNewline = []byte("\r\n")

func httpResponseRaw(code int, data []byte) []byte {
	buf := adapter.NewBuffer()
	defer adapter.ReleaseBuffer(buf)

	buf.WriteString("HTTP/1.1 ")
	buf.WriteString(strconv.Itoa(code))
	buf.WriteString(" ")
	buf.WriteString(http.StatusText(code))
	buf.Write(resNewline)

	buf.WriteString("Server: aurcab-xiliulou")
	buf.Write(resNewline)

	buf.WriteString("Content-Type: application/json; charset=UTF-8")
	buf.Write(resNewline)

	buf.WriteString("Connection: close")
	buf.Write(resNewline)

	buf.WriteString("Date: ")
	buf.WriteString(time.Now().UTC().Format(http.TimeFormat))
	buf.Write(resNewline)

	buf.WriteString("Content-Length: ")
	buf.WriteString(strconv.Itoa(len(data)))
	buf.Write(resNewline)
	buf.Write(resNewline)

	buf.Write(data)

	return buf.Bytes()
}
