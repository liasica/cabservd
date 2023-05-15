// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-02
// Based on cabservd by liasica, magicrolan@qq.com.

package core

import (
	"github.com/panjf2000/gnet/v2"
	"go.uber.org/zap"

	"github.com/auroraride/cabservd/internal/codec"
)

func Start(addr string, bean Hook, codec codec.Codec, options ...Option) {
	Hub = &hub{
		addr:  addr,
		Bean:  bean,
		codec: codec,
	}

	for _, option := range options {
		option.apply(Hub)
	}

	// go Hub.deadCheck()

	zap.L().Fatal(gnet.Run(
		Hub,
		Hub.addr,
		gnet.WithMulticore(true),
		gnet.WithReuseAddr(true),
	).Error())
}

// 每隔1分钟标记20分之前更新的电柜为离线
// TODO 是否发送消息
func (h *hub) deadCheck() {
	// ticker := time.NewTicker(time.Minute)
	// for {
	//     select {
	//     case t := <-ticker.C:
	//         _ = ent.Database.GetCabinet.Update().
	//             Where(
	//                 cabinet.Brand(g.Config.Brand),
	//                 cabinet.UpdatedAtLT(t.Add(-20*time.Minute)),
	//             ).
	//             SetOnline(false).
	//             Exec(context.Background())
	//     }
	// }
}
