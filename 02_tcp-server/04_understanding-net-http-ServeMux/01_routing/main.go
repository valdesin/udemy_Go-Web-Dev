package main

import (
	"io"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	switch req.URL.Path {
	case "/dog":
		io.WriteString(w, "Dogy Dogy Dogy!!!\n")
	case "/cat":
		io.WriteString(w, "Caty Caty Caty!!!\n")
	}
}

func main() {
	var h hotdog
	http.ListenAndServe(":9090", h)
}
