// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-17
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

import (
	"context"
	"math"

	"github.com/auroraride/cabservd/internal/ent"
	"github.com/auroraride/cabservd/internal/ent/cabinet"
)

// TODO 初始化获取所有电柜信息
func initialize() {
	ctx := context.Background()
	var items []string
	_ = ent.Database.Cabinet.Query().Order(ent.Asc(cabinet.FieldID)).Select(cabinet.FieldSerial).Scan(ctx, &items)
	total := len(items)
	size := 10
	max := int(math.Ceil(float64(total) / float64(size)))
	for i := 0; i < max; i++ {
		from := i * size
		to := (i + 1) * size
		if i == max-1 {
			to = total
		}
		attrs, _ := FetchBusinessAttr(&BusinessAttrRequest{
			NeedCellAttr: 1,
			SnList:       items[from:to],
		})
		// 异步更新数据
		go batchUpdate(attrs)
	}
}

// 更新电柜
func batchUpdate(attrs CabAttrs) {
	for _, attr := range attrs {
		ent.UpdateCabinet(attr)
	}
}
