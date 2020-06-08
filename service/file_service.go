package service

import (
	"../cons"
	"../datamodels"
	"../datasource"
	"../utils"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

type FileService struct {
}

func (fs FileService) MoveCut(srcFile datamodels.Movie, toFile datamodels.Movie) utils.Result {
	result := utils.Result{}
	root := srcFile.DirPath
	path := root + "\\" + toFile.Actress
	if toFile.Studio != "" {
		path = path + "\\" + toFile.Studio
	}
	title := toFile.Title
	title = strings.ReplaceAll(title, ":", "~")
	title = strings.ReplaceAll(title, ".", "~")

	dirname := "[" + toFile.Actress + "][" + toFile.Code + "]" + title
	dirpath := path + "\\" + dirname
	os.MkdirAll(dirpath, os.ModePerm)
	filename := dirname + "." + utils.GetSuffux(srcFile.Path)
	finalpath := dirpath + "\\" + filename
	jpgpath := utils.GetPng(finalpath, "jpg")
	nfopath := utils.GetPng(finalpath, "nfo")
	jpgOut, createErr := os.Create(jpgpath)
	if createErr != nil {
		//TODO 创建失败  标题 特殊字符处理 改为 演员+番号
		dirname = "[" + toFile.Actress + "][" + toFile.Code + "]"
		dirpath = path + "\\" + dirname
		os.MkdirAll(dirpath, os.ModePerm)
		filename = dirname + "." + utils.GetSuffux(srcFile.Path)
		finalpath = dirpath + "\\" + filename
		jpgpath = utils.GetPng(finalpath, "jpg")
		jpgOut, createErr = os.Create(jpgpath)
		if createErr != nil {
			result.Fail()
			fmt.Println("createErr:", createErr)
			os.Rename(finalpath, srcFile.Path)
			result.Message = "文件创建失败：" + jpgpath
			return result
		}
	}
	resp, downErr := http.Get(toFile.Jpg)
	if downErr != nil {
		result.Fail()
		fmt.Println("downErr:", downErr)
		os.Rename(finalpath, srcFile.Path)
		result.Message = "文件下载失败：" + toFile.Jpg
		return result
	}
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		result.Fail()
		fmt.Println("readErr:", readErr)
		os.Rename(finalpath, srcFile.Path)
		result.Message = "请求读取response失败"
		return result
	}
	jpgOut.Write(body)
	jpgOut.Close()
	pngErr := utils.ImageToPng(jpgpath)
	if pngErr != nil {
		result.Fail()
		fmt.Println("pngErr:", pngErr)
		os.Rename(finalpath, srcFile.Path)
		result.Message = "png生成失败"
		return result
	}
	os.Rename(srcFile.Path, finalpath)
	toFile.Jpg = jpgpath
	toFile.Nfo = nfopath
	toFile.Png = utils.GetPng(finalpath, "png")
	fs.MakeNfo(toFile)
	result.Success()
	result.Message = "【" + dirname + "】" + result.Message
	return result
}

func (fs FileService) MakeNfo(toFile datamodels.Movie) {
	nfo, _ := os.Create(toFile.Nfo)
	defer nfo.Close()
	nfoStr := "<?xml version=\"1.0\" encoding=\"utf-8\" standalone=\"yes\"?> \n"
	nfoStr += "<movie>\n"
	nfoStr += "<year>" + toFile.PTime + "</year>\n"
	nfoStr += "<title>" + toFile.Title + "</title>\n"
	nfoStr += "<releasedate>" + toFile.PTime + "</releasedate>\n"
	nfoStr += "<runtime>" + toFile.PTime + "</runtime>\n"
	nfoStr += "<poster>" + toFile.Jpg + "</poster>\n"
	nfoStr += "<thumb>" + toFile.Jpg + "</thumb>\n"
	nfoStr += "<fanart>" + toFile.Jpg + "</fanart>\n"
	nfoStr += "<maker>" + toFile.Supplier + "</maker>\n"
	nfoStr += "<studio>" + toFile.Studio + "</studio>\n"
	nfoStr += "<num>" + toFile.Code + "</num>\n"
	nfoStr += "<release>" + toFile.PTime + "</release>\n"
	nfoStr += "<cover>" + toFile.Jpg + ".jpg" + "</cover>\n"
	nfoStr += "<art>"
	nfoStr += "<poster>" + toFile.Png + "</poster>\n"
	nfoStr += "</art>"
	nfoStr += "<actor>"
	nfoStr += "<name>" + toFile.Actress + "</name>\n"
	nfoStr += "<type>Actor</type>\n"
	nfoStr += "</actor>\n"
	nfoStr += "<year>" + toFile.PTime + "</year>\n"
	nfoStr += "</movie>\n"
	nfo.WriteString(nfoStr)
}

func httpGet(url string) (*http.Response, error) {

	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("User-Agent", "Mozilla/6.0")
	client := &http.Client{}
	resp, err := client.Do(request)
	return resp, err

}

