// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-21
// Based on cabservd by liasica, magicrolan@qq.com.

package yundong

import (
	"encoding/binary"
	"errors"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/auroraride/adapter"
	"github.com/auroraride/adapter/defs/cabdef"
	"github.com/auroraride/adapter/log"
	"go.uber.org/zap"

	"github.com/auroraride/cabservd/internal/codec"
	"github.com/auroraride/cabservd/internal/core"
	"github.com/auroraride/cabservd/internal/g"
)

type Handler struct {
	core.Bean

	serverInfo *Server
	parser     *Parser
}

type Server struct {
	IP      []byte
	Port    []byte
	Address []byte
}

func New() (core.Hook, codec.Codec) {
	ip := net.ParseIP("39.106.77.239")
	arr := strings.Split(g.Config.Tcp.Bind, ":")
	port, _ := strconv.ParseUint(arr[1], 10, 64)

	p := make([]byte, 2)
	binary.BigEndian.PutUint16(p, uint16(port))

	info := &Server{
		IP:   ip[12:16],
		Port: p,
	}

	buf := adapter.NewBuffer()
	defer adapter.ReleaseBuffer(buf)

	t := make([]byte, 4)
	binary.BigEndian.PutUint32(t, uint32(time.Now().Unix()))
	buf.Write(t)
	buf.Write(info.IP)
	buf.Write(info.Port)

	info.Address = buf.Bytes()

	return &Handler{
		parser:     &Parser{},
		serverInfo: info,
	}, NewCodec()
}

// GetEmptyDeviation TODO 后续做在数据库中
func (h *Handler) GetEmptyDeviation() (voltage, current float64) {
	voltage = 45
	current = -1
	return
}

func (h *Handler) OnMessage(c *core.Client, data []byte) (serial string, res core.ResponseMessenger, fields []zap.Field, err error) {
	var (
		decoded string
		message []byte
		code    = data[0]
		sn      = data[1]
		b       = data[2:]
	)

	// 是否已发送响应报文
	defer func() {
		if c.Serial != "" {
			serial = c.Serial
		}
		if err == nil {
			fields = []zap.Field{
				zap.Int("code", int(code)),
				zap.Int("sn", int(sn)),
				log.Hex(data),
			}
			if decoded != "" {
				fields = append(fields, zap.String("decoded", decoded))
			}
			res = wrapPathroughData(code, sn, message)
		}
	}()

	switch code {
	default:
		err = errors.New("指令错误: " + strconv.Itoa(int(code)))
		return
	case CodeAlarm:
		// TODO 具体格式
		decoded = h.parser.Alarm(b[0])
	case CodeHeartbeat:
		decoded = h.parser.Gsm(b[0])
	case CodeUpgradeNotify:
		decoded = h.parser.UpdateNotify(b[0])
	case CodeLogin:
		message = h.serverInfo.Address
		serial, decoded = h.parser.Login(b)
	case CodePeriodMsg:
	case CodePassthrough:
		unpackPassthroughResponse(sn, b)
	case CodeDoorEvent:
	case CodeStartMaintaining:
		// go c.Send(code, sn, []byte{0, 0})
	case CodeStopMaintaining:
		// go c.Send(code, sn, []byte{0, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63})
	case CodeRequestExchange:
		// go c.Send(code, sn, []byte{48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63})
	case CodeGenerateOrder:
	case CodePayOrder:
	case CodeExchangeEnd:
	case CodeExchangeCancel:
	case CodeExchangeException:
	case CodeGetQrcode:
		message = h.getQrcode(c.Serial)
	case CodeGetAllowReserCount:
	case CodePeriodMsgV2:
	case CodePeriodMsgBlankBox:
	case CodeUnknown:
	case CodePeriodMsgV3:
	case CodePeriodMsgV4:
		decoded = h.parser.CabinetData(h, c.Serial, b)
	case CodeRequestExchangeV2:
	case CodeGenerateOrderV2:
	case CodeDoorEventV2:
	case CodeBinDisable:
	}

	return
}

func (h *Handler) SendOperate(serial string, typ cabdef.Operate, ordinal int, times int) (err error) {
	index := ordinal - 1
	switch typ {
	default:
		return adapter.ErrorOperateCommand
	case cabdef.OperateDoorOpen:
		CommandOpen(serial, index, times)
	case cabdef.OperateBinDisable:
		CommandDisable(serial, index, times)
	case cabdef.OperateBinEnable:
		CommandEnable(serial, index, times)
	}
	return
}

func (h *Handler) getQrcode(sn string) []byte {
	var (
		sb [16]byte
		qr [128]byte
	)
	copy(sb[:], adapter.ConvertString2Bytes(sn))
	// copy(qr[:], adapter.ConvertString2Bytes(sn))
	return append(sb[:], qr[:]...)
}
