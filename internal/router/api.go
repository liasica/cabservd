// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-11
// Based on cabservd by liasica, magicrolan@qq.com.

package router

import (
    "github.com/auroraride/adapter"
    "github.com/auroraride/cabservd/internal/app"
    "github.com/auroraride/cabservd/internal/controller/api"
    "github.com/auroraride/cabservd/internal/g"
    mw "github.com/auroraride/cabservd/internal/middleware"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/labstack/gommon/log"
    "net/http"
)

func Start() {
    r := echo.New()

    r.HTTPErrorHandler = func(err error, c echo.Context) {
        _ = app.Context(c).SendResponse(http.StatusInternalServerError, err)
    }

    echo.NotFoundHandler = func(c echo.Context) error {
        return app.Context(c).SendResponse(http.StatusNotFound, adapter.NotFound)
    }

    r.Validator = app.NewValidator()

    r.Use(
        mw.Context(),
        mw.Recover(),
        mw.User(),

        middleware.GzipWithConfig(middleware.GzipConfig{
            Level: 5,
        }),
        // TODO body dump middleware
    )

    r.POST("/operate/bin", api.Bin.Operate)

    r.POST("/exchange/usable", api.Exchange.Usable)
    r.POST("/exchange/do", api.Exchange.Do)

    log.Fatal(r.Start(g.Config.Api.Bind))
    // router := gin.Default()
    // _ = router.SetTrustedProxies([]string{"192.168.1.2"})
    //
    // // 引入HTML模板
    // tmpls := template.Must(assets.LoadTemplates())
    // router.SetHTMLTemplate(tmpls)
    //
    // // demo路由
    // router.POST("/demo/control", api.Demo.Control)
    // router.GET("/demo/exchange", api.Demo.Exchange)
    // router.POST("/demo/start", api.Demo.Start)
    // router.POST("/demo/status", api.Demo.Status)
    //
    // srv := &http.Server{
    //     Addr:    g.Config.Api.Bind,
    //     Handler: router,
    // }

    // go func() {
    // 服务连接
    // if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
    //     log.Fatalf("listen: %s\n", err)
    // }
    // }()
    //
    // // 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
    // quit := make(chan os.Signal)
    // signal.Notify(quit, os.Interrupt)
    // <-quit
    // log.Println("关闭服务器...")
    //
    // ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    // defer cancel()
    // if err := srv.Shutdown(ctx); err != nil {
    //     log.Fatal("服务器被终止: ", err)
    // }
    // log.Println("服务器退出")

}
