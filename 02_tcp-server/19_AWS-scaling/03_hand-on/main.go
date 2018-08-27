package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "valdesin:vlad1994033114@tcp(mydbinstance.cqureocv9nu0.us-east-2.rds.amazonaws.com:3306)/test01?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/instance", instance)
	http.HandleFunc("/amigos", amigos)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":80", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello from AWS")
}

func ping(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "OK")
}

func instance(res http.ResponseWriter, req *http.Request) {
	s := getInstantce()
	io.WriteString(res, s)
}

func amigos(res http.ResponseWriter, req *http.Request) {
	row, err := db.Query("SELECT Name FROM amigos;")
	check(err)

	s := getInstantce()
	s += "\nRETRIEVED RECORDS:\n"

	var name string
	for row.Next() {
		err = row.Scan(&name)
		check(err)
		s += name + "\n"
	}

	fmt.Fprintln(res, s)
}

func getInstantce() string {
	resp, err := http.Get("http://169.254.169.254/latest/meta-data/instance-id")
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	bs := make([]byte, resp.ContentLength)
	resp.Body.Read(bs)
	resp.Body.Close()

	return string(bs)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
