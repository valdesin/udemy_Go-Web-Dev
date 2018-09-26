package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var b bool
	d := `true`
	err := json.Unmarshal([]byte(d), &b)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(b)
}
