package main

import (
	"log"
	"os"
)

func Log() {
	logger := log.New(os.Stdout, "", log.Ldate)
	logger.Println("hello,world")
	logger.Panic("panic")
}

func main() {
	Log()
}
