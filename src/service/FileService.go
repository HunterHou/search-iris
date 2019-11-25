package service

import (
	"filepath"
	"io/ioutil"

	"../cons"
	"../model"
	"../utils"
)

var videoTypes = []string{cons.AVI, cons.MKV, cons.WMV, cons.MP4}

//遍历目录 获取文件库
func Walk(baseDir string) []model.File {
	var result []model.File
	files, _ := ioutil.ReadDir(baseDir)
	for _, path := range files {
		if path.IsDir() {
			subDir := filepath.Join(baseDir, path)
			childResult := Walk(subDir)
		} else {
			name := path.Name()
			suffix := utils.GetSuffux(name)
			if utils.HasItem(videoTypes, suffix) {
				file := model.NewFile(name, suffix, path.Size(), path.ModTime())
				result = append(result, file)
			}

		}
	}
	return result
}
