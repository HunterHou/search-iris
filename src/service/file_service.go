package service

import (
	"../cons"
	"../datamodels"
	"../datasource"
	"../utils/collectionUtils"
	"../utils/fileUtils"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

type FileService struct {
	fileList []datamodels.File
	fileMap  map[string]datamodels.File
}

func (fs FileService) FindOne(path string) datamodels.File {
	if len(datasource.FileLib) == 0 {
		var videoTypes = []string{cons.AVI, cons.MKV, cons.WMV, cons.MP4}
		var queryTypes []string
		queryTypes = collectionUtils.ExtandsItems(queryTypes, videoTypes)
		fs.ScanDisk(cons.BaseDir, queryTypes)
	}
	curFile := datasource.FileLib[path]
	return curFile
}

func (fs FileService) SortItems(lib []datamodels.File) {
	sort.Slice(lib, func(i, j int) bool {
		return lib[i].MTime > lib[j].MTime
	})
}

func (fs FileService) ScanAll() {
	var videoTypes = []string{cons.AVI, cons.MKV, cons.WMV, cons.MP4}
	var queryTypes []string
	queryTypes = collectionUtils.ExtandsItems(queryTypes, videoTypes)
	fs.ScanDisk(cons.BaseDir, queryTypes)
}

func (fs FileService) ScanDisk(baseDir []string, types []string) {
	datasource.FileLib = make(map[string]datamodels.File)
	files := Walks(baseDir, types)
	fs.fileList = files
	fs.fileMap = ArrayToMap(files)
	datasource.FileLib = fs.fileMap
	datasource.FileList = files

}

func (fs FileService) SearchByKeyWord(files []datamodels.File, keyWord string) []datamodels.File {

	if keyWord == "" {
		return files
	}

	var result []datamodels.File
	for _, file := range files {
		if strings.Contains(file.Code, keyWord) {
			result = append(result, file)
		} else if strings.Contains(file.Name, keyWord) {
			result = append(result, file)
		} else if strings.Contains(file.Actress, keyWord) {
			result = append(result, file)
		}
	}

	return result
}

func ArrayToMap(files []datamodels.File) map[string]datamodels.File {
	filemap := make(map[string]datamodels.File)
	for i := 0; i < len(files); i++ {
		curFile := files[i]
		filemap[curFile.Path] = curFile
	}
	return filemap
}
func Walks(baseDir []string, types []string) []datamodels.File {

	var wg sync.WaitGroup
	var datas = make(chan []datamodels.File, 10000)
	var result []datamodels.File
	wg.Add(len(baseDir))
	for i := 0; i < len(baseDir); i++ {
		go goWalk(baseDir[i], types, &wg, datas)
	}
	wg.Wait()
	close(datas)
	for {
		data, ok := <-datas
		if !ok {
			break
		}
		result = Expands(result, data)
	}
	return result

}
func goWalk(baseDir string, types []string, wg *sync.WaitGroup, datas chan []datamodels.File) {
	fmt.Println(baseDir)
	defer wg.Done()
	files := Walk(baseDir, types)
	datas <- files
}

//遍历目录 获取文件库
func Walk(baseDir string, types []string) []datamodels.File {
	var result []datamodels.File
	files, _ := ioutil.ReadDir(baseDir)
	for _, path := range files {

		pathAbs := filepath.Join(baseDir, path.Name())
		if path.IsDir() {
			childResult := Walk(pathAbs, types)
			result = Expands(result, childResult)
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

func Expands(originArr []datamodels.File, insertArr []datamodels.File) []datamodels.File {
	if len(insertArr) == 0 {
		return originArr
	}

	for i := 0; i < len(insertArr); i++ {
		originArr = append(originArr, insertArr[i])
	}
	return originArr
}
