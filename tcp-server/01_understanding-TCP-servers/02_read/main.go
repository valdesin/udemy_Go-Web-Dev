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
        go hundle(conn)
    }
}

func hundle(conn net.Conn) {
    scann := bufio.NewScanner(conn)
    for scann.Scan() {
        s := scann.Text()
        fmt.Println(s)
    }
    defer conn.Close()

    fmt.Println("Code got here.")
}
