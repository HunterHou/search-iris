package controller

import (
	"../cons"
	"../service"
	"github.com/kataras/iris"
)

type SettingController struct {
	Ctx     iris.Context
	Service service.FileService
}

func (fc FileController) GetSetting() {

	fc.Ctx.ViewData("BaseUrl", cons.BaseUrl)
	fc.Ctx.ViewData("Images", cons.Images)
	fc.Ctx.ViewData("Docs", cons.Docs)
	fc.Ctx.ViewData("Types", cons.Types)
	fc.Ctx.ViewData("BaseDir", cons.BaseDir)
	fc.Ctx.View("setting.html")
}
