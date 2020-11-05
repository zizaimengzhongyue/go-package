package main

import (
	"fmt"
	"strconv"
)

func Strconv() {
	b := []byte("hello,world")
	fmt.Println(string(strconv.AppendBool(b, true)))
	fmt.Println(string(strconv.AppendInt(b, 108, 10)))
	fmt.Println(string(strconv.AppendQuote(b, "world,hello")))

	a, err := strconv.Atoi("19")
	if err != nil {
		panic(err)
	}
	fmt.Println(a)

	s := strconv.Itoa(19)
	fmt.Println(s)

	fmt.Println(strconv.FormatBool(true))

	s = "true"
	bl, err := strconv.ParseBool(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(bl)
}

func main() {
	Strconv()
}
