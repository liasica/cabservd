// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-29
// Based on cabservd by liasica, magicrolan@qq.com.

package tower

import (
	"fmt"
	"strconv"

	"go.uber.org/zap"
)

type Attributes []*Attribute

// Attribute 属性信息
type Attribute struct {
	SignalData
	DoorID string `json:"doorId,omitempty"` // 柜门ID (可为空)
}

func (d SignalData) ValueString() (str string) {
	str = fmt.Sprintf("%v", d.Value)
	if str == "null" {
		str = ""
	}
	return
}

func (a *Attribute) GetOrdinal() (ordinal int, exists bool) {
	if a.DoorID == "" {
		return
	}

	var err error
	ordinal, err = strconv.Atoi(a.DoorID)
	if err != nil {
		zap.L().Error("仓位解析失败", zap.Error(err))
		return
	}

	exists = true
	return
}
