package main

import (
	"fmt"
	"os"
	"strings"
)

func Strings() {
	a := "hello"
	b := "hello,world"
	fmt.Println(strings.Compare(a, b))
	fmt.Println(strings.Contains(b, a))
	fmt.Println(strings.Index(b, a))
	fmt.Println(strings.ReplaceAll(b, "world", "new world"))

	fmt.Println(strings.ToLower("Hello"))
	fmt.Println(strings.ToUpper("Hello"))
	fmt.Println(strings.Trim("  hello,world  ", " "))
}

func Replacer() {
	r := strings.NewReplacer("hello", "world")
	fmt.Println(r.Replace("hello,world"))
	fmt.Println(r.Replace("world"))
	_, err := r.WriteString(os.Stdout, "hehe")
	if err != nil {
		panic(err)
	}
}

func main() {
	Strings()
	Replacer()
}
