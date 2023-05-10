// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-09
// Based on cabservd by liasica, magicrolan@qq.com.

package biz

import (
	"sync"

	"github.com/auroraride/adapter"
	"github.com/auroraride/adapter/defs/cabdef"
	"github.com/auroraride/adapter/rpc/pb"
	"github.com/google/uuid"
)

// 业务任务列表
// 数据格式为 serial+ordinal -> *Task
var tasks sync.Map

type Task struct {
	Biz      *pb.CabinetBiz
	Business adapter.Business
	Operate  cabdef.Operate

	Key uuid.UUID
	// 中断器, 发送中断消息
	Interrupter chan string
}

func Create(serial string, ordinal int, business adapter.Business, operate cabdef.Operate, user *adapter.User) (t *Task) {
	t = &Task{
		Key:         uuid.New(),
		Interrupter: make(chan string),
		Business:    business,
		Operate:     operate,

		Biz: &pb.CabinetBiz{
			Serial:  serial,
			Ordinal: int32(ordinal),
			Desc:    business.Text() + ":" + operate.Text(),
		},
	}
	if user != nil {
		t.Biz.User = user.String()
	}
	tasks.Store(t.Key, t)
	return
}

// Del 删除业务任务
// 一般用在中断或终止任务之后
func (t *Task) Del() {
	close(t.Interrupter)
	tasks.Delete(t.Key)
}

// Interrupt 发送中断消息
func (t *Task) Interrupt(message string) {
	t.Interrupter <- message
}

// Interrupt 发送消息并中断该电柜所有业务
func Interrupt(serial string, message string) (items []*Task) {
	// 查询该电柜所有任务并终止
	tasks.Range(func(_, v any) bool {
		item := v.(*Task)
		if item.Biz.Serial == serial {
			items = append(items, item)
			item.Interrupter <- message
		}
		return true
	})
	return
}

// List 该电柜正在执行的任务列表
func List(serial string) (items []*Task) {
	tasks.Range(func(_, v any) bool {
		item := v.(*Task)
		if item.Biz.Serial == serial {
			items = append(items, item)
		}
		return true
	})
	return
}
