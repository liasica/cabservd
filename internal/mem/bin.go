// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on cabservd by liasica, magicrolan@qq.com.

package mem

import (
	"strconv"
)

func binKey(serial string, ordinal int) string {
	return serial + "-" + strconv.Itoa(ordinal)
}
