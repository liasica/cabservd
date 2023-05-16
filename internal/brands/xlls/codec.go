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
	"github.com/panjf2000/gnet/v2"
)

type signer struct {
	parser *wildcat.HTTPParser
}

func (codec *signer) Decode(conn gnet.Conn) (b []byte, err error) {
	buf, _ := conn.Peek(-1)

	var offset int
	offset, err = codec.parser.Parse(buf)
	if err != nil {
		// 消息未收完继续收取
		if err == wildcat.ErrMissingData {
			err = adapter.ErrorIncompletePacket
		}
		return
	}

	// 是否POST
	method := adapter.ConvertBytes2String(codec.parser.Method)
	if method != allowMethod {
		_, _ = conn.Write(httpResponseRaw(http.StatusMethodNotAllowed, nil))
		_ = conn.Close()
		err = errors.New("请求方式未被允许: " + method)
		return
	}

	// 获取消息体长度
	bodyLen := int(codec.parser.ContentLength())
	// 未获取到消息体长度, 返回继续缓存消息
	if bodyLen == -1 {
		return nil, adapter.ErrorIncompletePacket
	}

	// 获取本次消息体长度
	n := bodyLen + offset

	// 若已缓存消息长度小于需要长度, 返回错误: 消息未接收完成, 继续缓存消息
	if conn.InboundBuffered() < n {
		return nil, adapter.ErrorIncompletePacket
	}

	// 消息体前4位存放path信息
	b = append(codec.parser.Path, buf[offset:n]...)

	_, _ = conn.Discard(n)

	return
}

func (codec *signer) Encode(data []byte) (b []byte) {
	return httpResponseRaw(http.StatusOK, data)
}
