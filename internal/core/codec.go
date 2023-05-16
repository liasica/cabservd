// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-10
// Based on cabservd by liasica, magicrolan@qq.com.

package core

import (
	"bytes"
	"encoding/binary"

	"github.com/auroraride/adapter"
)

const (
	bodySize           = 4
	linebreakDelimiter = '\n'
)

type Codec interface {
	Decode(c *Client) (b []byte, err error)
	Encode(data []byte) (b []byte)
}

type Default struct {
}

func (codec *Default) Decode(c *Client) (b []byte, err error) {
	return c.Next(-1)
}

func (codec *Default) Encode(message []byte) []byte {
	return message
}

// Linebreak 以\n为分割处理
type Linebreak struct{}

func (codec *Linebreak) Decode(c *Client) (b []byte, err error) {
	b, err = c.Peek(-1)
	if err != nil {
		return nil, err
	}

	n := bytes.IndexByte(b, linebreakDelimiter)
	if n < 0 {
		return nil, adapter.ErrorIncompletePacket
	}

	_, _ = c.Discard(n)

	b = b[:len(b)-1]
	return
}

func (codec *Linebreak) Encode(message []byte) []byte {
	return append(message, adapter.Newline...)
}

// HeaderLength 以头部4字节定义
type HeaderLength struct{}

func (codec *HeaderLength) Decode(c *Client) ([]byte, error) {
	buf, _ := c.Peek(bodySize)
	if len(buf) < bodySize {
		return nil, adapter.ErrorIncompletePacket
	}

	bodyLen := binary.BigEndian.Uint32(buf[:bodySize])
	msgLen := bodySize + int(bodyLen)
	if c.InboundBuffered() < msgLen {
		return nil, adapter.ErrorIncompletePacket
	}
	buf, _ = c.Peek(msgLen)
	_, _ = c.Discard(msgLen)

	return bytes.TrimSpace(bytes.Replace(buf[bodySize:msgLen], adapter.Newline, nil, -1)), nil
}

func (codec *HeaderLength) Encode(message []byte) []byte {
	msgLen := bodySize + len(message)

	data := make([]byte, msgLen)

	binary.BigEndian.PutUint32(data[:bodySize], uint32(len(message)))
	copy(data[bodySize:msgLen], message)

	return data
}
