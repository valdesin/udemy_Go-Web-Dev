package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":9090", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at foo: ", req.Method, "\n\n")
}
func bar(res http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at bar: ", req.Method, "\n\n")

	http.Redirect(res, req, "/", http.StatusMovedPermanently)
}
