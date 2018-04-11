package main

import (
    "net"
    "fmt"
    "log"
    "io/ioutil"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:9090")
    if err != nil {
        log.Fatalln(err)
    }
    defer conn.Close()

    bs, err := ioutil.ReadAll(conn)
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Println(string(bs))
}