func (fs FileService) RequestToFile(srcFile datamodels.Movie) (utils.Result, datamodels.Movie) {

	result := utils.Result{}
	newFile := datamodels.Movie{}
	if srcFile.Code == "" {
		result.Fail()
		return result, newFile
	}
	url := cons.BaseUrl + srcFile.Code
	resp, err := httpGet(url)
	if err != nil {
		fmt.Println("err", err)
		result.Fail()
		return result, newFile
	}
	defer resp.Body.Close()
	if 200 != resp.StatusCode {
		if strings.Contains(url, "_") {
			url = strings.ReplaceAll(url, "_", "-")
		} else if strings.Contains(url, "-") {
			url = strings.ReplaceAll(url, "-", "_")
		}
		resp, err = httpGet(url)
		if 200 != resp.StatusCode {
			fmt.Println("status error:", resp.StatusCode, resp.Status)
			result.Fail()
			result.Message = "请求失败：" + resp.Status + " url:" + url
			return result, newFile
		}

	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		result.Fail()
		result.Message = "html解析失败"
		fmt.Println("err:", err)
	}
	bigImage := doc.Find(".bigImage img")

	newFile.Id = srcFile.Id
	newFile.Title = bigImage.AttrOr("title", "")
	newFile.Jpg = bigImage.AttrOr("src", "")
	info := doc.Find(".header")
	info.Each(func(i int, selection *goquery.Selection) {
		item := selection.Text()
		if strings.HasPrefix(item, "發行日期:") {
			newFile.PTime = selection.Parent().Text()
			newFile.PTime = strings.Replace(newFile.PTime, "發行日期:", "", 1)
		} else if strings.HasPrefix(item, "長度:") {
			newFile.Length = selection.Parent().Text()
			newFile.Length = strings.Replace(newFile.Length, "長度:", "", 1)
		} else if strings.HasPrefix(item, "演員") {
			stars := doc.Find(".star-name")
			stars.Each(func(i int, selection *goquery.Selection) {
				starName := selection.Text()
				newFile.Actress += strings.TrimSpace(starName)
			})
		} else if strings.HasPrefix(item, "導演:") {
			newFile.Director = selection.Next().Text()
		} else if strings.HasPrefix(item, "製作商:") {
			newFile.Supplier = selection.Next().Text()
		} else if strings.HasPrefix(item, "發行商:") {
			newFile.Studio = selection.Next().Text()
		} else if strings.HasPrefix(item, "系列:") {
			newFile.Series = selection.Next().Text()
		} else if strings.HasPrefix(item, "識別碼:") {
			newFile.Code = selection.Next().Text()
		}
	})
	result.Success()
	result.Data = newFile
	return result, newFile
}

func (fs FileService) FindOne(Id string) datamodels.Movie {
	if len(datasource.FileLib) == 0 {
		fs.ScanAll()
	}
	curFile := datasource.FileLib[Id]
	return curFile
}

func (fs FileService) SortAct(lib []datamodels.Actress) {
	sort.Slice(lib, func(i, j int) bool {
		return lib[i].Cnt > lib[j].Cnt
	})
}

func (fs FileService) ScanAll() {
	dirList := []string{}
	for _, v := range cons.BaseDir {
		dirList = append(dirList, v)
	}
	cons.QueryTypes = []string{}
	cons.QueryTypes = utils.ExtandsItems(cons.QueryTypes, cons.VideoTypes)
	cons.QueryTypes = utils.ExtandsItems(cons.QueryTypes, cons.Docs)
	cons.QueryTypes = utils.ExtandsItems(cons.QueryTypes, cons.Images)

	fs.ScanDisk(dirList, cons.QueryTypes)
}
func (fs FileService) Delete(id string) {
	file := fs.FindOne(id)
	list := []string{file.Path, file.Png, file.Jpg, file.Nfo}
	for i := 0; i < len(list); i++ {
		err := os.Remove(list[i])
		if err != nil {
			fmt.Println(err)
		}
	}
	//TODO 删除父文件夹
	//dirname := path.Dir(file.Path)
	//fmt.Println(dirname)
	//deleteDir(dirname)

}

func deleteDir(filename string) {
	dirname, _ := path.Split(filename)
	files2, _ := ioutil.ReadDir(dirname)
	if len(files2) == 0 {
		os.Remove(dirname)
		return
	}
	dirname2, _ := path.Split(dirname)
	deleteDir(dirname2)
}

func (fs FileService) ScanDisk(baseDir []string, types []string) {
	datasource.FileLib = make(map[string]datamodels.Movie)
	files := Walks(baseDir, types)
	fileMap, actressMap, supplierMap, fileSize := ArrayToMap(files)
	var newFiles []datamodels.Movie
	for _, item := range fileMap {
		newFiles = append(newFiles, item)
	}
	datasource.FileLib = fileMap
	datasource.FileList = newFiles
	datasource.ActressLib = actressMap
	datasource.SupplierLib = supplierMap
	datasource.FileSize = fileSize

}

func (fs FileService) OnlyRepeat(files []datamodels.Movie) []datamodels.Movie {
	var result []datamodels.Movie
	codeMap := make(map[string]datamodels.Movie)
	for _, movie := range files {
		if movie.Code == "" {
			continue
		}
		ele, ok := codeMap[movie.Code]
		if ok {
			result = append(result, ele)
			result = append(result, movie)
			continue
		} else {
			codeMap[movie.Code] = movie
		}

	}
	return result
}

