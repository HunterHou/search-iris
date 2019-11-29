package controller

import "github.com/kataras/iris"

import (
	"../../cons"
	"../../datamodels"
	"../../service"
	"../../utils/collectionUtils"
)

type FileController struct {
	Ctx     iris.Context
	Service service.FileService
}

func (fc FileController) GetAll() []datamodels.File {
	var baseDir = "e:\\"
	var videoTypes = []string{cons.AVI, cons.MKV, cons.WMV, cons.MP4}
	var queryTypes []string
	queryTypes = collectionUtils.ExtandsItems(queryTypes, videoTypes)
	return fc.Service.ScanAll(baseDir, queryTypes)
}
func (fc FileController) GetSiew() {
	var baseDir = "e:\\"
	var videoTypes = []string{cons.AVI, cons.MKV, cons.WMV, cons.MP4}
	var queryTypes []string
	queryTypes = collectionUtils.ExtandsItems(queryTypes, videoTypes)
	list := fc.Service.ScanAll(baseDir, queryTypes)
	fc.Ctx.Gzip(true)
	fc.Ctx.ViewData("datas", list)
	fc.Ctx.ViewData("title", "文件列表")
	fc.Ctx.View("file_list.html")
}
