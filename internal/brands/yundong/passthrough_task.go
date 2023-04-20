// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-17
// Based on yundong by liasica, magicrolan@qq.com.

package yundong

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/auroraride/adapter"
	"github.com/auroraride/adapter/log"
	"github.com/auroraride/cabservd/internal/codec"
	"github.com/auroraride/cabservd/internal/core"
	"go.uber.org/zap"
)

var (
	// 透传指令编号
	passthroughSN atomic.Value

	// 透传任务列表 sn(byte) => *PassthroughTask
	passthroughTasks sync.Map

	// 透传指令默认超时时间 (30s)
	defaultPassthroughTaskTimeout = 30 * time.Second
)

func init() {
	passthroughSN.Store(defaultPassthroughSN)
}

// PassthroughTask 透传指令任务
type PassthroughTask struct {
	// 任务编码
	sn byte

	// 超时逻辑
	timeout *time.Timer

	// 是否需要等待成功
	success chan bool

	// 是否需要格式化任务结果
	result any
}

// Delete 删除任务
func (t *PassthroughTask) Delete() {
	// 关闭通道
	adapter.ChSafeClose(t.success)

	// 关闭定时任务
	t.timeout.Stop()

	// 移除任务列表
	passthroughTasks.Delete(t.sn)
}

func (t *PassthroughTask) Wait() bool {
	if t == nil || t.success == nil {
		return true
	}
	v := <-t.success
	return v
}

func (t *PassthroughTask) WaitResult() any {
	t.Wait()
	return t.result
}

// 发送透传指令
func sendPassthrough(serial string, param any) (t *PassthroughTask, err error) {
	var (
		c    *core.Client
		cmd  PassthroughCommand
		skip bool
	)

	c, err = core.GetClient(serial)
	if err != nil {
		return
	}

	t = &PassthroughTask{
		success: make(chan bool),
		// 超时设置
		timeout: time.AfterFunc(defaultPassthroughTaskTimeout, func() {
			t.Delete()
		}),
	}

	// TODO 保存任务
	switch v := param.(type) {
	default:
		skip = true

	// 电表读数请求
	case *PassthroughAmmeterRequest:
		cmd = PassthroughCommandAmmeter
		t.result = new(AmmeterResponse)

	// 开仓请求
	case *PassthroughOpenDoorRequest:
		cmd = PassthroughCommandOpenDoor

	// 无需传参的指令
	case PassthroughCommand:
		cmd = v
		param = nil
		// 取消任务保存
		t = nil

	// 禁启仓请求
	case *PassthroughBinAvailableRequest:
		cmd = PassthroughCommandBinUsable
	}

	// 指令错误
	if skip {
		err = ErrPassthroughCommand
		return
	}

	sn := NextPassthroughSN()

	if t != nil {
		// 获取下一个指令编码
		t.sn = sn

		// 保存任务
		passthroughTasks.Store(t.sn, t)
	}

	msg := PassthroughRequest{
		C:     cmd,
		TM:    time.Now().UnixMilli(),
		Param: param,
	}

	c.Info("发送透传指令", log.Payload(msg))

	b := msg.Bytes()

	err = c.SendMessage(wrapPathroughData(CodePassthrough, sn, b))

	// 如果发送失败, 直接删除任务 (此时任务会返回失败)
	if err != nil && t != nil {
		t.Delete()
	}
	return
}

type PathroughData struct {
	code    byte
	sn      byte
	message []byte
}

func wrapPathroughData(code, sn byte, message []byte) *PathroughData {
	return &PathroughData{
		code:    code,
		sn:      sn,
		message: message,
	}
}

func (d *PathroughData) GetMessage(c codec.Codec) ([]byte, []zap.Field) {
	buf := adapter.NewBuffer()
	defer adapter.ReleaseBuffer(buf)

	buf.WriteByte(d.code)
	buf.WriteByte(d.sn)
	if d.message != nil {
		buf.Write(d.message)
	}

	b := c.Encode(buf.Bytes())

	return b, []zap.Field{
		zap.Int("code", int(d.code)),
		zap.Int("sn", int(d.sn)),
		log.Hex(b),
	}
}
