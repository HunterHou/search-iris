package datamodels

type Dictionary struct {
	key    string
	values []string
}

func New() Dictionary {
	return Dictionary{}
}

func (dict *Dictionary) SetProperty(key string, value string) {

}
