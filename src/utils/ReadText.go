package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func WriteDir(path string, dirMap map[string]string) {
	outStream, openErr := os.OpenFile(path, os.O_RDWR, os.ModePerm)
	defer outStream.Close()
	if openErr != nil {
		fmt.Println("openErr", openErr)
	}
	writer := bufio.NewWriter(outStream)
	for _, v := range dirMap {
		writer.WriteString("dir=" + v + "\n")
	}
	writer.Flush()
}

func ReadDir(path string) []string {
	outStream, openErr := os.Open(path)
	defer outStream.Close()
	if openErr != nil {
		fmt.Println("openErr", openErr)
	}
	reader := bufio.NewReader(outStream)
	var dirs []string
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
		if line[0] == "dir" {
			dirs = append(dirs, line[1])
		}
	}
	return dirs
}
