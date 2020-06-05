package datamodels

type Dictionary struct {
	LibMap map[string][]string
	LibArr []string
}

func NewDictionary() Dictionary {
	return Dictionary{
		make(map[string][]string), nil,
	}
}

func (dict *Dictionary) PutProperty(key string, value string) {
	var values []string
	if dict.LibMap[key] != nil {
		for _, s := range dict.LibMap[key] {
			values = append(values, s)
		}
	} else {
	}
	values = append(values, value)
	dict.LibMap[key] = values
}
func (dict *Dictionary) GetProperty(key string) []string {
	return dict.LibMap[key]
}
