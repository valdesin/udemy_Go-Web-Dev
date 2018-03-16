package main

import (
    "os"
    "log"
    "text/template"
)

var tpl *template.Template

type sage struct {
    Name string
    Motto string
}

func init() {
    tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
    sage := sage {
        Name: "Buddha",
        Motto: "The belief of no beliefs",
    }

    err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sage)
    if err != nil {
        log.Fatalln(err)
    }
}
