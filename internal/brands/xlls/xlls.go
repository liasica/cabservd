// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-15
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

import (
	"net/http"

	"github.com/auroraride/adapter"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/color"
	"go.uber.org/zap"

	"github.com/auroraride/cabservd/internal/g"
)

var (
	appID     string
	appSecret []byte
	baseURL   string
	version   string
)

func Start() {
	conf := adapter.GetKoanf()

	// 读取配置
	appID = conf.Get("xiliulou.appId").(string)
	appSecret = []byte(conf.Get("xiliulou.appSecret").(string))
	baseURL = conf.Get("xiliulou.server").(string)
	version = conf.Get("xiliulou.version").(string)

	// 初始化先于回调调用并保存
	// TODO 待考虑: 是否使用队列保存回调信息等待初始化完成后统一更新
	initialize()

	// 创建echo
	e := echo.New()

	// 默认json序列化工具
	e.JSONSerializer = &adapter.DefaultJSONSerializer{}

	// 获取远程IP
	e.IPExtractor = echo.ExtractIPFromXFFHeader()

	// 使用中间件记录回调日志
	e.Use(dump())

	// 覆盖echo打印
	e.HideBanner = true
	e.HidePort = true
	colorer := color.New()
	bind := g.Config.Tcp.Bind
	colorer.Printf("⇨ 西六楼对接启动于 %s\n", colorer.Green(bind))

	// 创建接收器以便于接收西六楼反馈消息
	r := new(receiver)

	// 硬件操作结果通知
	e.POST(pathHardwareOperation, r.onImperfectMessage)

	// 业务结果通知
	e.POST(pathBusinesss, r.onBusinuess)

	// 离线换电结果通知
	e.POST(pathOfflineExchange, r.onImperfectMessage)

	// 格挡状态变化通知
	e.POST(pathCellChange, r.onBin)

	// 柜机状态变化通知
	e.POST(pathCabinetChange, r.onCab)

	// 硬件故障通知
	e.POST(pathHardwareFault, r.onImperfectMessage)

	// 自助开仓回调通知
	e.POST(pathSelfServiceOpen, r.onImperfectMessage)

	if err := e.Start(bind); err != nil && err != http.ErrServerClosed {
		zap.L().Fatal(err.Error())
	}
}
