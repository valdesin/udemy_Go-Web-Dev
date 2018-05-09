package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go serve(conn)

	}
}

func serve(conn net.Conn) {
	defer conn.Close()
	var method, URl string
	var i int
	s := bufio.NewScanner(conn)
	for s.Scan() {
		ln := s.Text()
		fmt.Println(ln)
		if i == 0 {
			data := strings.Fields(ln)
			method = data[0]
			URl = data[1]
			fmt.Println("METHOD: ", method)
			fmt.Println("URL: ", URl)
		}
		if ln == "" {
			fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
			break
		}
		i++
	}

	switch {
	case method == "GET" && URl == "/":
		handleIndex(conn)
	case method == "GET" && URl == "/apply":
		handleApply(conn)
	case method == "POST" && URl == "/apply":
		handleApplyPost(conn)
	default:
		handleDefaolt(conn)
	}

}

func handleIndex(conn net.Conn) {
	body := `	
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>GET INDEX</title>
	</head>
	<body>
		<h1>"GET INDEX"</h1>
		<a href="/">index</a><br>
		<a href="/apply">apply</a><br>
	</body>
	</html>
	`

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func handleApply(conn net.Conn) {
	body := `
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="utf-8">
			<title>GET APPLY</title>
		</head>
		<body>
			<h1>GET APPLY</h1>
			<a href="/">index</a><br>
			<a href="/apply">apply</a><br>
			<form method="POST" action="/apply">
				<input type="hidden" value="In my good death">
				<input type="submit" value="submit">
			</form>
		</body>
		</html>
	`
	io.WriteString(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func handleApplyPost(conn net.Conn) {
	body := `
	<!DOCTYPE html>
	<html>
	<head>
		<meta charset="utf-8">
		<title>POST APPLY</title>
	</head>
	<body>
		<h1>POST APPLY</h1>
		<a href="/">index</a><br>
		<a href="/apply">apply</a><br>
	</body>
	</html>
	`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func handleDefaolt(conn net.Conn) {
	body := `
	<!DOCTYPE html>
	<html>
	<head>
		<meta charset="utf-8">
		<title>default</title>
	</head>
	<body>
		<h1>Default</h1>
	</body>
	</html>
	`

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\t\n")
	io.WriteString(conn, "\t\n")
	io.WriteString(conn, body)
}
