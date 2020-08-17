package main

import (
	"fmt"
	"sync"
	"time"
)

func Hello() {
	fmt.Println("Hello")
}

func Do() {
	once := sync.Once{}
	for i := 0; i < 10000; i++ {
		go once.Do(Hello)
	}
	time.Sleep(time.Second)
}

func main() {
	Do()
}
