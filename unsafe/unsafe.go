package main

import (
	"fmt"
	"unsafe"
)

type Node struct {
	Key   string
	Value string
	Index int64
}

type Byte struct {
	K1 byte
}

type Bool struct {
	K1 bool
}

type Int64 struct {
	K1 int
}

type Complex struct {
	K1 byte
	K2 bool
	K3 int
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
	fmt.Printf("通过 unsafe.Offsetof 可以看到 struct.field 的偏移位置\n")
	node := Node{Key: "key123123", Value: "value"}
	fmt.Printf("unsafe.Offsetof: %+v\n", unsafe.Offsetof(node.Value))
}

func Alignof() {
	k1 := Byte{K1: 'a'}
	fmt.Printf("Byte.Alignof: %d, Byte.Sizeof: %d\n", unsafe.Alignof(k1), unsafe.Sizeof(k1))
	k2 := Bool{K1: true}
	fmt.Printf("Bool.Alignof: %d, Bool.Sizeof: %d\n", unsafe.Alignof(k2), unsafe.Sizeof(k2))
	k3 := Int64{K1: 1}
	fmt.Printf("Int.Alignof: %d, Int.Sizeof: %d\n", unsafe.Alignof(k3), unsafe.Sizeof(k3))
	k4 := Complex{K1: 'a', K2: true, K3: 8}
	fmt.Printf("Struct.Aliginof: %d, Struct.Sizeof: %d\n", unsafe.Alignof(k4), unsafe.Sizeof(k4))
}

func Slice() {
	a := make([]int, 3, 10)
	a[0] = 1
	a[1] = 2
	ptr := unsafe.Pointer(&a)
	base := unsafe.Pointer(&a[0])
	fmt.Printf("Slice.len: %d\n", *(*int)(unsafe.Pointer(uintptr(ptr) + uintptr(8))))
	fmt.Printf("Slice.cap: %d\n", *(*int)(unsafe.Pointer(uintptr(ptr) + uintptr(16))))
	fmt.Printf("Slice.Index.0: %d\n", *(*int)(unsafe.Pointer(uintptr(base))))
	fmt.Printf("Slice.Index.1: %d\n", *(*int)(unsafe.Pointer(uintptr(base) + uintptr(8))))
	fmt.Printf("Slice.Index.2: %d\n", *(*int)(unsafe.Pointer(uintptr(base) + uintptr(16))))
}

func main() {
	Convert()
	Offset()
	Alignof()
	Slice()
}
