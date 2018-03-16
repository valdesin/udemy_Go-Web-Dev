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

type car struct {
    Manufacturer string
    Model string
    Doors int
}

func init() {
    tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
    b := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	g := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	m := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}
    
    sages := []sage{b, g, m}
    cars := []car{f, c}
    data := struct {
        Wisdom []sage
		Transport []car
    }{
		Wisdom: sages,
		Transport: cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
