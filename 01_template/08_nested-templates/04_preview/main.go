package main

import (
    "log"
    "net/http"
    "html/template"
)

var tpl *template.Template

func init() {
    tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
    http.HandleFunc("/", index)
    http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
    http.ListenAndServe(":9000", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
    err := tpl.ExecuteTemplate(res, "index.gohtml", nil)
    if err != nil {
        log.Fatalln(err)
    }
}
