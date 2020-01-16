package controller

import (
	"fmt"

	"../../cons"
	"../../datamodels"
	"../../datasource"
	"../../service"
	"../../utils"
	"github.com/kataras/iris"
)

type FileController struct {
	Ctx     iris.Context
	Service service.FileService
}

func (fc FileController) GetAll() []datamodels.File {
	fc.Service.ScanAll()
	list := datasource.FileList
	return list
}
func (fc FileController) GetViews() {
	fc.Service.ScanAll()
	keyWord := fc.Ctx.URLParam("keyWord")
	fmt.Println("keyWord:", keyWord)
	datas := fc.Service.SearchByKeyWord(datasource.FileList, keyWord)

	fc.Ctx.ViewData("playIcon", cons.Play)
	fc.Ctx.ViewData("changeIcon", cons.Change)
	fc.Ctx.ViewData("openIcon", cons.Open)
	fc.Ctx.ViewData("replayIcon", cons.Replay)

	fc.Ctx.ViewData("dirList", cons.BaseDir)
	fc.Ctx.ViewData("totalCnt", len(datas))
	fc.Ctx.ViewData("datas", datas)
	fc.Ctx.ViewData("title", "文件列表")
	fc.Ctx.View("main.html")
}
func (fc FileController) PostPlay() {
	id := fc.Ctx.PostValue("id")
	file := fc.Service.FindOne(id)
	fmt.Println("id:", file.Path)
	utils.ExecCmdStart(file.Path)
}
