package main

import (
	"encoding/json"
	"fmt"
)

type code struct {
	Code     int
	Descript string
}

func main() {
	var sCode []code

	rcvd := `[{"Code":200,"Descript":"StatusOK"},{"Code":301,"Descript":"StatusMovedPermanently"},{"Code":302,"Descript":"StatusFound"},{"Code":303,"Descript":"StatusSeeOther"},{"Code":307,"Descript":"StatusTemporaryRedirect"},{"Code":400,"Descript":"StatusBadRequest"},{"Code":401,"Descript":"StatusUnauthorized"},{"Code":402,"Descript":"StatusPaymentRequired"},{"Code":403,"Descript":"StatusForbidden"},{"Code":404,"Descript":"StatusNotFound"},{"Code":405,"Descript":"StatusMethodNotAllowed"},{"Code":418,"Descript":"StatusTeapot"},{"Code":500,"Descript":"StatusInternalServerError"}]`
	err := json.Unmarshal([]byte(rcvd), &sCode)
	if err != nil {
		fmt.Println(err)
	}

	for _, data := range sCode {
		fmt.Printf("%d --- %s\n", data.Code, data.Descript)
	}
}
