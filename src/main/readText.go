package main

import (
	"../datasource"
	"../service"
	"fmt"
)

func main() {
	service.ReadDictionary("D:\\code\\search-iris\\src\\dirList.ini")
	for index, name := range datasource.DictLib.LibMap {
		if len(name) <= 1 {
			fmt.Println(index, ":", name)
		} else {
			for _, s := range name {
				fmt.Println(index, ":", s)
			}

		}

	}
}
