package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hello %s!\n", ps.ByName("user"))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:user", Hello)

	http.ListenAndServe(":9090", router)
}
