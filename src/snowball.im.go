package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"net/http"
	"github.com/gorilla/mux"
	"strings"
)

func LoginHandler(r http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(r, "Login here - me")
}

func HomeHandler(r http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(r, "home asas")
}

func RedirectToHome(r http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	if !strings.HasSuffix(path, "/"){
		path += "/"
	}
	http.Redirect(r, req, path+"home", http.StatusFound)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", LoginHandler)
	r.HandleFunc("/home", HomeHandler)
	r.HandleFunc("/", RedirectToHome)

	n := negroni.Classic()

	n.UseHandler(r)
	n.Run(":8080")
}
