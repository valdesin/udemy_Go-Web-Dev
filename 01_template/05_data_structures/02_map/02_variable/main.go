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
    sage := map[string]string {
        "India":    "Gandhi",
		"America":  "MLK",
		"Meditate": "Buddha",
		"Love":     "Jesus",
		"Prophet":  "Muhammad",
    }

    err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sage)
    if err != nil {
        log.Fatalln(err)
    }
}
