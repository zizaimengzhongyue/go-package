package main

import (
	"fmt"
	"sync"
	"testing"
)

type Node struct {
	Val int
	ID  int
}

func New() {
	fmt.Printf("sync.Pool 维护一个资源池，对资源进行复用\n")

	nodeMaker := &sync.Pool{New: func() interface{} { return &Node{} }}
	node, _ := nodeMaker.Get().(*Node)
	node.Val = 5
	node.ID = 4
	fmt.Printf("sync.Pool.Get: %+v\n", node)
	nodeMaker.Put(node)
	node, _ = nodeMaker.Get().(*Node)
	fmt.Printf("sync.Pool.Get: %+v\n", node)

	nMaker := &sync.Pool{New: func() interface{} { return Node{} }}
	n, _ := nMaker.Get().(Node)
	n.Val = 5
	n.ID = 4
	fmt.Printf("sync.Pool.New: %+v\n", n)
	fmt.Printf("sync.Pool.Addr: %p\n", &n)
	nMaker.Put(node)
	n2, _ := nMaker.Get().(Node)
	fmt.Printf("sync.Pool.Get: %+v\n", n2)
	fmt.Printf("sync.Pool.Addr: %p\n", &n2)
}

func main() {
	New()
}
