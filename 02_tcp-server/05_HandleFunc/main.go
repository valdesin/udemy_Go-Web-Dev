package main

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Dog Dog Dog")
}

func c(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Cat Cat Cat")
}

func main() {
	http.HandleFunc("/dog", d)
	http.HandleFunc("/cat", c)

	http.ListenAndServe(":9090", nil)
}
