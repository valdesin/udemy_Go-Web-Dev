package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/authenticate", auth)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":9090", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		c = &http.Cookie{
			Name:  "session",
			Value: "",
		}
	}

	if req.Method == http.MethodPost {
		e := req.FormValue("email")
		c.Value = e + "|" + getCode(e)
	}

	http.SetCookie(res, c)

	io.WriteString(res, `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="utf-8">
			<title>DOCUMENT</title>
		</head>
		<body>
			<form method="POST">
				<input type="email" name="email">
				<input type="submit">
			</form>
			<a href="/authenticate">Validate This `+c.Value+`</a>
		</body>
		</html>
	`)
}

func auth(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	if c.Value == "" {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	arrV := strings.Split(c.Value, "/")
	email := arrV[0]
	codeRcvd := arrV[1]
	codeCheck := getCode(email + "s")

	if codeRcvd != codeRcvd {
		fmt.Println("HMAC codes didn't match")
		fmt.Println(codeRcvd)
		fmt.Println(codeCheck)
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	io.WriteString(res, `
	<!DOCTYPE html>
	<html>
		<body>
		<h1>`+codeRcvd+` RECIVED</h1>
		<h1>`+codeCheck+` RECALCULATED</h1>
		</body>
	</html>
	`)
}

func getCode(s string) string {
	h := hmac.New(sha256.New, []byte("my-key"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
