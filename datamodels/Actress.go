package datamodels

import "../utils"

type Actress struct {
	Name    string
	Url     string
	Cnt     int
	size    int64
	SizeStr string
}

func NewActres(name string, url string, size int64) Actress {
	return Actress{
		Name:    name,
		Url:     url,
		Cnt:     1,
		size:    size,
		SizeStr: utils.GetSizeStr(size),
	}
}
func (act *Actress) PlusCnt() {
	act.Cnt = act.Cnt + 1
}

func (act *Actress) PlusSize(size int64) {
	act.size = act.size + size
	act.SizeStr = utils.GetSizeStr(act.size)
}

func (act Actress) PngBase64() string {
	path := act.Url
	if !utils.ExistsFiles(path) {
		path = act.Url
	}
	return utils.ImageToString(path)
}
