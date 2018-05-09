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
    sage := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad"}
    err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sage)
    if err != nil {
        log.Fatalln(err)
    }
}
