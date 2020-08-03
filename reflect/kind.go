package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 10
	kind := reflect.TypeOf(a).Kind()
	fmt.Printf("Kind 是一个 unint 类型的数字，只有 String 一个方法，输出该类型的字符串名称\n")
	fmt.Printf("    Kind.String: %s\n", kind)
}
