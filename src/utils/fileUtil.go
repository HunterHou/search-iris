package utils

import "strings"

func GetSuffux(name string) string {

	var suffix string
	if name == "" {
		return suffix
	}

	nameArr := strings.Split(name, ".")
	suffix = nameArr[len(nameArr)-1]
	return suffix

}
