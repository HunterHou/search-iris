package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/mvc"
	"strings"
)
import (
	"./cons"
	"./utils"
	"./web/controller"
	"path/filepath"
)

var curDir string
var staticDir string


func init() {
	curDir, _ := filepath.Abs(".")
	if !strings.HasSuffix(curDir, "src") {
		curDir += "/src"
	}
	staticDir = curDir + "/web/static"
	cons.Play =utils.ImageToString(staticDir +"/image/play.jpg")
	cons.Open =utils.ImageToString(staticDir +"/image/open.jpg")
	cons.Change =utils.ImageToString(staticDir +"/image/change.jpg")
	cons.Replay =utils.ImageToString(staticDir +"/image/replay.jpg")
}

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

	app.RegisterView(iris.Django(staticDir, ".html"))
	app.HandleDir("/", staticDir)
	app.Logger().SetLevel("debug")
	mvc.New(app).Handle(new(controller.TestController))
	mvc.New(app).Handle(new(controller.FileController))
	app.Run(iris.Addr("127.0.0.1:8000"), iris.WithConfiguration(iris.Configuration{
		DisableStartupLog:    false,
		FireMethodNotAllowed: false,
		TimeFormat:           "2019-11-10 18:10:33",
		Charset:              "uft-8",
	}))

	utils.ExecCmdExplorer("http://127.0.0.1:8000/views")
}
