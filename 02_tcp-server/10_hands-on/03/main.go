package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.html"))
}

func main() {
	han := http.FileServer(http.Dir("public"))
	http.Handle("/pics/", han)
	http.HandleFunc("/", dog)
	http.ListenAndServe(":9090", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Fatalln(err)
	}
}