package model

import (
	"fmt"
	"time"
)

type File struct {
	name     string
	path     string
	code     string
	actress  string
	fileType string
	dirPath  string
	size     int64
	SizeStr  string
	cTime    string
	mTime    string
}

func BuildFile(name string, fileType string, size int64, modTime time.Time) File {
	result := File{}
	result.name = name
	result.fileType = fileType
	result.size = size
	result.SizeStr = getSizeStr(size)
	result.mTime = modTime.Format("2006-01-02 15:04:05")
	return result
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
