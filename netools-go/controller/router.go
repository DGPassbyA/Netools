package controller

import (
	"main/controller/api"
	"main/middleware"
	"main/config"
	"net/http"
	"os"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
	"github.com/sirupsen/logrus"
)

func Router() {
	app := iris.New()
	app.Logger().SetLevel("warn")
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
		MaxAge:           600,
		AllowedMethods:   []string{iris.MethodGet, iris.MethodPost, iris.MethodOptions, iris.MethodHead, iris.MethodDelete, iris.MethodPut},
		AllowedHeaders:   []string{"*"},
	}))

	app.AllowMethods(iris.MethodOptions)

	app.Any("/", func(ctx iris.Context) {
		_, _ = ctx.HTML("<h1>Powered by HoshinoMaster</h1>")
	})

	mvc.Configure(app.Party("/api"), func(m *mvc.Application) {
		m.Party("/test", middleware.CheckAuthToken).Handle(new(api.TestController))
		m.Party("/text").Handle(new(api.TextController))
		m.Party("/file").Handle(new(api.FileController))
	})

	server := &http.Server{Addr: config.ServerAddr}
	err := app.Run(iris.Server(server), iris.WithConfiguration(iris.Configuration{
		DisableStartupLog:                 false,
		DisableInterruptHandler:           false,
		DisablePathCorrection:             false,
		EnablePathEscape:                  false,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: false,
		DisableAutoFireStatusCode:         false,
		EnableOptimizations:               true,
		TimeFormat:                        "2006-01-02 15:04:05",
		Charset:                           "UTF-8",
	}))
	if err != nil {
		logrus.Error(err)
		os.Exit(-1)
	}
}
