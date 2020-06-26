package apihelper

const (
	ErrorCodeServerError = "server_error"
	ErrorCodeBadArgument = "bad_argument"
)

type ApiError interface {
	error
	StatusCode() int
	Code() string
}

type apiError struct {
	statusCode int
	code       string
	msg        string
}

func NewError(code, msg string) ApiError {
	return &apiError{code: code, msg: msg}
}

func NewErrorWithStatusCode(statusCode int, code, msg string) ApiError {
	return &apiError{statusCode: statusCode, code: code, msg: msg}
}

func (this *apiError) StatusCode() int {
	return this.statusCode
}

func (this *apiError) Code() string {
	return this.code
}

func (this *apiError) Error() string {
	return this.msg
}
