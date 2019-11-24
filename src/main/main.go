package main

import (
	"fmt"
	"search/src/service"
)

func main() {
	var baseDir = "F:\\emby\\tomake\\秋山祥子\\inmad\\[秋山祥子] [ATID-349]女教師玩具化計画 秋山祥子"
	items := service.Walk(baseDir)
	for i := 0; i < len(items); i++ {
		fmt.Println(items[i])
	}
}
