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
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	io.WriteString(conn, "\r\nIN-MEMORY DATABASE\r\n\r\n"+
		"USE:\r\n"+
		"\tSET key value \r\n"+
		"\tGET key \r\n"+
		"\tDEL key \r\n\r\n"+
		"EXAMPLE:\r\n"+
		"\tSET fav chocolate \r\n"+
		"\tGET fav \r\n\r\n\r\n")

	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fls := strings.Fields(ln)

		if len(fls) < 1 {
			continue
		}
		switch fls[0] {
		case "GET":
			fmt.Println("GET", fls[1])
			k := fls[1]
			v := data[k]
			fmt.Println(v)
			fmt.Fprintf(conn, "%s\r\n", v)
		case "SET":
			if len(fls) != 3 {
				fmt.Fprintln(conn, "EXPECTED VALUE\r\n")
				continue
			}
			k := fls[1]
			v := fls[2]
			data[k] = v
			fmt.Println(data)
		case "DEL":
			k := fls[1]
			delete(data, k)
		default:
			fmt.Fprintln(conn, "INVALID COMMAND "+fls[0]+"\r\n")
			continue
		}
	}
}
