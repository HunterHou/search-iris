package service

import (
	"io/ioutil"
	"path/filepath"

	"../model"
	"../utils/collectionUtils"
	"../utils/fileUtils"
)

//遍历目录 获取文件库
func Walk(baseDir string, types []string) []model.File {
	var result []model.File
	files, _ := ioutil.ReadDir(baseDir)
	for _, path := range files {

		pathAbs := filepath.Join(baseDir, path.Name())
		if path.IsDir() {
			childResult := Walk(pathAbs, types)
			result = expands(result, childResult)
		} else {
			name := path.Name()
			suffix := fileUtils.GetSuffux(name)
			if collectionUtils.HasItem(types, suffix) {
				file := model.NewFile(baseDir, pathAbs, name, suffix, path.Size(), path.ModTime())
				result = append(result, file)
			}

		}
	}
	return result
}

func expands(originArr []model.File, insertArr []model.File) []model.File {
	if len(insertArr) == 0 {
		return originArr
	}

	for i := 0; i < len(insertArr); i++ {
		originArr = append(originArr, insertArr[i])
		i++
	}
	return originArr
}
