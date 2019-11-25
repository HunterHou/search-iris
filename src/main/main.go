package main

import (
	"fmt"

	"../service"
)

func main() {
	var baseDir = "e:\\"
	items := service.Walk(baseDir)
	for i := 0; i < len(items); i++ {
		fmt.Println(items[i])
	}
}
