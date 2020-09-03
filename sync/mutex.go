package main

import (
	"fmt"
	"sync"
)

func Increase() {
	a := 0
	wg := &sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			a++
		}()
	}
	wg.Wait()
	fmt.Printf("预期得到的值是 1000，实际得到的值是 %d\n", a)
}

func Mutex() {
	fmt.Println("mutex 是个互斥锁，可以通过 mutex.Lock 和 mutex.Unlock 保证临界区的代码同一时间只有一个 goroutine 在执行")
	mutex := &sync.Mutex{}
	a := 0
	wg := &sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutex.Lock()
			defer mutex.Unlock()
			a++
		}()
	}
	wg.Wait()
	fmt.Printf("预期得到的值是 10000，实际得到的值是 %d\n", a)
}

func main() {
	Increase()
	Mutex()
}
