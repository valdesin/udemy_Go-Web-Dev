package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9090")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer conn.Close()

	fmt.Fprint(conn, "GET / HTTP/1.1")

	scan := bufio.NewScanner(conn)
	for scan.Scan() {
		ls := scan.Text()
		fmt.Println(ls)
	}
}
