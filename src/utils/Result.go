package utils

type Result struct {
	MessageCode
	Data interface{}
}

func (res *Result) Success() {
	res.MessageCode.Success()
}
func NewResult() Result {
	return Result{}
}
func (res *Result) Fail() {
	res.MessageCode.Fail()
}
