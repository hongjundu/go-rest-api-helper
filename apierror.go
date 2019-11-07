package apihelper

type ApiError interface {
	error
	Code() int
}

type apiError struct {
	code int
	msg  string
}

func NewError(code int, msg string) ApiError {
	return &apiError{code: code, msg: msg}
}

func (this *apiError) Code() int {
	return this.code
}

func (this *apiError) Error() string {
	return this.msg
}
