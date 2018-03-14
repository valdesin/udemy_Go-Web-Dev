package main

import (
    "io"
    "os"
    "fmt"
    "log"
    "strings"
) 

func main() {
    name := os.Args[1]
    fmt.Println(os.Args[0])
    fmt.Println(os.Args[1])
    str := fmt.Sprint(`
            <!DOCTYPE html>
            <html lang="en">
                <head>
                    <meta charset="UTF-8">
                    <title>Hello World!</title>
                </head>
                <body>
                    <h1>` +
		                name +
		           `</h1>
                </body>
            </html>
    `)

    dst, err := os.Create("index.html")
    if err != nil{
        log.Fatal(err)
    }
    defer dst.Close()

    io.Copy(dst, strings.NewReader(str))
}
