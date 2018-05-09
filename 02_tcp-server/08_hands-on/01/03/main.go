package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(foo))
	http.Handle("/dog", http.HandlerFunc(dog))
	http.Handle("/me", http.HandlerFunc(myName))

	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "WELCOME!\n")
}

func dog(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "waw graw auuu!\n")
}

func myName(w http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("mn.gohtml"))

	err := tpl.ExecuteTemplate(w, "mn.gohtml", "Vlad!")
	if err != nil {
		log.Fatalln(err)
	}
}
