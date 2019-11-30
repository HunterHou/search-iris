package controller

import (
	"github.com/kataras/iris"
)
import (
	"../../cons"
	"../../datasource"
	"../../service"
	"../../utils"
	"../../utils/collectionUtils"
)

type TestController struct {
	Ctx     iris.Context
	Service service.FileService
}

func (ts TestController) GetHello() {
	name := ts.Ctx.URLParam("name")
	ts.Ctx.HTML("<h1> hello " + name + "</h1>")
}

func (ts TestController) GetScan() {
	var videoTypes = []string{cons.AVI, cons.MKV, cons.WMV, cons.MP4}
	var queryTypes []string
	queryTypes = collectionUtils.ExtandsItems(queryTypes, videoTypes)
	fileService := service.FileService{}
	fileService.ScanDisk(cons.BaseDir, queryTypes)
	ts.Ctx.JSON(utils.Success())
}

func (ts TestController) GetResult() {
	result := utils.Result{}
	result.Success()
	result.Data = datasource.FileLib
	ts.Ctx.JSON(result)
}
