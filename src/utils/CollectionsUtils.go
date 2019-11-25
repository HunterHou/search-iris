package utils

//判断集合是否包含
func HasItem(lib []string, item string) bool {
	if lib == nil {
		return false
	}
	if len(lib) == 0 {
		return false
	}
	for i := 0; i < len(lib); i++ {
		if item == lib[i] {
			return true
		}
	}
	return false

}
