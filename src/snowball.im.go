package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"net/http"
	"github.com/gorilla/mux"
	"strings"
	"github.com/dineshsaravanan/snowball.im/src/logger"
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
	//Create new negroni
	n := negroni.New()

	//router
	r := mux.NewRouter()
	r.HandleFunc("/login", LoginHandler)
	r.HandleFunc("/home", HomeHandler)
	r.HandleFunc("/", RedirectToHome)

	//use router
	n.UseHandler(r)

	//Creating static middleware, that will serve files from /public folder, we are also making sure other folders are
	//off limits
	static := NewStatic(http.Dir("./public"))
	static.Prefix = "/public"

	//Settimg up default middlewares
	n.Use(negroni.NewRecovery())
	n.Use(negroni.NewLogger())
	n.Use(logger.NewLogger())
	n.Use(static);

	//start listening to port 8080 for http requests
	n.Run(":8080")
}
