package main

import (
	"encoding/json"
	"github.com/hongjundu/go-rest-api-helper"
	"log"
	"net/http"
)

func main() {
	log.Printf("server starts...")

	http.HandleFunc("/ok", wrapHandler(handlerReturnsOK))
	http.HandleFunc("/error", wrapHandler(handlerReturnsError))

	log.Fatal(http.ListenAndServe(":8080", nil))

	log.Printf("server exists")
}

func handlerReturnsOK(r *http.Request) (response interface{}, err error) {
	type Resp struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}
	response = &Resp{FirstName: "Jon", LastName: "Snow"}
	return
}

func handlerReturnsError(r *http.Request) (response interface{}, err error) {
	err = apihelper.NewError("some_error_code", "some error message")
	return
}

func wrapHandler(handler func(r *http.Request) (interface{}, error)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var response interface{}
		var responseBody []byte

		if response, err = handler(r); err == nil {
			responseBody, err = json.Marshal(response)
		}

		if err != nil {
			responseBody, _ = json.Marshal(go-rest-api-helper.NewErrorResponse(err))
		}

		w.Write(responseBody)
	}
}
