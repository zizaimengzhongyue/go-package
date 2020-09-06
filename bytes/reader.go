package main

import (
	"bytes"
	"fmt"
)

func Info() {
	fmt.Printf("Bytes.Reader 针对 bytes 实现了 io.Reader 这个 interface\n")
}

func Reader() {
	reader := bytes.NewReader([]byte("hello,world"))
	fmt.Printf("Reader.Len: %d\n", reader.Len())
	fmt.Printf("Reader.Size: %d\n", reader.Size())
}

func main() {
	Info()
	Reader()
}
