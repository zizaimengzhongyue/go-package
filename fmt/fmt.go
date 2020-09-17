package main

import (
	"errors"
	"fmt"
	"os"
)

func Fmt() {
	err := errors.New("test")
	erf := fmt.Errorf("wrappers: %w", err)
	fmt.Println(erf)
	fmt.Fprintf(os.Stdout, "hello,world\n")
	fmt.Fprintln(os.Stdout, "world,hello")
	var a int
	fmt.Scanf("%d", &a)
	fmt.Println(a)
	str := "hello: %s"
	fmt.Println(fmt.Sprintf(str, "张三"))
	var name string
	var score int
	fmt.Sscan("张三 35", &name, &score)
	fmt.Println(name, score)
}

func main() {
	Fmt()
}
