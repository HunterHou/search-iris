package utils

import (
	"encoding/base64"
	"os"
)

func ImageToString(path string) string {
	file, _ := os.Open(path)
	defer file.Close()
	sourceBuffer := make([]byte, 500000)
	n, _ := file.Read(sourceBuffer)
	return base64.StdEncoding.EncodeToString(sourceBuffer[:n])
}
