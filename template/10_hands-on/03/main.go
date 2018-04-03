package main

import (
    "os"
    "log"
    "text/template"
)

type menu struct {
    Breakfast, Lunch, Dinner []string
}

var tpl *template.Template

func init() {
    tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
    menu := menu {
        Breakfast: []string{"Grand Slam","Kiwi Slam","French Slam"},
        Lunch: []string{"Chicken Fried Steak", "Swiss Chard", "Soup Albondigas"},
        Dinner: []string{"Two Tacos-Beef", "Burritos Chile Verde", "Fajitas Chicken"},
    }

    err := tpl.Execute(os.Stdout, menu)
    if err != nil {
        log.Fatalln(err)
    }
}
