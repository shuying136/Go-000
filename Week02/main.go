package main

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

type appHandler func(writer http.ResponseWriter,
	request *http.Request) error

func errWrapper(
	handler appHandler) func(
	http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter,
		request *http.Request) {
		err := handler(writer, request)

		if err != nil {
			fmt.Printf("original error:%T %v\n", errors.Cause(err), errors.Cause(err))
			http.Error(writer,
				err.Error(),
				http.StatusBadRequest)
			//fmt.Printf("stack trace:\n%+v\n",err)

		}
	}
}

func main() {
	http.HandleFunc("/getResult",
		errWrapper(selectValue))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
}
