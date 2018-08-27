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

	http.HandleFunc("/", index)
	http.HandleFunc("/amigos", amigos)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/del", del)
	http.HandleFunc("/drop", drop)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err = http.ListenAndServe(":80", nil)
	check(err)

}

func index(res http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(res, "at index")
	check(err)
}

func amigos(res http.ResponseWriter, req *http.Request) {
	row, err := db.Query("SELECT Name FROM amigos;")
	check(err)
	defer row.Close()

	var s, name string
	s = "RETRIVED RECORDS:\n"

	for row.Next() {
		err := row.Scan(&name)
		check(err)
		s += name + "\n"
	}
	fmt.Fprintln(res, s)
}

func create(res http.ResponseWriter, req *http.Request) {
	s, err := db.Prepare("CREATE TABLE customer (name VARCHAR(20));")
	check(err)
	defer s.Close()

	r, err := s.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintf(res, "CTRAETED TABLE customer", n)
}

func insert(res http.ResponseWriter, req *http.Request) {
	stm, err := db.Prepare(`INSERT INTO customer VALUES ("Vlad");`)
	check(err)
	defer stm.Close()

	r, err := stm.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(res, "INSERTED RECORD: ", n)
}

func read(res http.ResponseWriter, req *http.Request) {
	row, err := db.Query(`SELECT * FROM customer;`)
	check(err)
	defer row.Close()

	var name string
	for row.Next() {
		err = row.Scan(&name)
		check(err)
		fmt.Fprintln(res, "RETRIEVED RECORD:", name)
	}
}

func update(res http.ResponseWriter, req *http.Request) {
	stm, err := db.Prepare(`UPDATE customer SET name="Fill" WHERE name="Vlad";`)
	check(err)
	defer stm.Close()

	r, err := stm.Exec()
	check(err)

	n, err := r.RowsAffected()

	fmt.Fprintln(res, "UPDATED RECORD: ", n)

}

func del(res http.ResponseWriter, req *http.Request) {
	stm, err := db.Prepare(`DELETE FROM customer WHERE name="Fill";`)
	check(err)
	defer stm.Close()

	r, err := stm.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(res, "DELETE RECORD: ", n)
}

func drop(res http.ResponseWriter, req *http.Request) {
	stm, err := db.Prepare(`DROP TABLE customer;`)
	check(err)
	defer stm.Close()

	_, err = stm.Exec()
	check(err)

	fmt.Fprintln(res, "DROPED TABLE customer")
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
