package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/fovico.ico", http.NotFoundHandler())
	http.ListenAndServe(":9090", nil)
}

func foo(res http.ResponseWriter, q *http.Request) {
	v := q.FormValue("v")
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `
		<form method="GET">
			<input type="text" name="v">
			<input type="submit">
		</form>
	<br>`+v)
}
