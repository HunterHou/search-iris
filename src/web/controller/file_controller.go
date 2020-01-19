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

func (fc FileController) GetMovies() utils.Page {
	if len(datasource.FileList) == 0 {
		fc.Service.ScanAll()
	}
	list := datasource.FileList
	result := utils.NewPage()
	result.CurCnt = len(list)
	result.Data = list
	return result
}
func (fc FileController) GetActess() utils.Page {
	if len(datasource.FileList) == 0 {
		fc.Service.ScanAll()
	}
	list := datasource.ActressLib
	result := utils.NewPage()
	result.CurCnt = len(list)
	result.Data = list
	return result
}
func (fc FileController) GetSupplier() utils.Page {
	if len(datasource.SupplierLib) == 0 {
		fc.Service.ScanAll()
	}
	list := datasource.SupplierLib
	result := utils.NewPage()
	result.CurCnt = len(list)
	result.Data = list
	return result
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
func (fc FileController) PostDelete() {
	id := fc.Ctx.PostValue("id")
	fc.Service.Delete(id)
	result := utils.Success()
	fc.Ctx.JSON(result)
}

func (fc FileController) PostSync() {
	id := fc.Ctx.PostValue("id")
	curFile := fc.Service.FindOne(id)
	result, newFile := fc.Service.RequestToFile(curFile)
	fc.Service.MoveCut(curFile, newFile)
	fc.Ctx.JSON(result)
}

func (fc FileController) GetViews() {
	if len(datasource.FileList) == 0 {
		fc.Service.ScanAll()
	}
	keyWord := fc.Ctx.URLParam("keyWord")
	pageNo, errNo := fc.Ctx.URLParamInt("pageNo")
	if errNo != nil || pageNo == 0 {
		pageNo = 1
	}
	pageSize, errSize := fc.Ctx.URLParamInt("pageSize")
	if errSize != nil || pageSize == 0 {
		pageSize = 100
	}
	totalCnt := len(datasource.FileList)
	datas := fc.Service.SearchByKeyWord(datasource.FileList, keyWord)
	resultCnt := len(datas)
	fc.Service.SortMovies(datas)
	datas = fc.Service.GetPage(datas, pageNo, pageSize)
	curCnt := len(datas)

	page := utils.NewPage()
	page.KeyWord = keyWord
	page.PageNo = pageNo
	page.PageSize = pageSize
	page.Data = datas
	page.TotalCnt = totalCnt
	page.CurCnt = curCnt
	page.ResultCnt = resultCnt
	page = page.SetResultCnt(resultCnt)

	fc.Ctx.ViewData("playIcon", cons.Play)
	fc.Ctx.ViewData("changeIcon", cons.Change)
	fc.Ctx.ViewData("openIcon", cons.Open)
	fc.Ctx.ViewData("replayIcon", cons.Replay)
	fc.Ctx.ViewData("closeIcon", cons.Close)

	fc.Ctx.ViewData("page", page)
	fc.Ctx.ViewData("curPage", page.PageNo)
	fc.Ctx.ViewData("dirList", cons.BaseDir)
	fc.Ctx.ViewData("title", "文件列表")
	fc.Ctx.View("main.html")
}
func (fc FileController) GetView() {
	if len(datasource.FileList) == 0 {
		fc.Service.ScanAll()
	}
	datas := []datamodels.Actress{}
	list := datasource.ActressLib
	for _, data := range list {
		datas = append(datas, data)
	}
	fc.Service.SortAct(datas)
	totalCnt := len(datas)
	resultCnt := len(datas)
	curCnt := len(datas)

	page := utils.NewPage()
	page.Data = datas
	page.TotalCnt = totalCnt
	page.CurCnt = curCnt
	page.ResultCnt = resultCnt

	fc.Ctx.ViewData("page", page)
	fc.Ctx.ViewData("curPage", page.PageNo)
	fc.Ctx.ViewData("dirList", cons.BaseDir)
	fc.Ctx.ViewData("title", "列表")
	fc.Ctx.View("act.html")
}
