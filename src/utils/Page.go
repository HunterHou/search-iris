package utils

type Page struct {
	PageNo    int
	PageSize  int
	StartNo   int
	TotalPage int
	EachPage  []int

	Data      interface{}
	KeyWord   string
	TotalCnt  int
	ResultCnt int
	CurCnt    int

	ResultSize string
	TotalSize  string
	CurSize    string
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

func (p *Page) SetResultCnt(resultCnt int) {
	p.ResultCnt = resultCnt
	totalPage := resultCnt/p.PageSize + 1
	p.TotalPage = totalPage
	var pageList []int
	for i := 0; i < totalPage; i++ {
		pageList = append(pageList, (i + 1))
	}
	p.EachPage = pageList
}
