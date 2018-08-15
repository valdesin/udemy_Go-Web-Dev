package main

import (
	"html/template"
	"net/http"

	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
}

var tpl *template.Template
var dbSessions = map[string]string{}
var dbUsers = map[string]user{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	bs, _ := bcrypt.GenerateFromPassword([]byte("session"), bcrypt.MinCost)
	dbUsers["test@test.com"] = user{"test@test.com", bs, "Vlad", "Fill"}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":9090", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	u := getUser(res, req)
	tpl.ExecuteTemplate(res, "index.gohtml", u)
}

func bar(res http.ResponseWriter, req *http.Request) {
	if !alreadyLogIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	u := getUser(res, req)
	tpl.ExecuteTemplate(res, "bar.gohtml", u)
}

func signup(res http.ResponseWriter, req *http.Request) {
	if alreadyLogIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	var u user
	if req.Method == http.MethodPost {

		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")

		if _, ok := dbUsers[un]; ok {
			http.Error(res, "Username already taken", http.StatusForbidden)
			return
		}

		uId, err := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: uId.String(),
		}

		http.SetCookie(res, c)
		dbSessions[c.Value] = un

		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(res, "Internal server error", http.StatusInternalServerError)
		}

		u = user{un, bs, f, l}
		dbUsers[un] = u

		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(res, "signup.gohtml", u)
}

func login(res http.ResponseWriter, req *http.Request) {
	if alreadyLogIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	var u user
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")

		u, ok := dbUsers[un]
		if !ok {
			http.Error(res, "Username and/or password do not match", http.StatusForbidden)
			return
		}

		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(res, "Username and/or password do not match", http.StatusForbidden)
			return
		}

		uId, err := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: uId.String(),
		}
		http.SetCookie(res, c)
		dbSessions[c.Value] = un

		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(res, "login.gohtml", u)
}

func logout(res http.ResponseWriter, req *http.Request) {
	if !alreadyLogIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	c, _ := req.Cookie("session")
	delete(dbSessions, c.Value)
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(res, c)

	http.Redirect(res, req, "/login", http.StatusSeeOther)
}
