package main

import (
	"bytes"
	"fmt"
)

func Buffer() {
	buf := bytes.NewBuffer([]byte("hello,world"))
	fmt.Printf("Buffer.Bytes: %s\n", string(buf.Bytes()))
	fmt.Printf("Buffer.Cap: %d\n", buf.Cap())
	fmt.Printf("Buffer.Len: %d\n", buf.Len())
}

func main() {
	Buffer()
}
