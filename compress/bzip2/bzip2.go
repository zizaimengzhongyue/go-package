package main

import (
	"compress/bzip2"
	"fmt"
	"io/ioutil"
	"os"
)

func Bzip2() {
	f, err := os.Open("./bzip2.go.bz2")
	if err != nil {
		panic(err)
	}
	reader := bzip2.NewReader(f)
	bts, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bts))
}

func main() {
	Bzip2()
}
