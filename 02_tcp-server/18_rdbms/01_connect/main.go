package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "valdesin:vlad1994033114@tcp(mydbinstance.cqureocv9nu0.us-east-2.rds.amazonaws.com:3306)/test01?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err = http.ListenAndServe(":8080", nil)
	check(err)
}

func index(res http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(res, "Successfully complete")
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
