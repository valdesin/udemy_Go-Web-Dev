package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/sponge.jpg", dogImg)
	http.ListenAndServe(":9090", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/sponge.jpg">`)
}

func dogImg(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "sponge.jpg")
}
