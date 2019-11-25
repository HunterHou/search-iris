package model

import (
	"fmt"
	"time"

	"../utils/fileUtils"
)

//声明一个File结构体 表示一个文件信息
type File struct {
	code     string
	name     string
	path     string
	actress  string
	fileType string
	dirPath  string
	size     int64
	sizeStr  string
	cTime    string
	mTime    string
}

//使用工厂模式 返回一个 File 实例
func NewFile(dir string, path string, name string, fileType string, size int64, modTime time.Time) File {
	result := File{
		code:     fileUtils.GetCode(name),
		name:     name,
		path:     path,
		actress:  fileUtils.GetCode(name),
		fileType: fileType,
		dirPath:  dir,
		size:     size,
		sizeStr:  fileUtils.GetSizeStr(size),
		cTime:    "",
		mTime:    modTime.Format("2006-01-02 15:04:05"),
	}
	return result
}

func (this File) GetFileInfo() string {
	info := fmt.Sprintf("name: %v\t code:%v\t fileType:%v\t sizeStr:%v\t actress:%v\t path:%v\t",
		this.name, this.code, this.fileType, this.sizeStr, this.actress, this.path)
	return info
}
