// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-15
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

import (
	"errors"
	"net/http"

	"github.com/auroraride/adapter"
	"github.com/evanphx/wildcat"

	"github.com/auroraride/cabservd/internal/core"
)

type signer struct{}

func (codec *signer) Decode(c *core.Client) (b []byte, err error) {
	buf, _ := c.Peek(-1)
	parser := wildcat.NewHTTPParser()

	var offset int
	offset, err = parser.Parse(buf)
	if err != nil {
		// 消息未收完继续收取
		if err == wildcat.ErrMissingData {
			err = adapter.ErrorIncompletePacket
		}
		return
	}

	realIP := parser.FindHeader(headerXRealIP)
	if realIP == nil {
		realIP = parser.FindHeader(headerXForwardedFor)
	}
	if realIP != nil {
		c.SetIP(adapter.ConvertBytes2String(realIP))
	}

	// 是否POST
	method := adapter.ConvertBytes2String(parser.Method)
	if method != allowMethod {
		_, _ = c.Write(httpResponseRaw(http.StatusMethodNotAllowed, nil))
		_ = c.Close()
		err = errors.New("请求方式未被允许: " + method)
		return
	}

	// 获取消息体长度
	bodyLen := int(parser.ContentLength())
	// 未获取到消息体长度, 返回继续缓存消息
	if bodyLen == -1 {
		return nil, adapter.ErrorIncompletePacket
	}

	// 获取本次消息体长度
	msgLen := bodyLen + offset

	// 若已缓存消息长度小于需要长度, 返回错误: 消息未接收完成, 继续缓存消息
	if c.InboundBuffered() < msgLen {
		return nil, adapter.ErrorIncompletePacket
	}

	buf, _ = c.Peek(msgLen)

	// 消息体前4位存放path信息
	b = append(parser.Path, buf[offset:msgLen]...)

	_, _ = c.Discard(msgLen)

	return
}

func (codec *signer) Encode(data []byte) (b []byte) {
	return httpResponseRaw(http.StatusOK, data)
}
