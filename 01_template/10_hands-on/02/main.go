package main

import (
    "os"
    "log"
    "text/template"
)

type hotel struct {
    Name, Address, City, Zip, Region string
}

var tpl *template.Template

func init() {
    tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
    hotels := []hotel {
        hotel {
            Name: "McCloud Hotel",
            Address: "408 Main Street",
            City: "McCloud",
            Zip: "96057",
            Region: "Northern",
        },

        hotel {
                Name: "Pelican Inn & Suites",
                Address: "6316 Moonstone Beach Dr",
                City: "Cambria",
                Zip: "93428-1806",
                Region: "Central",
         },

        hotel {
            Name: "Manchester Grand Hyatt San Diego",
            Address: "One Market Place",
            City: "San Diego",
            Zip: "92101-7714",
            Region: "Southern",
        },
    }

    err := tpl.Execute(os.Stdout, hotels)
    if err != nil {
        log.Fatalln(err)
    }
}
