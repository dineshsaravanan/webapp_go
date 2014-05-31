package main

import (
	"fmt"
	"net/http"
)

func hello(r http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(r, "Hello world!");
}

func main() {
	http.HandleFunc("/", hello)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
