package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/sponge.jpg", dogImg)
	http.ListenAndServe(":9090", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
		<img src="/sponge.jpg">
		`)
}

func dogImg(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("sponge.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
	}
	defer f.Close()

	io.Copy(w, f)
}
