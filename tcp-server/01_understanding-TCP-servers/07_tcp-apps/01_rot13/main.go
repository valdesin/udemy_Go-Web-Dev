package main

import (
    "fmt"
    "log"
    "net"
    "bufio"
    "strings"
    "time"
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
        }

    go handle(conn)
    }
}

func handle(conn net.Conn) {
    err := conn.SetDeadline(time.Now().Add(5 * time.Second))
    if err != nil {
        log.Fatalln("CONNECTION TIMEOUT")
    }
    defer conn.Close()

    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        ls := strings.ToLower(scanner.Text())
        bs := []byte(ls)
        r13 := rot13(bs)
        fmt.Fprintf(conn, "%s - %s",ls, r13)
    }
}

func rot13(bs []byte) []byte {
    var r13 = make([]byte, len(bs))
    for i, v := range bs {
        if v <= 109 {
            r13[i] = v + 13
        } else {
            r13[i] = v - 13
        }
    }
    return r13
}
