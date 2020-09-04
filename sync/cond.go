package main

import (
	"fmt"
	"sync"
	"time"
)

func Cond() {
	condTyp := `
    type Cond struct {
    noCopy noCopy

    // L is held while observing or changing the condition
    L Locker

    notify  notifyList
    checker copyChecker
}`
	wait := `
func (c *Cond) Wait() {
    c.checker.check()
    t := runtime_notifyListAdd(&c.notify)
    c.L.Unlock()
    runtime_notifyListWait(&c.notify, t)
    c.L.Lock()
}
`
	fmt.Println(condTyp)
	fmt.Printf("cond.Wait 会释放锁，Wait 具体代码:\n")
	fmt.Println(wait)
	cond := &sync.Cond{L: &sync.Mutex{}}
	a := 0
	go func() {
		for {
			cond.L.Lock()
			for a >= 10 {
				cond.Wait()
			}
			a++
			//fmt.Println(a)
			cond.Signal()
			cond.L.Unlock()
		}
	}()
	go func() {
		for {
			cond.L.Lock()
			for a <= 0 {
				cond.Wait()
			}
			a--
			//fmt.Println(a)
			cond.Signal()
			cond.L.Unlock()
		}
	}()
	time.Sleep(10 * time.Second)
	fmt.Println("finished")
}

func main() {
	Cond()
}
