package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/bared", bared)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":9090", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at foo: ", req.Method, "\n\n")
}
func bar(res http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at bar: ", req.Method, "\n\n")

	http.Redirect(res, req, "/", http.StatusSeeOther)
}
func bared(res http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at bared: ", req.Method, "\n\n")

	tpl.ExecuteTemplate(res, "index.gohtml", nil)
}
