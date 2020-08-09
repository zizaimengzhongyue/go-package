package main

import (
	"fmt"
	"sync/atomic"
)

type Node struct {
	Key   string
	Value string
}

func StoreAndLoad() {
	a := 9
	val := atomic.Value{}
	val.Store(a)
	fmt.Printf("Value.Store.Int: %+v\n", val.Load())
	b := Node{Key: "key1", Value: "value1"}
	val = atomic.Value{}
	val.Store(b)
	fmt.Printf("Value.Store.Struct: %+v\n", val.Load())
	c := []int{1, 2, 3}
	val = atomic.Value{}
	val.Store(c)
	fmt.Printf("Value.Store.Slice: %+v\n", val.Load())
	d := map[string]string{"key1": "value1", "key2": "value2"}
	val = atomic.Value{}
	val.Store(d)
	fmt.Printf("Value.Store.Map: %+v\n", val.Load())
}

func main() {
	StoreAndLoad()
}
