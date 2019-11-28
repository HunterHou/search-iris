package main

import "github.com/kataras/iris"
import (
	"../cons"
	"../datasource"
	"../service"
	"../utils/collectionUtils"
)

func main() {
	app := iris.New()
	app.Handle("GET", "/", func(ctx iris.Context) {
		app.Logger().Info(ctx.Path())
		ctx.HTML("<h1>hello world!!!</h1>")

	})
	app.Get("/hello", func(ctx iris.Context) {
		app.Logger().Info(ctx.Path())
		name := ctx.URLParam("name")
		ctx.WriteString("hello " + name)
	})
	app.Get("/welcome", func(ctx iris.Context) {
		app.Logger().Info(ctx.Path())
		name := ctx.URLParam("name")
		ctx.WriteString("hello " + name)
	})
	app.Get("/json", func(ctx iris.Context) {
		app.Logger().Info(ctx.Path())
		var baseDir = "e:\\"
		var videoTypes = []string{cons.AVI, cons.MKV, cons.WMV, cons.MP4}
		var queryTypes []string
		queryTypes = collectionUtils.ExtandsItems(queryTypes, videoTypes)
		service.ScanDisk(baseDir, queryTypes)
		ctx.JSON(datasource.FileLib)
	})
	app.Run(iris.Addr("127.0.0.1:8000"), iris.WithConfiguration(iris.Configuration{
		DisableStartupLog:    false,
		FireMethodNotAllowed: false,
		TimeFormat:           "2019-11-10 18:10:33",
		Charset:              "uft-8",
	}))
}
