package model

import (
	"fmt"
	"time"
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
		code:     "",
		name:     name,
		path:     path,
		actress:  "",
		fileType: fileType,
		dirPath:  dir,
		size:     size,
		sizeStr:  getSizeStr(size),
		cTime:    "",
		mTime:    modTime.Format("2006-01-02 15:04:05"),
	}
	return result
}

func (this File) GetFileInfo() string {
	info := fmt.Sprintf("%v\t %v\t %v\t %v\t %v\t %v\t",
		this.name, this.code, this.fileType, this.sizeStr, this.actress, this.path)
	return info
}

func getSizeStr(fSize int64) string {

	fileSize := float64(fSize)
	result := ""
	if fileSize <= 1024 {
		result = fmt.Sprintf("%.f", fileSize)
	} else if fileSize <= 1024*1024 {
		size := float64(fileSize / 1024)
		result = fmt.Sprintf("%.f", size) + " k"
	} else if fileSize <= 1024*1024*1024 {
		size := float64(fileSize / (1024 * 1024))
		result = fmt.Sprintf("%.2f", size) + " M"
	} else if fileSize <= 1024*1024*1024*1024 {
		size := float64(fileSize / (1024 * 1024 * 1024))
		result = fmt.Sprintf("%.2f", size) + " G"
	} else if fileSize <= 1024*1024*1024*1024*1024 {
		size := float64(fileSize / (1024 * 1024 * 1024 * 1024))
		result = fmt.Sprintf("%.2f", size) + " T"
	} else {
		size := float64(fileSize / (1024 * 1024 * 1024 * 1024))
		result = fmt.Sprintf("%.2f", size) + " T"
	}
	return result
}
