package main

import (
	"../cons"
	"../service"
	"../utils/collectionUtils"
	"fmt"
)

func main() {
	var baseDir = "e:\\"
	var videoTypes = []string{cons.AVI, cons.MKV, cons.WMV, cons.MP4}
	var imageTypes = []string{cons.JPG, cons.PNG, cons.GIF}
	var queryTypes []string
	queryTypes = collectionUtils.ExtandsItems(videoTypes, imageTypes)
	items := service.Walk(baseDir, queryTypes)
	i := 0
	for {
		if len(items) == i {
			break
		}
		fmt.Println(items[i])
		i++
	}
}
