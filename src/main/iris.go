package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/mvc"
)
import (
	"../web/controller"
)

func main() {
	app := iris.New()
	app.Handle("GET", "/", func(ctx iris.Context) {
		app.Logger().Info(ctx.Path())
		ctx.HTML("<h1>hello world!!!</h1>")

	})
	customLogger := logger.New(logger.Config{
		Status:             true,
		IP:                 true,
		Method:             true,
		Path:               true,
		Query:              false,
		Columns:            false,
		MessageContextKeys: nil,
		MessageHeaderKeys:  nil,
		LogFunc:            nil,
		LogFuncCtx:         nil,
		Skippers:           nil,
	})
	app.Use(customLogger)

	app.RegisterView(iris.Django("./views", ".html"))
	app.Logger().SetLevel("debug")
	mvc.New(app).Handle(new(controller.TestController))
	mvc.New(app).Handle(new(controller.FileController))
	app.Run(iris.Addr("127.0.0.1:8000"), iris.WithConfiguration(iris.Configuration{
		DisableStartupLog:    false,
		FireMethodNotAllowed: false,
		TimeFormat:           "2019-11-10 18:10:33",
		Charset:              "uft-8",
	}))
}
