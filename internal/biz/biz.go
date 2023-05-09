// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-09
// Based on cabservd by liasica, magicrolan@qq.com.

package biz

import "sync"

// 业务任务列表
// 数据格式为 key(string) -> chan string(中断消息)
var tasks sync.Map

func Add(key string) (ch chan string) {
	ch = make(chan string)
	tasks.Store(key, ch)
	return
}

func Get(key string) (ch chan string, ok bool) {
	var v any
	v, ok = tasks.Load(key)
	if !ok {
		return
	}

	ch = v.(chan string)
	return
}

// Del 删除业务任务
// 一般用在中断或终止任务之后
func Del(key string, ch chan string) {
	close(ch)
	tasks.Delete(key)
}

// Send 发送消息并中断业务
func Send(key string, message string) {
	ch, ok := Get(key)
	if ok {
		ch <- message
	}
}
