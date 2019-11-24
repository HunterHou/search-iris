package main

import (
	"fmt"
	"search/src/service"
)

func main() {
	var baseDir = "F:\\emby"
	items := service.Walk(baseDir)
	for i := 0; i < len(items); i++ {
		fmt.Println(items[i])
	}
}
