package main

import (
	"fmt"
	"html"
)

func Html() {
	str := `Hello: <world> a=b&c=d?e=f`
	en := html.EscapeString(str)
	fmt.Println(en)
	fmt.Println(html.UnescapeString(en))
}

func main() {
	Html()
}
