package main

import (
    "fmt"
    "log"
    "net"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:9090")
    if err != nil {
        log.Fatalln(err)
    }
    defer conn.Close()

    fmt.Fprintln(conn, "I dialed you!")
}
