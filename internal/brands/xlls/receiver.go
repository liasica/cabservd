// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-17
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/liasica/go-helpers/silk"
	"go.uber.org/zap"

	"github.com/auroraride/cabservd/internal/ent"
)

type receiver struct {
}

func (r *receiver) send(c echo.Context, n Notifyer) error {
	return c.JSON(http.StatusOK, &NotifyResult[any]{
		RequestID: n.GetRequestID(),
	})
}

func (r *receiver) onBin(c echo.Context) (err error) {
	n := new(BinNotify)
	err = c.Bind(n)
	if err != nil {
		zap.L().Error("仓位属性格式化失败", zap.Error(err))
		return
	}

	if n.Sn == "" {
		zap.L().Error("仓位仓位上报无sn")
		return
	}

	cab := &CabAttr{Sn: silk.String(n.Sn), CellAttrList: BinAttrs{&n.BinAttr}}
	ent.UpdateCabinet(cab)
	c.Set("payload", n)

	return r.send(c, n)
}

func (r *receiver) onCab(c echo.Context) (err error) {
	n := new(CabNotify)
	err = c.Bind(n)
	if err != nil {
		zap.L().Error("电柜属性格式化失败", zap.Error(err), zap.String("id", n.RequestID))
		return
	}

	ent.UpdateCabinet(n)
	c.Set("payload", n)

	return r.send(c, n)
}
