package fileUtils

import (
	"path/filepath"
	"strings"
)

func GetSuffux(fielname string) string {

	var suffix string
	if fielname == "" {
		return suffix
	}
	suffix = filepath.Ext(fielname)
	suffix = strings.ToLower(suffix)
	if strings.Contains(suffix, ".") {
		suffix = strings.TrimPrefix(suffix, ".")
	}
	return suffix

}
