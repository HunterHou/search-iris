package controller

import (
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

func (fc FileController) PostPlay() {
	id := fc.Ctx.PostValue("id")
	file := fc.Service.FindOne(id)
	utils.ExecCmdStart(file.Path)
}

func (fc FileController) PostOpendir() {
	id := fc.Ctx.PostValue("id")
	file := fc.Service.FindOne(id)
	utils.ExecCmdStart(file.DirPath)
}

func (fc FileController) PostInfo() {
	id := fc.Ctx.PostValue("id")
	file := fc.Service.FindOne(id)
	file.Png = file.PngBase64()
	file.Jpg = utils.ImageToString(file.Jpg)
	fc.Ctx.JSON(file)
}

func (fc FileController) GetFresh() {
	fc.Service.ScanAll()
	result := utils.Success()
	fc.Ctx.JSON(result)
}

func (fc FileController) GetViews() {
	fc.Service.ScanAll()
	keyWord := fc.Ctx.URLParam("keyWord")
	pageNo, errNo := fc.Ctx.URLParamInt("pageNo")
	if errNo != nil {
		pageNo = 1
	}
	pageSize, errSize := fc.Ctx.URLParamInt("pageSize")
	if errSize != nil {
		pageSize = 100
	}
	totalCnt := len(datasource.FileList)
	datas := fc.Service.SearchByKeyWord(datasource.FileList, keyWord)
	datas = fc.Service.GetPage(datas, pageNo, pageSize)
	fc.Service.SortItems(datas)

	page := utils.NewPage()
	page.KeyWord = keyWord
	page.PageNo = pageNo
	page.PageSize = pageSize
	page.Data = datas
	page = page.SetTotalCnt(totalCnt)

	fc.Ctx.ViewData("playIcon", cons.Play)
	fc.Ctx.ViewData("changeIcon", cons.Change)
	fc.Ctx.ViewData("openIcon", cons.Open)
	fc.Ctx.ViewData("replayIcon", cons.Replay)

	fc.Ctx.ViewData("page", page)
	fc.Ctx.ViewData("curPage", page.PageNo)
	fc.Ctx.ViewData("dirList", cons.BaseDir)
	fc.Ctx.ViewData("title", "文件列表")
	fc.Ctx.View("main.html")
}
