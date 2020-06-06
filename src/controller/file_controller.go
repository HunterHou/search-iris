package controller

import (
	"../cons"
	"../datamodels"
	"../datasource"
	"../service"
	"../utils"
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

func (fc FileController) GetPlay() {
	id := fc.Ctx.URLParam("id")
	file := fc.Service.FindOne(id)
	utils.ExecCmdStart(file.Path)
	//files, err := ioutil.ReadFile(file.Path)
	//if err!=nil {
	//	fmt.Println(err)
	//}

	// fc.Ctx.ContentType("video/mp4")
	// f,_:=os.Open(file.Path)
	// sourceBuffer := make([]byte, 5000)
	// f.Read(sourceBuffer)
	// fc.Ctx.Write(sourceBuffer)

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
	datasource.SortMovieForce()
	result := utils.Success()
	fc.Ctx.JSON(result)
}
func (fc FileController) PostDelete() {
	id := fc.Ctx.PostValue("id")
	fc.Service.Delete(id)
	result := utils.Success()
	fc.Ctx.JSON(result)
}
func (fc FileController) PostRemovedir() {
	id := fc.Ctx.PostValue("id")
	delete(cons.BaseDir, id)
	go service.FlushDictionart(cons.DirFile)
	result := utils.Success()
	fc.Ctx.JSON(result)
}
func (fc FileController) PostAdddir() {
	id := fc.Ctx.PostValue("id")
	key, path := utils.DirpathForId(id)
	_, ok := cons.BaseDir[key]
	result := utils.NewResult()
	if ok {
		result.Fail()
	} else {
		cons.BaseDir[key] = path
		result.Success()
		go service.FlushDictionart(cons.DirFile)
	}
	fc.Ctx.JSON(result)

}

func (fc FileController) PostSync() {
	id := fc.Ctx.PostValue("id")
	curFile := fc.Service.FindOne(id)
	result, newFile := fc.Service.RequestToFile(curFile)
	if result.Code != 200 {
		fc.Ctx.JSON(result)
		return
	}
	result = fc.Service.MoveCut(curFile, newFile)
	fc.Ctx.JSON(result)

}
func (fc FileController) PostMknfo() {
	id := fc.Ctx.PostValue("id")
	curFile := fc.Service.FindOne(id)
	result, newFile := fc.Service.RequestToFile(curFile)
	if result.Code != 200 {
		fc.Ctx.JSON(result)
		return
	}
	newFile.Png = curFile.Png
	newFile.Jpg = curFile.Jpg
	newFile.Nfo = curFile.Nfo
	fc.Service.MakeNfo(newFile)
	result.Success()
	fc.Ctx.JSON(result)

}

func (fc FileController) GetViews() {
	if len(datasource.FileList) == 0 {
		fc.Service.ScanAll()
	}
	keyWord := fc.Ctx.URLParam("keyWord")
	onlyRepeat := fc.Ctx.URLParam("onlyRepeat")
	sortField := fc.Ctx.URLParamDefault("sortField", datasource.DefSortField)
	sortType := fc.Ctx.URLParamDefault("sortType", datasource.DefSortType)
	pageNo, errNo := fc.Ctx.URLParamInt("pageNo")
	if errNo != nil || pageNo == 0 {
		pageNo = 1
	}
	pageSize, errSize := fc.Ctx.URLParamInt("pageSize")
	if errSize != nil || pageSize == 0 {
		pageSize = 55
	}
	page := utils.NewPage()
	page.KeyWord = keyWord
	page.PageNo = pageNo
	page.PageSize = pageSize

	dataSource := datasource.FileList
	if onlyRepeat == "on" {
		keyWord = ""
		dataSource = fc.Service.OnlyRepeat(dataSource)
	}
	page.TotalCnt = len(dataSource)
	page.TotalSize = utils.GetSizeStr(datasource.FileSize)
	datasource.SortMovies(sortField, sortType, false)
	datas := fc.Service.SearchByKeyWord(dataSource, keyWord)
	page.SetResultCnt(len(datas), pageNo)
	page.ResultSize = utils.GetSizeStr(fc.Service.DataSize(datas))
	datas = fc.Service.GetPage(datas, pageNo, pageSize)
	page.CurCnt = len(datas)
	page.CurSize = utils.GetSizeStr(fc.Service.DataSize(datas))
	page.Data = datas

	fc.Ctx.ViewData("playIcon", cons.Play)
	fc.Ctx.ViewData("changeIcon", cons.Change)
	fc.Ctx.ViewData("openIcon", cons.Open)
	fc.Ctx.ViewData("replayIcon", cons.Replay)
	fc.Ctx.ViewData("closeIcon", cons.Close)
	fc.Ctx.ViewData("StopIcon", cons.Stop)

	fc.Ctx.ViewData("sortField", sortField)
	fc.Ctx.ViewData("sortType", sortType)
	fc.Ctx.ViewData("onlyRepeat", onlyRepeat)

	fc.Ctx.ViewData("page", page)
	fc.Ctx.ViewData("curPage", page.PageNo)
	fc.Ctx.ViewData("dirList", cons.BaseDir)
	fc.Ctx.ViewData("title", "文件列表")
	fc.Ctx.View("main.html")
}
func (fc FileController) GetStar() {
	if len(datasource.FileList) == 0 {
		fc.Service.ScanAll()
	}
	var datas []datamodels.Actress
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
