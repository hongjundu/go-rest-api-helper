package apihelper

const (
	ErrorCodeServerError = "server_error"
)

type ApiError interface {
	error
	Code() string
}

type apiError struct {
	code string
	msg  string
}

func NewError(code, msg string) ApiError {
	return &apiError{code, msg}
}

func (this *apiError) Code() string {
	return this.code
}

func (this *apiError) Error() string {
	return this.msg
}
