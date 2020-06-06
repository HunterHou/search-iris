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

func (fc SettingController) GetSetting() {

	fc.Ctx.ViewData("BaseUrl", cons.BaseUrl)
	fc.Ctx.ViewData("Images", cons.Images)
	fc.Ctx.ViewData("Docs", cons.Docs)
	fc.Ctx.ViewData("Types", cons.Types)
	fc.Ctx.ViewData("BaseDir", cons.BaseDir)
	fc.Ctx.View("setting.html")
}

func (fc SettingController) PostSettingsave() {

	cons.BaseUrl = fc.Ctx.PostValue("BaseUrl")
	dirs := fc.Ctx.PostValues("BaseDir")
	cons.Images = fc.Ctx.PostValues("Images")
	cons.VideoTypes = fc.Ctx.PostValues("VideoTypes")
	cons.Docs = fc.Ctx.PostValues("Docs")
	cons.SetBaseDir(dirs)
	service.FlushDictionart(cons.DirFile)
	fc.Ctx.Redirect("/setting")
}
