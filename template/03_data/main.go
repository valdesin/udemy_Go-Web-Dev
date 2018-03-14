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
    err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 244)
    if err != nil {
        log.Fatalln(err)
    }
}
