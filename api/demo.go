// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-11
// Based on cabservd by liasica, magicrolan@qq.com.

package api

import (
    "context"
    "fmt"
    "github.com/auroraride/cabservd/internal/core/kaixin"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/ent/cabinetbin"
    "github.com/auroraride/cabservd/internal/errs"
    "github.com/gin-gonic/gin"
    "net/http"
)

type demo struct {
    start bool
}

var Demo = new(demo)

func (*demo) Control(c *gin.Context) {
    var req struct {
        SN     string                `form:"sn" json:"sn"`
        Params []kaixin.ControlParam `json:"params"`
    }
    err := c.Bind(&req)
    if err == nil {
        err = kaixin.Control(req.SN, kaixin.ControlRequest{ParamList: req.Params})
    }

    c.JSON(http.StatusOK, gin.H{
        // "result":
        "err": fmt.Sprintf("%v", err),
    })
}

func (*demo) Exchange(c *gin.Context) {
    var req struct {
        SN string `form:"sn" json:"sn"`
    }

    err := c.Bind(&req)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err,
        })
        return
    }

    // 获取仓位状态
    var items ent.CabinetBins
    items, err = ent.Database.CabinetBin.Query().
        Where(cabinetbin.Sn(req.SN)).
        Order(ent.Asc(cabinetbin.FieldIndex)).
        All(context.Background())

    c.HTML(http.StatusOK, "demo/exchange.go.html", gin.H{
        "items": items,
        "sn":    req.SN,
    })
}

type bin struct {
    Name string  `json:"name"`
    Soc  float64 `json:"soc"`
}

func (d *demo) Start(c *gin.Context) {
    var req struct {
        SN string `form:"sn" json:"sn"`
    }
    err := c.Bind(&req)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": errs.ParamValidateFailed})
        return
    }

    // 获取仓位状态
    var items ent.CabinetBins
    items, err = ent.Database.CabinetBin.Query().
        Where(cabinetbin.Sn(req.SN)).
        Order(ent.Desc(cabinetbin.FieldSoc)).
        All(context.Background())

    if len(items) == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": errs.CabinetNoFully.Error()})
        return
    }

    // 获取满电仓位
    max := items[0]
    if max.BatterySn == "" || max.Soc == 0 || max.Soc < 50 {
        c.JSON(http.StatusBadRequest, gin.H{"error": errs.CabinetNoFully.Error()})
        return
    }
    fully := bin{
        Name: max.Name,
        Soc:  max.Soc,
    }

    // 获取空电仓位
    min := items[len(items)-1]
    if min.BatterySn != "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": errs.CabinetNoEmpty.Error()})
        return
    }
    empty := bin{
        Name: min.Name,
        Soc:  min.Soc,
    }

    // 开始执行换电任务
    err = d.run(req.SN, fully, empty)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 返回
    c.JSON(http.StatusOK, gin.H{
        "fully": fully,
        "empty": empty,
    })
}

func (*demo) Status(c *gin.Context) {
}

func (d *demo) run(sn string, fully, empty bin) (err error) {
    // 第一步, 开启空电仓门
    // 第二步, 识别仓门是否关闭
    // 第三步, 开启满电仓门
    // 第四步, 识别电池取出并关闭仓门
}
