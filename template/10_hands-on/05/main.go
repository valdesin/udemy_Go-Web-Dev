package main

import (
    "os"
    "log"
    "time"
    "net/http"
    "strconv"
    "text/template"
    "encoding/csv"
)

type Record struct {
    Date time.Time
    Open float64
}

var records []Record

func main() {
    http.HandleFunc("/", index)
    http.ListenAndServe(":9090", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
    csvData := csvReader("table.csv")

    tpl := template.Must(template.ParseFiles("tpl.gohtml"))

    err := tpl.Execute(res, csvData)
    if err != nil {
        log.Fatalln(err)
    }
    //log.Println(records)
}

func csvReader(tName string) []Record {
    src, err := os.Open(tName)
    if err != nil {
        log.Fatalln(err)
    }
    defer src.Close()
    
    csvRead := csv.NewReader(src)
    data, err  :=csvRead.ReadAll()
    if err != nil {
        log.Fatalln(err)
    }
     for i:=1; i < 10; i++ {
            //log.Println(csvData[i][0])
            date, _ := time.Parse("2006-01-02", string(data[i][0]))
            open, _ := strconv.ParseFloat(data[i][1], 64)
            r := Record {
                Date: date,
                Open: open,
            }
            records = append(records, r)
        
    }
    
    return records
}
