package main

import (
    "os"
    "log"
    "text/template"
)

var tpl *template.Template

func init() {
    tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
    xs := []string{"zero", "one", "two", "three", "four", "five"}

    data := struct {
        Words []string
        Lname string
    }{
        Words: xs,
        Lname: "Filonenko",
    }

    err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
    if err != nil {
        log.Fatalln(err)
    }
}
