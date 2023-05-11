// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-11
// Based on cabservd by liasica, magicrolan@qq.com.

package biz

var (
	exchangeFirstStepRetryTimes = 1
	exchangeThirdStepRetryTimes = 1
)

type retryer struct {
}

func New() {
}

func Next() bool {
	return false
}

func WithexchangeFirstStepRetryTimes(n int) {
	exchangeFirstStepRetryTimes = n
}
