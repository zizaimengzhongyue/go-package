package main

import (
	"fmt"
	"os"
	"path"
)

func Path() {
	p, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(p)
	fmt.Println(path.Base(p))
	fmt.Println(path.Clean(p))
	fmt.Println(path.Dir(p))
	fmt.Println(path.Ext(p))
	fmt.Println(path.IsAbs(p))
	bol, err := path.Match(p, "*go-package*")
	if err != nil {
		panic(err)
	}
	fmt.Println(bol)
	dir, f := path.Split(p)
	fmt.Println(dir, f)
}

func main() {
	Path()
}
