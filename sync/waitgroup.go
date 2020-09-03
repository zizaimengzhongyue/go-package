package main

import (
	"fmt"
	"sync"
)

func WaitGroup() {
	fmt.Printf("sync.WaitGroup 包含 Add、Done 和 Wait 三个方法，典型的使用场景是等待一批 goroutine 全部退出\n")
	wg := &sync.WaitGroup{}
	a := 0
	for i := 0; i < 10; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			a++
		}()
	}
	wg.Wait()
	fmt.Printf("WaitGroup.Wait: %d\n", a)
}

func main() {
	WaitGroup()
}
