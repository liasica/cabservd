// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-10
// Based on cabservd by liasica, magicrolan@qq.com.

package core

import (
    "bufio"
    "bytes"
    "encoding/binary"
    "github.com/auroraride/cabservd/internal/errs"
    "github.com/panjf2000/gnet/v2"
)

var (
    newline = []byte{'\n'}
    space   = []byte{' '}
)

const (
    bodySize = 4
)

type Codec interface {
    Decode(c gnet.Conn) (b []byte, err error)
}

// Newline 以\n为分割处理
type Newline struct{}

func (codec *Newline) Decode(c gnet.Conn) (b []byte, err error) {
    b, err = bufio.NewReader(c).ReadBytes('\n')
    if err != nil {
        return
    }
    return
}

// HeaderLength 以头部4字节定义
type HeaderLength struct{}

func (codec *HeaderLength) Decode(c gnet.Conn) ([]byte, error) {
    buf, _ := c.Peek(bodySize)
    if len(buf) < bodySize {
        return nil, errs.IncompletePacket
    }

    bodyLen := binary.BigEndian.Uint32(buf[:bodySize])
    msgLen := bodySize + int(bodyLen)
    if c.InboundBuffered() < msgLen {
        return nil, errs.IncompletePacket
    }
    buf, _ = c.Peek(msgLen)
    _, _ = c.Discard(msgLen)

    return bytes.TrimSpace(bytes.Replace(buf[bodySize:msgLen], newline, nil, -1)), nil
}
