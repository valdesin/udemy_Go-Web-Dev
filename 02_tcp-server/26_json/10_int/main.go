package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var a int
	d := `43`
	err := json.Unmarshal([]byte(d), &a)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(a)
}
