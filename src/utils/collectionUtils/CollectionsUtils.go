package collectionUtils

import "strings"

//判断集合是否包含
func HasItem(lib []string, item string) bool {
	if lib == nil {
		return false
	}
	if len(lib) == 0 {
		return false
	}
	for i := 0; i < len(lib); i++ {

		if strings.Compare(item, lib[i]) == 0 {
			return true
		}
	}
	return false
}

func ExtandsItems(lib []string, items []string) []string {
	if len(items) == 0 || items == nil {
		return lib
	}
	for i := 0; i < len(items); i++ {
		lib = append(lib, items[i])
	}
	return lib

}

func IndexOf(lib []string, item string) int {
	if lib == nil {
		return -1
	}
	if len(lib) == 0 {
		return -1
	}
	for i := 0; i < len(lib); i++ {

		if strings.Compare(item, lib[i]) == 0 {
			return i
		}
	}
	return -1
}

func RemoveItem(lib []string, item string) []string {
	index := IndexOf(lib, item)
	if index != -1 {
		return ExtandsItems(lib[0:index], lib[(index+1):len(lib)])
	}
	return lib
}
