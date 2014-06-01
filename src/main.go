package main

import (
	"github.com/codegangsta/negroni"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/dineshsaravanan/snowball.im/src/logger"
	s "github.com/dineshsaravanan/negroni-contrib/sessions"
	"github.com/gorilla/sessions"
)

func main() {
	//Create new negroni
	logger.GetStaticLogger().Logger.Println("Starting app...")

	//router
	r := mux.NewRouter()
	r.HandleFunc("/login", LoginHandler)
	r.HandleFunc("/home", HomeHandler)
	r.HandleFunc("/", RedirectToHome)

	//Negroni is mixed here
	n := negroni.New()

	//Negroni ingredients
	//use router
	n.UseHandler(r)

	//Creating static middleware, that will serve files from /public folder, we are also making sure other folders are
	//off limits
	static := NewStatic(http.Dir("./public"))
	static.Prefix = "/public"

	//Settimg up default middlewares
	n.Use(negroni.NewRecovery())
	n.Use(negroni.NewLogger())
	n.Use(s.NewSessions("alright", sessions.NewCookieStore()))

	n.Use(static);

	//start listening to port 8080 for http requests
	n.Run(":8080")
}
