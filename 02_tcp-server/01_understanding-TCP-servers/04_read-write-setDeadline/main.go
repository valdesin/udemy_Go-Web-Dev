package main

import (
    "net"
    "fmt"
    "log"
    "time"
    "bufio"
)

func main() {
    li, err := net.Listen("tcp", ":8080")
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
    err := conn.SetDeadline(time.Now().Add(10 * time.Second))
    if err != nil {
        log.Fatalln("CONNECTION TIMEOUT")
    }

    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        ls := scanner.Text()
        fmt.Println(ls)
        fmt.Fprintf(conn, "I heard you say: %s\n", ls)
    }
    defer conn.Close()

    fmt.Println("***CODE GOT HERE***")
}
