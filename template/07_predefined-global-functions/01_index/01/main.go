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

    err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", xs)
    if err != nil {
        log.Fatalln(err)
    }
}
