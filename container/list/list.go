package main

import (
	"container/list"
	"fmt"
)

func List() {
	ls := list.New()
	ls.PushFront(1)
	ls.PushBack(2)
	fmt.Println(ls.Front().Value.(int), ls.Back().Value.(int))
	ls.PushFront(9)
	ls.PushFront(8)
	ls.PushFront(7)
	ls.PushFront(6)
	elem := ls.Front()
	fmt.Println(elem.Value.(int))
	elem = elem.Next()
	fmt.Println(elem.Value.(int))
	elem = elem.Next()
	fmt.Println(elem.Value.(int))
}

func main() {
	List()
}
