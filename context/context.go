package main

import (
	"context"
	"fmt"
	"time"
)

func ContextWithCancel() {
	done := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		}
		done <- true
	}()
	cancel()
	<-done
}

func ContextWithTimeout() {
	done := make(chan bool)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("goroutine01: ", ctx.Err())
		case <-time.After(3 * time.Second):
			fmt.Println("goroutine01: context exceeded")
		}
		done <- true
	}()
	select {
	case <-ctx.Done():
		fmt.Println("main gouroutine: context finished")
	}
	<-done
}

func ContextWithValue() {
	done := make(chan bool)
	ctx := context.WithValue(context.Background(), "hello", "world")
	go func() {
		fmt.Println(ctx.Value("hello"))
		done <- true
	}()
	<-done
}

func main() {
	ContextWithCancel()
	ContextWithTimeout()
	ContextWithValue()
}
