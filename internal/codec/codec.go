// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-10
// Based on cabservd by liasica, magicrolan@qq.com.

package codec

import (
    "bufio"
    "bytes"
    "encoding/binary"
    "github.com/auroraride/adapter"
    "github.com/panjf2000/gnet/v2"
)

const (
    bodySize = 4
)

type Codec interface {
    Decode(conn gnet.Conn) (b []byte, err error)
    Encode(data []byte) (b []byte)
}

// Newline 以\n为分割处理
type Newline struct{}

func (codec *Newline) Decode(conn gnet.Conn) (b []byte, err error) {
    b, err = bufio.NewReader(conn).ReadBytes('\n')
    if err != nil {
        return
    }
    return
}

func (codec *Newline) Encode(message []byte) []byte {
    return append(message, adapter.Newline...)
}

// HeaderLength 以头部4字节定义
type HeaderLength struct{}

func (codec *HeaderLength) Decode(conn gnet.Conn) ([]byte, error) {
    buf, _ := conn.Peek(bodySize)
    if len(buf) < bodySize {
        return nil, adapter.ErrorIncompletePacket
    }

    bodyLen := binary.BigEndian.Uint32(buf[:bodySize])
    msgLen := bodySize + int(bodyLen)
    if conn.InboundBuffered() < msgLen {
        return nil, adapter.ErrorIncompletePacket
    }
    buf, _ = conn.Peek(msgLen)
    _, _ = conn.Discard(msgLen)

    return bytes.TrimSpace(bytes.Replace(buf[bodySize:msgLen], adapter.Newline, nil, -1)), nil
}

func (codec *HeaderLength) Encode(message []byte) []byte {
    msgLen := bodySize + len(message)

    data := make([]byte, msgLen)

    binary.BigEndian.PutUint32(data[:bodySize], uint32(len(message)))
    copy(data[bodySize:msgLen], message)

    return data
}
