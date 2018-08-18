package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *mysql.DB
var err error

func main() {
	db, err = mysql.Open("mysql", "valdesin:vlad1994033114@tcp(mydbinstance.cqureocv9nu0.us-east-2.rds.amazonaws.com:3306)/01_test?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()

	http.HandleFunc("/", index)
	http.HandleFunc("/amigos", amigos)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/del", del)
	http.HandleFunc("/drop", drop)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err = http.ListenAndServe(":8080", nil)
	check(err)

}

func index(res http.ResponseWriter, req *http.Request) {}

func amigos(res http.ResponseWriter, req *http.Request) {}

func create(res http.ResponseWriter, req *http.Request) {}

func insert(res http.ResponseWriter, req *http.Request) {}

func read(res http.ResponseWriter, req *http.Request) {}

func update(res http.ResponseWriter, req *http.Request) {}

func del(res http.ResponseWriter, req *http.Request) {}

func drop(res http.ResponseWriter, req *http.Request) {}

func check(err error) {
	if err != nill {
		fmt.Println(err)
	}
}
