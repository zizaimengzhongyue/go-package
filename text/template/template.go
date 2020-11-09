package main

import (
	"fmt"
	"text/template"
)

func Template() {
	s := "<b>hello,world</b>"
	fmt.Println(template.HTMLEscapeString(s))

	t, o := template.IsTrue("hello")
	fmt.Println(t, o)
	t, o = template.IsTrue("")
	fmt.Println(t, o)

	j := `<script>var a = "hello" </script>`
	fmt.Println(template.JSEscapeString(j))

	u := `http://www.baidu.com?k=v`
	fmt.Println(template.URLQueryEscaper(u))
}

func main() {
	Template()
}
