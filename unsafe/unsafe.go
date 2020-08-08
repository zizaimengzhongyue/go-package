package main

import (
	"fmt"
	"unsafe"
)

type Node struct {
	Key   string
	Value string
}

func Convert() {
	fmt.Printf("通过 unsafe.Pointer 可以实现任意类型指针之间的互转\n")
	a := 8
	b := *(*int64)(unsafe.Pointer(&a))
	fmt.Printf("Int.Int64: %d\n", b)
	c := 1.23
	d := *(*int)(unsafe.Pointer(&c))
	fmt.Printf("Float.Int: %d\n", d)
	e := 8
	f := *(*float32)(unsafe.Pointer(&e))
	fmt.Printf("Int.Float: %f\n", f)
	var g int64 = 10
	h := *(*float32)(unsafe.Pointer(&g))
	fmt.Printf("Int64.Float: %f\n", h)
	var i float64 = 1.23
	j := *(*int32)(unsafe.Pointer(&i))
	fmt.Printf("Float64.Int64: %d\n", j)
	node := Node{Key: "key", Value: "value"}
	n := *(*int)(unsafe.Pointer(&node))
	fmt.Printf("Node.Int: %d\n", n)
}

func Offset() {
	fmt.Printf("通过 uintptr 可以实现对指针偏移量的计算\n")
	list := []int{0, 1, 2}
	base := uintptr(unsafe.Pointer(&list[0]))
	ln := unsafe.Sizeof(list[0])
	fmt.Printf("uintptr.Offset.0: %d\n", *(*int)(unsafe.Pointer(base)))
	fmt.Printf("uintptr.Offset.1: %d\n", *(*int)(unsafe.Pointer(base + ln)))
	fmt.Printf("uintptr.Offset.2: %d\n", *(*int)(unsafe.Pointer(base + ln*2)))
}

func main() {
	Convert()
	Offset()
}
