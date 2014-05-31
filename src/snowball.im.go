package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func Hello(r http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name := vars["name"]
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(r, "Hello %s!", name)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", Hello)
	r.HandleFunc("/{name}", Hello)

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
