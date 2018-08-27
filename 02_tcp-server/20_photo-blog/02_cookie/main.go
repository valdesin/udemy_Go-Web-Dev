package main

import (
	"html/template"
	"net/http"

	"github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":9090", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	c := getCookie(res, req)
	tpl.ExecuteTemplate(res, "index.gohtml", c)
}

func getCookie(res http.ResponseWriter, req *http.Request) string {
	c, err := req.Cookie("session")
	if err != nil {
		uId, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: uId.String(),
		}
		http.SetCookie(res, c)
	}
	return c.Value
}
