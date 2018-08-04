package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":9090", nil)
}

func foo(res http.ResponseWriter, r *http.Request) {
	q := r.FormValue("q")
	fmt.Fprintln(res, "My search is: "+q)
}
