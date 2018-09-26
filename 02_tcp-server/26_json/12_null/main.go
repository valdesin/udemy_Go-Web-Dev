package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var n []string
	dJ := `null`
	err := json.Unmarshal([]byte(dJ), &n)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(n)
	fmt.Println(len(n))
	fmt.Println(cap(n))
}
