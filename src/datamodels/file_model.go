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
	Jpg      string
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
	id = strings.ReplaceAll(id, "\\", "&")
	result := File{
		Id:       id,
		Code:     fileUtils.GetCode(name),
		Name:     name,
		Path:     path,
		Png:      getPng(path, "png"),
		Jpg:      getPng(path, "jpg"),
		Actress:  fileUtils.GetActress(name),
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
	path := f.Png
	if !fileUtils.ExistsFiles(path) {
		path = f.Path
	}
	return utils.ImageToString(path)
}

func getPng(path string, suffix string) string {
	path = strings.ReplaceAll(path, fileUtils.GetSuffux(path), suffix)
	return path
}
func (f File) GetPng() string {
	path := f.Path
	return getPng(path, "png")
}