func (fs FileService) SearchByKeyWord(files []datamodels.Movie, keyWord string) []datamodels.Movie {

	if keyWord == "" {
		return files
	}

	var result []datamodels.Movie
	var size int64
	for _, file := range files {
		if strings.Contains(strings.ToUpper(file.Code), strings.ToUpper(keyWord)) {
			result = append(result, file)
			size = size + file.Size
		} else if strings.Contains(strings.ToUpper(file.Name), strings.ToUpper(keyWord)) {
			result = append(result, file)
			size = size + file.Size
		} else if strings.Contains(strings.ToUpper(file.Actress), strings.ToUpper(keyWord)) {
			result = append(result, file)
			size = size + file.Size
		}
	}

	return result
}

func (fs FileService) GetPage(files []datamodels.Movie, pageNo int, pageSize int) []datamodels.Movie {

	if len(files) == 0 {
		return files
	}
	size := len(files)
	start := (pageNo - 1) * pageSize

	end := size
	if size-start > pageSize {
		end = start + pageSize
	}
	if len(files) < pageSize {
		return files
	}
	if start > end {
		start = 0
	}
	data := files[start:end]
	return data
}

func (fs FileService) DataSize(data []datamodels.Movie) int64 {
	var dataSize int64
	for _, d := range data {
		dataSize = dataSize + d.Size
	}
	return dataSize
}

func ArrayToMap(files []datamodels.Movie) (map[string]datamodels.Movie, map[string]datamodels.Actress, map[string]datamodels.Supplier, int64) {
	filemap := make(map[string]datamodels.Movie)
	actessmap := make(map[string]datamodels.Actress)
	suppliermap := make(map[string]datamodels.Supplier)
	var size int64
	for i := 0; i < len(files); i++ {
		curFile := files[i]
		size = size + curFile.Size
		filemap[curFile.Id] = curFile
		curActress, ok := actessmap[curFile.Actress]
		if ok {
			curActress.PlusCnt()
			curActress.PlusSize(curFile.Size)
			actessmap[curFile.Actress] = curActress
		} else {
			actessmap[curFile.Actress] = datamodels.NewActres(curFile.Actress, curFile.Png, curFile.Size)
		}
		curSupplier, okS := suppliermap[curFile.Supplier]
		if okS {
			curSupplier.Plus()
			suppliermap[curFile.Supplier] = curSupplier
		} else {
			suppliermap[curFile.Supplier] = datamodels.NewSupplier(curFile.Supplier)
		}

	}
	return filemap, actessmap, suppliermap, size
}
func Walks(baseDir []string, types []string) []datamodels.Movie {

	var wg sync.WaitGroup
	var dataMovie = make(chan []datamodels.Movie, 10000)
	var result []datamodels.Movie
	wg.Add(len(baseDir))
	for i := 0; i < len(baseDir); i++ {
		go goWalk(baseDir[i], types, &wg, dataMovie)
	}
	wg.Wait()
	close(dataMovie)
	for {
		data, ok := <-dataMovie
		if !ok {
			break
		}
		result = ExpandsMovie(result, data)
	}

	return result

}
func goWalk(baseDir string, types []string, wg *sync.WaitGroup, datas chan []datamodels.Movie) {
	defer wg.Done()
	files := Walk(baseDir, types)
	datas <- files
}

//遍历目录 获取文件库
func Walk(baseDir string, types []string) []datamodels.Movie {
	var result []datamodels.Movie
	files, _ := ioutil.ReadDir(baseDir)
	for _, path := range files {

		pathAbs := filepath.Join(baseDir, path.Name())
		if path.IsDir() {
			childResult := Walk(pathAbs, types)
			result = ExpandsMovie(result, childResult)
		} else {
			name := path.Name()
			suffix := utils.GetSuffux(name)
			if utils.HasItem(types, suffix) {
				file := datamodels.NewFile(baseDir, pathAbs, name, suffix, path.Size(), path.ModTime())
				result = append(result, file)
			}

		}
	}
	return result
}

func ExpandsMovie(originArr []datamodels.Movie, insertArr []datamodels.Movie) []datamodels.Movie {
	if len(insertArr) == 0 {
		return originArr
	}

	for i := 0; i < len(insertArr); i++ {
		originArr = append(originArr, insertArr[i])
	}
	return originArr
}
func ExpandsActess(originArr []datamodels.Actress, insertArr []datamodels.Actress) []datamodels.Actress {
	if len(insertArr) == 0 {
		return originArr
	}

	for i := 0; i < len(insertArr); i++ {
		originArr = append(originArr, insertArr[i])
	}
	return originArr
}
func ExpandsSupplier(originArr []datamodels.Supplier, insertArr []datamodels.Supplier) []datamodels.Supplier {
	if len(insertArr) == 0 {
		return originArr
	}

	for i := 0; i < len(insertArr); i++ {
		originArr = append(originArr, insertArr[i])
	}
	return originArr
}
