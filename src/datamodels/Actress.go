package datamodels

import "../utils"

type Actress struct {
	Name string
	Url  string
	Cnt  int
}

func NewActres(name string, url string) Actress {
	return Actress{
		Name: name,
		Url:  url,
		Cnt:  1,
	}
}
func (act *Actress) Plus() {
	act.Cnt = act.Cnt + 1
}

func (act Actress) PngBase64() string {
	path := act.Url
	if !utils.ExistsFiles(path) {
		path = act.Url
	}
	return utils.ImageToString(path)
}
