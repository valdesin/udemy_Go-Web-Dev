package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
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

		s := bufio.NewScanner(conn)
		for s.Scan() {
			ln := s.Text()
			fmt.Println(ln)
		}

		io.WriteString(conn, "I see you connected!!!\n")
		conn.Close()
	}
}
