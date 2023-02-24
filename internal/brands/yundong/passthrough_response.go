// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-17
// Based on yundong by liasica, magicrolan@qq.com.

package yundong

import (
    "github.com/auroraride/adapter"
    jsoniter "github.com/json-iterator/go"
    "go.uber.org/zap"
    "strconv"
)

// AmmeterResponse 电表数据返回
type AmmeterResponse struct {
    EMeterValue float64 `json:"eMeterValue"`
}

func (res *AmmeterResponse) Unmarshal(str string) {
    if str == "" {
        return
    }

    err := jsoniter.Unmarshal(adapter.ConvertString2Bytes(str), res)
    if err != nil {
        zap.L().Error("电表数据解析错误", zap.Error(err))
    }

    // TODO 存储电表数据
    zap.L().Info("电表数据为: " + strconv.FormatFloat(res.EMeterValue, 'f', 2, 64) + "kW·h")
}
