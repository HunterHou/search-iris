package utils

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func DirpathForId(path string) (string, string) {
	id, _ := url.QueryUnescape(path)
	id = strings.ReplaceAll(id, "\\\\", "\\")
	id = strings.ReplaceAll(id, "\\", "~")
	newpath := strings.ReplaceAll(id, "~", "\\")
	return id, newpath
}

func GetPng(path string, suffix string) string {
	path = strings.ReplaceAll(path, GetSuffux(path), suffix)
	return path
}

func ExistsFiles(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func GetSuffux(fielname string) string {

	var suffix string
	if fielname == "" {
		return suffix
	}
	suffix = filepath.Ext(fielname)
	suffix = strings.ToLower(suffix)
	if strings.Contains(suffix, ".") {
		suffix = strings.TrimPrefix(suffix, ".")
	}
	return suffix

}

// 获取文件名
func GetTitle(filename string) string {
	result := ""
	if filename == "" {
		return result
	}
	last_suffix := path.Ext(filename)
	filename = strings.TrimSuffix(filename, last_suffix)
	return filename

}

// 根据 文件名称  分析番号 [] 中包含 '-'符号...
func GetActress(fileName string) string {
	code := ""
	rights := strings.Split(fileName, "[")
	if len(rights) <= 1 {
		return GetTitle(fileName)
	}
	for index, value := range rights {
		if index == 0 {
			continue
		}
		right := value
		lefts := strings.Split(right, "]")
		for _, left := range lefts {
			if !strings.Contains(left, "-") {
				return left
			}
		}
	}
	return code
}

// 根据 文件名称  分析番号 [] 中包含 '-'符号...
func GetCode(fileName string) string {
	code := ""
	rights := strings.Split(fileName, "[")
	if len(rights) <= 1 {
		return GetTitle(fileName)
	}
	for index, value := range rights {
		if index == 0 {
			continue
		}
		right := value
		lefts := strings.Split(right, "]")
		for _, left := range lefts {
			if strings.Contains(left, "-") || strings.Contains(left, "_") {
				return left
			}else{
				code = left
			}
		}
	}
	if strings.Contains(code,".mp4"){
		code=strings.ReplaceAll(code,".mp4","")
	}
	return code
}

func GetSizeStr(fSize int64) string {

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
