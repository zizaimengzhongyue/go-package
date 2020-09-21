package main

import (
	"fmt"
	"math"
)

func Math() {
	a := 1.23
	b := 1.2300000000001
	fmt.Println(math.Max(a, b))
	a = 1.25
	b = 1.23
	fmt.Println(math.Mod(a, b))
}

func main() {
	Math()
}
