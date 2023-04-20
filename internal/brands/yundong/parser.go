// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-15
// Based on yundong by liasica, magicrolan@qq.com.

package yundong

import "strconv"

type Parser struct {
}

type Field struct {
	Length int
	Name   string
}

type Data interface {
	SetField(b []byte, name string)
}

func Parse[T Data](data T, b []byte, fields []Field) (curr int) {
	curr = 0
	for _, field := range fields {
		data.SetField(b[curr:curr+field.Length], field.Name)
		curr += field.Length
	}
	return
}

func (p *Parser) Gsm(b byte) string {
	// TODO Save
	return "GSM信号强度=" + strconv.Itoa(int(b))
}

func (p *Parser) UpdateNotify(b byte) string {
	// TODO Save
	return "http升级结果通知=" + strconv.Itoa(int(b))
}
