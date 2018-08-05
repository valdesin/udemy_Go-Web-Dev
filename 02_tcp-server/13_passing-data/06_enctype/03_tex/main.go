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
	FirstName string
	LastName  string
	Subscribe bool
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":9090", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	bs := make([]byte, req.ContentLength)
	req.Body.Read(bs)
	body := string(bs)

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := tpl.ExecuteTemplate(res, "index.gohtml", body)
	if err != nil {
		http.Error(res, err.Error(), 500)
		log.Fatalln(err)
	}
}
