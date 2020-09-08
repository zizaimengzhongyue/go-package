package main

import (
	"container/ring"
	"fmt"
)

func Ring() {
	rn := ring.New(10)
	for i := 0; i < rn.Len(); i++ {
		rn.Value = i
		rn = rn.Next()
	}
	for i := 0; i < rn.Len(); i++ {
		fmt.Println(rn.Value)
		rn = rn.Prev()
	}
}

func main() {
	Ring()
}
