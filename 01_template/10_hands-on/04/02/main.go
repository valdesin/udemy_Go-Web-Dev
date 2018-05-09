package main

import (
    "os"
    "log"
    "text/template"
)

type menu struct {
    Name string
    Price float64
}

type meat struct {
    Kind string
    Menu []menu
}

type restaurant struct {
    Name string
    Meat []meat
}

var tpl *template.Template

func init() {
    tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
    
    rsts := []restaurant {
        restaurant {
            Name: "Chine WOK",
            Meat: []meat{
                meat{                
                    Kind: "breakfast",
                    Menu: []menu{
                        menu{
                            Name:"Two Fresh Eggs",
                            Price: 3.99,
                        },
                        menu{
                            Name:"Western Omelet",
                            Price: 7.39,
                        },
                    },
                },
                meat{                
                    Kind: "lunch",
                    Menu: []menu{
                        menu{
                            Name:"Spaghetti",
                            Price: 2.95,
                        },
                        menu{
                            Name:"Indian Taco",
                            Price: 6.99,
                        },
                    },
                },
                meat{                
                    Kind: "dinner",
                    Menu: []menu{
                        menu{
                            Name:"Chicken Fried Steak",
                            Price: 8.99,
                        },
                        menu{
                            Name:"Fried Shrimp Basket",
                            Price: 11.99,
                        },
                    },
                },
            },
        },
        restaurant {
            Name: "MULLET'S",
            Meat: []meat{
                meat{                
                    Kind: "breakfast",
                    Menu: []menu{
                        menu{
                            Name:"Two Fresh Eggs",
                            Price: 3.99,
                        },
                        menu{
                            Name:"Fish & Chips",
                            Price: 7.39,
                        },
                    },
                },
                meat{                
                    Kind: "lunch",
                    Menu: []menu{
                        menu{
                            Name:"Pulled Pork Salad",
                            Price: 10.95,
                        },
                        menu{
                            Name:"Home Soup Of The Day",
                            Price: 4.99,
                        },
                    },
                },
                meat{                
                    Kind: "dinner",
                    Menu: []menu{
                        menu{
                            Name:"Fried Catfish",
                            Price: 10.99,
                        },
                        menu{
                            Name:"Fried Shrimp Basket",
                            Price: 11.99,
                        },
                    },
                },
            },
        },

    }

    err := tpl.Execute(os.Stdout, rsts)
    if err != nil {
        log.Fatalln(err)
    }
}
