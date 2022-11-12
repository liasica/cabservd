// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-06
// Based on cabservd by liasica, magicrolan@qq.com.

package core

import "github.com/google/uuid"

type Context struct {
    ID string // 链接ID
}

func NewContext() *Context {
    return &Context{ID: uuid.New().String()}
}
