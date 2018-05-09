package main

import (
    "fmt"
    "log"
    "net"
    "bufio"
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
            log.Fatalln(err)
            continue
        }
         go handle(conn)
    }
}

func handle(conn net.Conn) {
    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        ls := scanner.Text()
        fmt.Println(ls)
        fmt.Fprintf(conn, "I heard you say: %s\n", ls)
    }
    defer conn.Close()

    fmt.Println("Code got here.")
}
