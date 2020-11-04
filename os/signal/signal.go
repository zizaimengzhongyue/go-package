package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Ignore() {
	fmt.Println("SIGKILL 和 SIGSTOP 无法被拦截")
	signal.Ignore(syscall.SIGKILL)
	fmt.Println("kill -9 还是可以杀死该进程")
	signal.Ignore(syscall.SIGINT)
	fmt.Println("SIGINT 信号被忽略，ctrl + c 被无法终止该进程")
}

func Notify() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGFPE)
	go func(ch chan os.Signal) {
		for {
			_ = <-ch
			fmt.Println("我捕获了一次 SIGFPE 信号")
		}
	}(ch)
	fmt.Println("SIGFPE 信号被捕获")
}

func Do() {
	for {
	}
}

func main() {
	Ignore()
	Notify()
	Do()
	signal.Reset()
}
