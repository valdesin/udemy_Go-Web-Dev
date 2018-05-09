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
	io.WriteString(w, `<img src="/sponge.jpg">`)
}

func dogImg(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("sponge.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
	}

	http.ServeContent(w, req, f.Name(), fi.ModTime(), f)
}
