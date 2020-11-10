package main

import (
	"fmt"
	"unicode/utf8"
)

func UTF8() {
	a := '好'
	b := make([]byte, 4)
	fmt.Println(utf8.EncodeRune(b, a))
	fmt.Println(b)

	s := "你好，hello"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
	fmt.Println(utf8.RuneCount([]byte(s)))
}

func main() {
	UTF8()
}
