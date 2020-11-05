package main

import (
	"fmt"
	"plugin"
)

type P interface {
	Ping() string
}

func Plugin() {
	p, err := plugin.Open("./ping/ping.so")
	if err != nil {
		panic(err)
	}

	i, err := p.Lookup("I")
	if err != nil {
		panic(err)
	}

	im, ok := i.(P)
	if !ok {
		panic(p)
	}

	fmt.Println(im.Ping())
}

func main() {
	Plugin()
}
