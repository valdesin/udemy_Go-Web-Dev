package main

import (
    "os"
    "log"
    "html/template"
)

type Page struct {
    Title, Heading, Input string
}

var tpl *template.Template

func init() {
    tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
    p := Page {
        Title: "Nothing Escaped",
        Heading: "Nothing is escaped with text/template",
        Input: `<script>alert("Yow!");</script>`,
    }

    err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", p)
    if err != nil {
        log.Fatalln(err)
    }
}
