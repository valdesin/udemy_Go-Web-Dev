package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	c1 := getCode("vlad@gmail.com")
	fmt.Println(c1)
	c2 := getCode("vlad@gmail.com")
	fmt.Println(c2)

	fmt.Println(compareMAC(c1, c2))
}

func getCode(s string) string {
	h := hmac.New(sha256.New, []byte("my-password"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func compareMAC(m1, m2 string) bool {
	if !hmac.Equal([]byte(m1), []byte(m2)) {
		return false
	}
	return true
}
