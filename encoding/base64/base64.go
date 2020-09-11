package main

import (
	"encoding/base64"
	"fmt"
)

func Base64() {
	str := "hello,world"
	en := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println(en)
	de, err := base64.StdEncoding.DecodeString(en)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(de))
}

func main() {
	Base64()
}
