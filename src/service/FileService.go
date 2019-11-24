package service

import (
	"io/ioutil"
	"search/src/cons"
	"search/src/model"
	"search/src/utils"
)

var videoTypes = []string{cons.AVI, cons.MKV, cons.WMV, cons.MP4}

func Walk(baseDir string) []model.File {
	var result []model.File
	files, _ := ioutil.ReadDir(baseDir)
	for _, filePath := range files {
		if filePath.IsDir() {
			continue
		} else {
			name := filePath.Name()
			suffix := utils.GetSuffux(name)
			if HasItem(videoTypes, suffix) {
				file := model.BuildFile(name, suffix, filePath.Size(), filePath.ModTime())
				result = append(result, file)
			}

		}
	}
	return result
}

func HasItem(lib []string, item string) bool {
	if lib == nil {
		return false
	}
	if len(lib) == 0 {
		return false
	}
	for i := 0; i < len(lib); i++ {
		if item == lib[i] {
			return true
		}
	}
	return false

}
