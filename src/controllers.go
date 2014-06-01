package main

import (
	"fmt"
	"net/http"
	"strings"
)

func LoginHandler(r http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(r, GetLoginView(nil))
}

func HomeHandler(r http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(r, GetHomeView(nil))
}

func RedirectToHome(r http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	if !strings.HasSuffix(path, "/"){
		path += "/"
	}
	http.Redirect(r, req, path+"home", http.StatusFound)
}
