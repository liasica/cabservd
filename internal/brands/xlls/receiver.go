// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-17
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

import (
	"net/http"
	"sync"

	"github.com/auroraride/adapter"
	"github.com/labstack/echo/v4"
	"github.com/liasica/go-helpers/silk"
	"go.uber.org/zap"

	"github.com/auroraride/cabservd/internal/ent"
)

type receiver struct {
}

var (
	bizNotifiers sync.Map
)

// 存储业务回调通知器
func storeBizNotifier(key string, ch chan *BusinessNotify) {
	bizNotifiers.Store(key, ch)
}

// 关闭并移除业务回调通知器
func removeBizNotifier(key string, ch chan *BusinessNotify) {
	adapter.ChSafeClose(ch)
	bizNotifiers.Delete(key)
}

// 获取业务回调通知器
func loadBizNotifier(key string) chan *BusinessNotify {
	v, _ := bizNotifiers.Load(key)
	if v == nil {
		return nil
	}
	return v.(chan *BusinessNotify)
}

// 发送回调结果
func (r *receiver) send(c echo.Context, n Notifyer) error {
	return c.JSON(http.StatusOK, &NotifyResult[any]{
		RequestID: n.GetRequestID(),
	})
}

// 获取回调数据并存储到上下文中
func getReceiverData[T any](c echo.Context) (n *T, err error) {
	n = new(T)
	err = c.Bind(n)
	if err != nil {
		zap.L().Error("[回调] 数据格式化失败", zap.Error(err))
		return
	}
	c.Set("payload", n)
	return
}

// 仓位变动回调
func (r *receiver) onBin(c echo.Context) error {
	data, err := getReceiverData[BinNotify](c)
	if err != nil {
		return err
	}

	if data.Sn == "" {
		zap.L().Error("仓位仓位上报无sn")
		return nil
	}

	cab := &CabAttr{Sn: silk.String(data.Sn), CellAttrList: BinAttrs{&data.BinAttr}}
	ent.UpdateCabinet(cab)

	return r.send(c, data)
}

// 柜机变动回调
func (r *receiver) onCab(c echo.Context) error {
	data, err := getReceiverData[CabNotify](c)
	if err != nil {
		return err
	}

	ent.UpdateCabinet(data)

	return r.send(c, data)
}

// 业务回调
func (r *receiver) onBusinuess(c echo.Context) error {
	data, err := getReceiverData[BusinessNotify](c)
	if err != nil {
		return err
	}

	// 尝试获取通知器并传递数据
	ch := loadBizNotifier(data.OrderNo)
	if ch != nil {
		adapter.ChSafeSend(ch, data)
	}

	return r.send(c, data)
}
