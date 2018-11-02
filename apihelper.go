package apihelper

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func ReadJsonRequestBody(readCloser io.ReadCloser, obj interface{}) error {
	if body, err := ioutil.ReadAll(readCloser); err == nil {
		return json.Unmarshal(body, obj)
	} else {
		return err
	}
}

func ReadFormValueInt(r *http.Request, formKey string, defaultValue int64) (n int64, err error) {
	s := r.FormValue(formKey)

	if len(s) == 0 {
		return defaultValue, nil
	}

	n, err = strconv.ParseInt(s, 10, 64)
	if err != nil {
		err = NewError(ErrorCodeBadArgument, err.Error())
	}

	return
}
