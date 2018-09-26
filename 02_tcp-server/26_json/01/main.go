package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	FName string
	LName string
	Items []string
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/msh", msh)
	http.HandleFunc("/encod", encod)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":9090", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	s := `<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>FOO</title>
	</head>
	<body>
	You are at foo
	</body>
	</html>`
	res.Write([]byte(s))
}

func msh(res http.ResponseWriter, req *http.Request) {
	p1 := person{
		FName: "Vlad",
		LName: "Fill",
		Items: []string{"Gun", "ShotGun", "Rifle"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte(bs))
}

func encod(res http.ResponseWriter, req *http.Request) {
	p1 := person{
		FName: "Robert",
		LName: "Rodriges",
		Items: []string{"Pistol", "bat", "M16"},
	}

	res.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(res).Encode(p1)
	if err != nil {
		log.Println(err)
	}
}
