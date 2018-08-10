package main

import (
	"fmt"
	"net/http"
	"time"
)

var t = time.Now().Add(time.Duration(10))

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":9090", nil)
}

func set(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:    "Time & COOKIE",
		Value:   "cookie with set time alive!",
		Expires: t,
	})
}

func read(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("Time & COOKIE")
	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	}

	fmt.Fprintln(res, c)
}
