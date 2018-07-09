package apihelper

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func ReadJsonRequestBody(readCloser io.ReadCloser, obj interface{}) error {
	if body, err := ioutil.ReadAll(readCloser); err == nil {
		return json.Unmarshal(body, obj)
	} else {
		return err
	}
}
