package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("favico.ico", http.NotFoundHandler())
	http.ListenAndServe(":9090", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	var s string
	fmt.Println(req.Method)
	if req.Method == http.MethodPost {

		f, h, err := req.FormFile("f")
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		fmt.Println("\nname ", f, "\nheaders ", h, "\nerror ", err)

		sb, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(sb)
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `
		<form method="POST" enctype="multipart/form-data">
		<input type="file" name="f">
		<br>
		<input type="submit">
		</form>
	`+s)
}
