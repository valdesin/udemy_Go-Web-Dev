package main

import (
    "os"
    "log"
    "text/template"
)

func main() {
    tpl, err := template.ParseGlob("templates/*")
    if err != nil {
        log.Fatalln(err)
    }

    err = tpl.Execute(os.Stdout, nil)
    if err != nil {
        log.Fatalln(err)
    }

    err = tpl.ExecuteTemplate(os.Stdout, "01_vespa.gmao", nil)
    if err != nil {
        log.Fatalln(err)
    }

    err = tpl.ExecuteTemplate(os.Stdout, "two.gmao", nil)
    if err != nil {
        log.Fatalln(err)
    }

    err = tpl.ExecuteTemplate(os.Stdout, "one.gmao", nil)
    if err != nil {
        log.Fatalln(err)
    }
}
