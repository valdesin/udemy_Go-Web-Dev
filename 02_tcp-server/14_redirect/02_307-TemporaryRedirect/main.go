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
	fmt.Print("Your request method at foo: ", req.Method, "\n\n\n")
}
func bar(res http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at dar: ", req.Method, "\n\n\n")

	http.Redirect(res, req, "/", http.StatusTemporaryRedirect)
}
func bared(res http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at dared: ", req.Method, "\n\n\n")

	tpl.ExecuteTemplate(res, "index.gohtml", nil)
}
