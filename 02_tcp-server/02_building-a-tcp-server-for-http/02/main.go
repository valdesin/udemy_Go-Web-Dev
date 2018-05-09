package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	i := 0
	scan := bufio.NewScanner(conn)
	for scan.Scan() {
		ls := scan.Text()
		if i == 0 {
			url := strings.Fields(ls)[1]
			fmt.Fprintln(conn, "URLL:", url)
			i++
		}
	}
}
