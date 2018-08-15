package main

import (
	"html/template"
	"net/http"

	"github.com/satori/go.uuid"
)

type user struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template
var dbUser = map[string]user{}
var dbSession = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":9090", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		uId, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: uId.String(),
		}
		http.SetCookie(res, c)
	}

	var u user
	if un, ok := dbSession[c.Value]; ok {
		dbUser[un] = u
	}

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		u = user{un, f, l}
		dbSession[c.Value] = un
		dbUser[un] = u
	}
	tpl.ExecuteTemplate(res, "index.gohtml", u)
}

func bar(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	un := dbSession[c.Value]
	u := dbUser[un]

	tpl.ExecuteTemplate(res, "bar.gohtml", u)
}
