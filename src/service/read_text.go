package service

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
import (
	"../cons"
	"../datamodels"
)

func FlushDictionart(path string) {
	dict := datamodels.NewDictionary()
	dict.PutProperty("BaseUrl", cons.BaseUrl)
	for _, dir := range cons.BaseDir {
		dict.PutProperty("dir", dir)
	}
	for _, image := range cons.Images {
		dict.PutProperty("Images", image)
	}
	WriteDictionary(path, dict)

}

func WriteDictionary(path string, dict datamodels.Dictionary) {
	outStream, openErr := os.OpenFile(path, os.O_TRUNC|os.O_RDWR, os.ModePerm)
	defer outStream.Close()
	if openErr != nil {
		fmt.Println("openErr", openErr)
	}
	writer := bufio.NewWriter(outStream)
	for key, value := range dict.LibMap {
		for _, v := range value {
			writer.WriteString(key + "=" + v + "\n")
		}
	}
	writer.Flush()
}

func ReadDictionary(path string) datamodels.Dictionary {
	outStream, openErr := os.Open(path)
	defer outStream.Close()
	if openErr != nil {
		fmt.Println("openErr", openErr)
	}
	reader := bufio.NewReader(outStream)
	dict := datamodels.NewDictionary()
	for {
		lineStr, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		lineStr = strings.TrimSpace(lineStr)
		if lineStr == "" {
			continue
		}
		line := strings.Split(lineStr, "=")
		dict.PutProperty(line[0], line[1])
	}
	return dict
}
