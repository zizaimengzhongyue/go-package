package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"text/scanner"
)

func Scanner() {
	bts, err := ioutil.ReadFile("./scanner.go")
	if err != nil {
		panic(err)
	}

	s := &scanner.Scanner{}
	s.Init(bytes.NewReader(bts))
	for t := s.Scan(); t != scanner.EOF; t = s.Scan() {
		fmt.Println(s.Position, s.TokenText())
	}
}

func main() {
	Scanner()
}
