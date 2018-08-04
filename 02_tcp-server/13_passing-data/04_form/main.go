package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favico.ico", http.NotFoundHandler())
	http.ListenAndServe(":9090", nil)
}

func foo(res http.ResponseWriter, q *http.Request) {
	f := q.FormValue("firstName")
	l := q.FormValue("lastName")
	sub := q.FormValue("subscribe") == "on"

	err := tpl.ExecuteTemplate(res, "index.gohtml", person{f, l, sub})
	if err != nil {
		http.Error(res, err.Error(), 500)
		log.Fatalln(err)
	}
}
