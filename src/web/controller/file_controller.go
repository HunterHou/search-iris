package controller

import (
	"github.com/kataras/iris"
	"io/ioutil"
)
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
	var videoTypes = []string{cons.AVI, cons.MKV, cons.WMV, cons.MP4}
	var queryTypes []string
	queryTypes = collectionUtils.ExtandsItems(queryTypes, videoTypes)
	return fc.Service.ScanAll(cons.BaseDir, queryTypes)
}
func (fc FileController) GetViews() {
	var videoTypes = []string{cons.AVI, cons.MKV, cons.WMV, cons.MP4}
	var queryTypes []string
	queryTypes = collectionUtils.ExtandsItems(queryTypes, videoTypes)
	list := fc.Service.ScanAll(cons.BaseDir, queryTypes)
	fc.Ctx.Gzip(true)
	fc.Ctx.ViewData("datas", list)
	fc.Ctx.ViewData("title", "文件列表")
	fc.Ctx.View("file_list.html")
}
func (fc FileController) GetPlay() {
	url := fc.Ctx.URLParam("url")
	file := fc.Service.FindOne(url)
	data, _ := ioutil.ReadFile(string(file.Path))
	fc.Ctx.Header("Content-Type", "video/mp4")
	fc.Ctx.Write(data)
}
