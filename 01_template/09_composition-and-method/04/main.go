package main

import (
    "os"
    "log"
    "text/template"
)

type person struct {
    Name string
    Age int
}

func (p person) SomeProcessing() int {
    return 7
}

func (p person) AgeDbl() int {
    return p.Age * 2
}

func (p person)TakesArg(x int) int {
    return x * 2
}

var tpl  *template.Template

func init() {
    tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
    p1 := person {
        Name: "Vlad",
        Age: 42,
    }

    err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", p1)
    if err != nil {
        log.Fatalln(err)
    }
}
