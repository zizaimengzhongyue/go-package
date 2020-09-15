package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	ID   int
	Name string
}

func JSON() {
	str := `{"id": 1, "name": "张三"}`
	p := &Person{}
	err := json.Unmarshal([]byte(str), p)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", p)
	bts, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bts))
}

func main() {
	JSON()
}
