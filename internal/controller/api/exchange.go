// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-27
// Based on cabservd by liasica, magicrolan@qq.com.

package api

import "github.com/gin-gonic/gin"

type exchange struct {
}

func NewExchange() *exchange {
    return &exchange{}
}

func (*exchange) Usable(c *gin.Context) {

}
