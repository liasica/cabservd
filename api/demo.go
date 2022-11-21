// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-11
// Based on cabservd by liasica, magicrolan@qq.com.

package api

import (
    "fmt"
    "github.com/auroraride/cabservd/internal/core/kaixin"
    "github.com/gin-gonic/gin"
    "net/http"
)

type demo struct{}

var Demo = new(demo)

func (*demo) Control(c *gin.Context) {
    var req struct {
        DeviceId string                `json:"deviceId"`
        Params   []kaixin.ControlParam `json:"params"`
    }
    err := c.Bind(&req)
    if err == nil {
        err = kaixin.Control(req.DeviceId, kaixin.ControlRequest{ParamList: req.Params})
    }

    c.JSON(http.StatusOK, gin.H{
        // "result":
        "err": fmt.Sprintf("%v", err),
    })
}

func (*demo) Exchange(c *gin.Context) {
    var req struct {
        DeviceId string `json:"deviceId"`
    }
    err := c.Bind(&req)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "err": err,
        })
        return
    }
    // 获取空仓位
}
