package main

import (
	"bytes"
	"fmt"
)

func Bytes() {
	bts := []byte("hello,world")
	sub := []byte("hello")
	fmt.Printf("bytes.Count: %d\n", bytes.Count(bts, sub))
	fmt.Printf("bytes.Contains: %t\n", bytes.Contains(bts, sub))
	fmt.Printf("bytes.ContainsAny: %t\n", bytes.ContainsAny(bts, string(sub)))
	fmt.Printf("bytes.IndexByte: %d\n", bytes.IndexByte(bts, 'o'))
	fmt.Printf("bytes.LastIndex: %d\n", bytes.LastIndex(bts, []byte("o")))
	fmt.Printf("bytes.SplitN: %+v\n", bytes.SplitN(bts, []byte(","), -1))
	fmt.Printf("bytes.Split: %+v\n", bytes.Split(bts, []byte(",")))
	fmt.Printf("bytes 还有 Join,HasPrefix,HasSuffix,Map,Repeat,ToUpper,ToLower 等常用方法，任何针对单个 byte 的操作都可以先看下 bytes 库是否有已经封装好的方法\n")
}

func main() {
	Bytes()
}
