package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

type session struct {
	un           string
	lastActivity time.Time
}

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]session{}
var dbSessionCleaned time.Time

const sessionLength int = 30

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	dbSessionCleaned = time.Now()
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/singup", singup)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/checkUserName", checkUserName)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":9090", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	u := getUser(res, req)
	showSessions()
	tpl.ExecuteTemplate(res, "index.gohtml", u)
}

func bar(res http.ResponseWriter, req *http.Request) {
	u := getUser(res, req)
	if !alreadyLoggedIn(res, req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "007" {
		http.Error(res, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}
	showSessions()
	tpl.ExecuteTemplate(res, "bar.gohtml", u)
}

func singup(res http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(res, req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	var u user
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := req.FormValue("role")

		if _, ok := dbUsers[un]; ok {
			http.Error(res, "Username already taken!", http.StatusForbidden)
			return
		}
		uId, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: uId.String(),
		}
		c.MaxAge = sessionLength
		http.SetCookie(res, c)
		dbSessions[c.Value] = session{un, time.Now()}
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(res, "Internal server error", http.StatusInternalServerError)
			return
		}
		u := user{un, bs, f, l, r}
		dbUsers[un] = u
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	showSessions()
	tpl.ExecuteTemplate(res, "singup.gohtml", u)
}

func login(res http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(res, req) {
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

		uId, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: uId.String(),
		}
		c.MaxAge = sessionLength
		http.SetCookie(res, c)
		dbSessions[c.Value] = session{un, time.Now()}
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	showSessions()
	tpl.ExecuteTemplate(res, "login.gohtml", u)
}

func logout(res http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(res, req) {
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

	if time.Now().Sub(dbSessionCleaned) > (time.Second * 30) {
		go cleanSessions()
	}
	http.Redirect(res, req, "/login.gohtml", http.StatusSeeOther)
}

func checkUserName(res http.ResponseWriter, req *http.Request) {
	sampleUsers := map[string]bool{
		"test@example.com": true,
		"jame@bond.com":    true,
		"moneyp@uk.gov":    true,
	}
	bs, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}

	sbs := string(bs)
	fmt.Println("USERNAME: ", sbs)

	fmt.Fprint(res, sampleUsers[sbs])
}
