package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>hello world!!!</h1>")
	})
	app.Get("/go", func(ctx iris.Context) { ctx.WriteString("hello go") })
	app.Get("/json", func(ctx iris.Context) { ctx.JSON(iris.Map{"message": "hello json"}) })
	app.Run(iris.Addr("127.0.0.1:8080"), iris.WithConfiguration(iris.Configuration{
		DisableStartupLog:    false,
		FireMethodNotAllowed: false,
		TimeFormat:           "2019-11-10 18:10:33",
		Charset:              "uft-8",
	}))
}
