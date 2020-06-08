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

func (p *Page) SetResultCnt(resultCnt int, pageNo int) {
	p.ResultCnt = resultCnt
	totalPage := resultCnt/p.PageSize + 1
	p.TotalPage = totalPage
	var pageList []int
	var headNum = 7
	var middNum = 4
	var middPage = pageNo
	if pageNo <= 5 || pageNo >= (totalPage-5) {
		middPage = totalPage / 2
	}
	for i := 0; i < totalPage; i++ {
		if i < headNum || i > totalPage-headNum {
			pageList = append(pageList, (i + 1))
			continue
		}
		if i < (middPage+middNum) && i > (middPage-middNum) {
			pageList = append(pageList, (i + 1))
			continue
		}

	}
	p.EachPage = pageList
}
