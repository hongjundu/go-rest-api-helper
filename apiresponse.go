package apihelper

const (
	ResponseStatusOK    = "ok"
	ResponseStatusError = "error"
)

func NewOKResponse(data interface{}) interface{} {
	return &apiResponse{Status: ResponseStatusOK, Data: data}
}

func NewListResponse(count int, list interface{}) interface{} {
	return NewOKResponse(&responseListData{List: list, Count: count})
}

func NewErrorResponse(err error) interface{} {
	if err == nil {
		return NewOKResponse(nil)
	} else {
		if apiErr, ok := err.(ApiError); ok {
			return &apiResponse{Status: ResponseStatusError, Code: apiErr.Code(), Msg: apiErr.Error()}
		} else {
			return &apiResponse{Status: ResponseStatusError, Code: 500, Msg: err.Error()}
		}
	}
}

// ok response format:
// {"status":"ok","data":{ ... }}
// error response format:
// {"status":"error","code":"error_code","msg":"error message description"}

type apiResponse struct {
	Status string      `json:"status"`
	Code   int         `json:"code,omitempty"`
	Msg    string      `json:"msg,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

type responseListData struct {
	Count int         `json:"count"`
	List  interface{} `json:"list,omitempty"`
}
