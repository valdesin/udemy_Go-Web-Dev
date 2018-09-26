package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var name string
	cvds := `"Vlad"`
	err := json.Unmarshal([]byte(cvds), &name)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(name)
}
