package utils

type Result struct {
	MessageCode
	Data interface{}
}

func ResultSuccess() Result {
	result := Result{}
	result.Success()
	return result
}
