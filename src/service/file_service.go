package service

import (
	"io/ioutil"
	"path/filepath"

	"../datamodels"
	"../datasource"
	"../utils/collectionUtils"
	"../utils/fileUtils"
)

type FileService struct {
	fileList []datamodels.File
	fileMap  map[string]datamodels.File
}

func (fs FileService) ScanDisk(baseDir string, types []string) {
	datasource.FileLib = make(map[string]datamodels.File)
	files := Walk(baseDir, types)
	fs.fileList = files
	fs.fileMap = ArrayToMap(files)
	datasource.FileLib = fs.fileMap

}
func (fs FileService) ScanAll(baseDir string, types []string) []datamodels.File {
	files := Walk(baseDir, types)
	fs.fileList = files
	return files
}

func ArrayToMap(files []datamodels.File) map[string]datamodels.File {
	filemap := make(map[string]datamodels.File)
	for i := 0; i < len(files); i++ {
		curFile := files[i]
		filemap[curFile.Path] = curFile
	}
	return filemap
}

//遍历目录 获取文件库
func Walk(baseDir string, types []string) []datamodels.File {
	var result []datamodels.File
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
				file := datamodels.NewFile(baseDir, pathAbs, name, suffix, path.Size(), path.ModTime())
				result = append(result, file)
			}

		}
	}
	return result
}

func expands(originArr []datamodels.File, insertArr []datamodels.File) []datamodels.File {
	if len(insertArr) == 0 {
		return originArr
	}

	for i := 0; i < len(insertArr); i++ {
		originArr = append(originArr, insertArr[i])
		i++
	}
	return originArr
}
