package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":9090", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	cxt := req.Context()

	cxt = context.WithValue(cxt, "userId", 1994)
	cxt = context.WithValue(cxt, "name", "vlad")

	result := dbAcces(cxt)
	fmt.Fprintln(res, result)
}

func dbAcces(ctx context.Context) int {
	uid := ctx.Value("userId").(int)
	return uid
}

func bar(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	log.Println(ctx)
	fmt.Fprintln(res, ctx)
}
