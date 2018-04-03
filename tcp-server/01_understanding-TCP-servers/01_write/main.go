package main

import (
    "io"
    "fmt"
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
            log.Fatalln(err)
            continue
        }
        
        io.WriteString(conn, "\nHello from TCP server\n")
        fmt.Fprintln(conn, "How is your day?")
        fmt.Fprintf(conn, "%v", "Well I hope") 
    }
}
