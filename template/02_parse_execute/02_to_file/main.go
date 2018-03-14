package main

import (
    "os"
    "log"
    "text/template"
)

func main() {
    tpl, err := template.ParseFiles("tpl.gohtml")
    if err != nil {
        log.Fatalln(err)
    }

    dst, err := os.Create("index.html")
    if err != nil {
        log.Fatalln(err)
    }
    defer dst.Close()

    err = tpl.Execute(dst, nil)
    if err != nil {
        log.Fatalln(err)
    }
}
