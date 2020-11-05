package main

import (
	"fmt"
	"sort"
)

type Node struct {
	Nums []int
}

func (n *Node) Len() int {
	return len(n.Nums)
}

func (n *Node) Less(i int, j int) bool {
	return n.Nums[i] < n.Nums[j]
}

func (n *Node) Swap(i int, j int) {
	n.Nums[i], n.Nums[j] = n.Nums[j], n.Nums[i]
}

func Sort() {
	f := []float64{1.00, 1.00000000003, 1.0000000002}
	sort.Float64s(f)
	fmt.Println(f)

	a := []int{3, 1, 2, 5, 4}
	sort.Ints(a)
	fmt.Println(a)
}

func Interface() {
	n := &Node{Nums: []int{3, 2, 1, 4, 5}}
	sort.Sort(n)
	fmt.Println(n.Nums)
}

func main() {
	Sort()
	Interface()
}
