package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":80", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Oh yeah, I'm running on AWS. \nHello Vladyslav how are you ?")
}
