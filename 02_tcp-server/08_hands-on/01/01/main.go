package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/me", myName)

	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "WELCOME!\n")
}

func dog(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "waw graw auuu!\n")
}

func myName(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Vlad")
}
