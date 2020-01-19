package datamodels

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
func (act Actress) Plus() Actress {
	act.Cnt = act.Cnt + 1
	return act
}
