package main

import (
    "os"
    "fmt"
    "log"
    "net"
    "bufio"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:9090")
    if err != nil {
        log.Fatalln(err)
    }
    defer conn.Close()

    read := bufio.NewReader(os.Stdin)
    bs, _, _ := read.ReadLine()
    fmt.Fprintln(conn, string(bs))

    scan := bufio.NewScanner(conn)
    for scan.Scan() {
        fmt.Println(scan.Text())
    }
}
