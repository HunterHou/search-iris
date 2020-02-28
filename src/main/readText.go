package main

import (
	"../utils"
	"fmt"
)

func main() {
	dirs := utils.ReadDir("D:\\code\\search-iris\\src\\dirList.txt")
	for index, name := range dirs {
		fmt.Println(index, ":", name)
	}
}
