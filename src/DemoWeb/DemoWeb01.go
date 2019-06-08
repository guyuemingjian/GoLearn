package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/goDemo01",
		func(writer http.ResponseWriter, request *http.Request) {
			fmt.Fprint(writer,"<h1>hello world</h1>")
		})

	http.ListenAndServe(":8081",nil)
}
