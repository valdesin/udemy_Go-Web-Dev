package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var c *http.Cookie
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/bowzer", bowzer)
	http.HandleFunc("/dog/bowzer/bowzerpics", bowzerpics)
	http.HandleFunc("/cat", cat)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":9090", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("user-name")
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%T\t", c)
	}
	tpl.ExecuteTemplate(res, "index.gohtml", c)
}

func bowzer(res http.ResponseWriter, req *http.Request) {
	c = &http.Cookie{
		Name:  "user-name",
		Value: "some value",
		//Path:  "/dog",
	}
	http.SetCookie(res, c)
	tpl.ExecuteTemplate(res, "bowzer.gohtml", c)
}

func bowzerpics(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("user-name")
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%T\t", c)
	}
	tpl.ExecuteTemplate(res, "bowzerpics.gohtml", c)
}

func cat(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("user-name")
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%T\t", c)
	}
	tpl.ExecuteTemplate(res, "cat.gohtml", c)
}
