package utils

type Page struct {
	PageNo    int
	PageSize  int
	StartNo   int
	TotalPage int
	EachPage  []int
	TotalCnt  int
	Data      interface{}
	KeyWord   string
	CurCnt    int
}

func NewPage() Page {
	return Page{
		PageNo:   0,
		PageSize: 0,
		StartNo:  0,
		TotalCnt: 0,
		Data:     nil,
		KeyWord:  "",
	}
}

func (p Page) SetTotalCnt(totalCnt int) Page {
	p.TotalCnt = totalCnt
	totalPage := totalCnt/p.PageSize + 1
	p.TotalPage = totalPage
	var pageList []int
	for i := 0; i < totalPage; i++ {
		pageList = append(pageList, (i + 1))
	}
	p.EachPage = pageList
	return p
}
