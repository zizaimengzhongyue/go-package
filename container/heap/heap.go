package main

import (
	"container/heap"
	"fmt"
)

type Heap struct {
	List []int
}

func (this *Heap) Len() int {
	return len(this.List)
}

func (this *Heap) Less(i, j int) bool {
	return this.List[i] > this.List[j]
}

func (this *Heap) Swap(i, j int) {
	this.List[i], this.List[j] = this.List[j], this.List[i]
}

func (this *Heap) Push(x interface{}) {
	this.List = append(this.List, x.(int))
}

func (this *Heap) Pop() interface{} {
	ln := this.Len()
	ans := this.List[ln-1]
	this.List = this.List[0 : ln-1]
	return ans
}

func HeapF() {
	h := &Heap{List: []int{3, 5, 4, 2, 1}}
	heap.Init(h)
	heap.Push(h, 20)
	heap.Push(h, 19)
	ln := h.Len()
	for i := 0; i < ln; i++ {
		fmt.Println(heap.Pop(h))
	}
}

func main() {
	HeapF()
}
