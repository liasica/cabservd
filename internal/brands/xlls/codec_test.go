// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-16
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

import (
	"bytes"
	"testing"

	"github.com/auroraride/adapter"
)

var (
	sa = []byte{71, 69, 84}
	sb = []byte{71, 69, 84}
	sc = []byte{80, 79, 83, 84}

	ssa = "GET"
	ssb = "GET"
	ssc = "POST"
)

func BenchmarkCompare1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytes.Compare(sa, sb)
		bytes.Compare(sa, sc)
	}
}

func BenchmarkCompare2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ssa == adapter.ConvertBytes2String(sb)
		_ = ssa == adapter.ConvertBytes2String(sc)
	}
}
