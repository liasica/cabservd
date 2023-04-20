// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-21
// Based on cabservd by liasica, magicrolan@qq.com.

package yundong

import (
	"bytes"
	"encoding/binary"

	"github.com/auroraride/adapter"
	"github.com/panjf2000/gnet/v2"
)

type Codec struct {
	pkgStart   []byte
	headerSize int
}

func NewCodec() *Codec {
	return &Codec{
		pkgStart:   []byte{0xAA, 0x55},
		headerSize: 6,
	}
}

func (codec *Codec) Decode(conn gnet.Conn) (data []byte, err error) {
	buf, _ := conn.Peek(codec.headerSize)
	if len(buf) < codec.headerSize {
		return nil, adapter.ErrorIncompletePacket
	}

	if buf[0] != codec.pkgStart[0] || buf[1] != codec.pkgStart[1] {
		return nil, adapter.ErrorIncorrectPacket
	}

	bodyLen := int(binary.BigEndian.Uint16(buf[4:]))
	msgLen := codec.headerSize + bodyLen
	if conn.InboundBuffered() < msgLen {
		return nil, adapter.ErrorIncompletePacket
	}
	buf, _ = conn.Peek(msgLen)
	_, _ = conn.Discard(msgLen)

	data = make([]byte, 2+bodyLen)
	data[0] = buf[2]
	data[1] = buf[3]

	copy(data[2:], buf[codec.headerSize:])
	return
}

func (codec *Codec) Encode(b []byte) []byte {
	// buf := adapter.NewBuffer()
	// defer adapter.ReleaseBuffer(buf)
	buf := bytes.Buffer{}
	buf.Write(codec.pkgStart)

	// code
	buf.WriteByte(b[0])

	// sn
	buf.WriteByte(b[1])

	l := make([]byte, 2)
	binary.BigEndian.PutUint16(l, uint16(len(b)-2))
	buf.Write(l)
	buf.Write(b[2:])

	return buf.Bytes()
}
