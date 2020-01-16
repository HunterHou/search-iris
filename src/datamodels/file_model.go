package datamodels

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"../utils"
	"../utils/fileUtils"
)

//声明一个File结构体 表示一个文件信息
type File struct {
	Id       string
	Code     string "json:'code'"
	Name     string
	Path     string
	Png      string
	Actress  string
	FileType string
	DirPath  string
	Size     int64
	SizeStr  string
	CTime    string
	MTime    string
}

func NewFile(dir string, path string, name string, fileType string, size int64, modTime time.Time) File {
	// 使用工厂模式 返回一个 File 实例
	id, _ := url.QueryUnescape(path)
	id= strings.ReplaceAll(id,"\\","&")
	result := File{
		Id:       id,
		Code:     fileUtils.GetCode(name),
		Name:     name,
		Path:     path,
		Png:      getPng(path),
		Actress:  fileUtils.GetCode(name),
		FileType: fileType,
		DirPath:  dir,
		Size:     size,
		SizeStr:  fileUtils.GetSizeStr(size),
		CTime:    "",
		MTime:    modTime.Format("2006-01-02 15:04:05"),
	}
	return result
}

func (f File) GetFileInfo() string {
	//
	info := fmt.Sprintf("name: %v\t code:%v\t fileType:%v\t sizeStr:%v\t actress:%v\t path:%v\t",
		f.Name, f.Code, f.FileType, f.SizeStr, f.Actress, f.Path)
	return info
}
func (f File) PngBase64() string {
	return utils.ImageToString(f.Png)
}

func encode() {

}

func getPng(path string) string {
	path = strings.ReplaceAll(path, fileUtils.GetSuffux(path), "png")
	return path
}
func (f File) GetPng() string {
	//
	path := f.Path
	return getPng(path)
}
