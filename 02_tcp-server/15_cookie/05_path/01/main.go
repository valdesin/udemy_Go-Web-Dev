package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":9090", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(res, http.StatusText(400), http.StatusBadRequest)
		return
	}
	fmt.Println(c)
	fmt.Fprintln(res, "YOUR COOKIE: ", c)
}

func set(res http.ResponseWriter, req *http.Request) {
	c := &http.Cookie{
		Name:  "my-cookie",
		Value: "some valuoes!!!",
	}

	http.SetCookie(res, c)
	fmt.Println(c)
	fmt.Fprintln(res, "YOUR COOKIE!!! ", c)
}

func read(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(res, http.StatusText(400), http.StatusBadRequest)
		return
	}
	fmt.Println(c)
	fmt.Fprintln(res, "YOUR COOKIE: ", c)
}
