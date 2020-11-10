package main

import (
	"fmt"
	"unicode"
)

func Unicode() {
	s := '\u0001'
	fmt.Println(unicode.IsControl(s))

	s = '\u0035'
	fmt.Println(unicode.IsDigit(s))

	s = '\u25A0'
	fmt.Println(unicode.IsGraphic(s))

	s = '\u0046'
	fmt.Println(unicode.IsLetter(s))

	s = '\u0062'
	fmt.Println(unicode.IsLower(s))

	s = '\u2022'
	fmt.Println(unicode.IsMark(s))

	s = '\u0035'
	fmt.Println(unicode.IsNumber(s))

	// unicode 库里面对于常用字符集都写好了，可以直接判断
	s = '和'
	fmt.Println(unicode.In(s, unicode.Han))
}

func main() {
	Unicode()
}
