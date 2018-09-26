package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}

func StoreUsers(m map[string]User) {
	f, err := os.Create("data")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	json.NewEncoder(f).Encode(m)
}

func LoadUsers() map[string]User {
	m := make(map[string]User)
	f, err := os.Open("data")
	if err != nil {
		fmt.Println(err)
		return m
	}
	defer f.Close()

	json.NewDecoder(f).Decode(&m)

	return m
}
