package controller

import (
	"github.com/kataras/iris"
)
import (
	"../../datamodels"
	"../../datasource"
	"../../service"
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
	list := datasource.FileList
	fc.Ctx.Gzip(true)
	fc.Ctx.ViewData("datas", list)
	fc.Ctx.ViewData("title", "文件列表")
	fc.Ctx.View("file_list.html")
}
func (fc FileController) GetPlay() {

}
