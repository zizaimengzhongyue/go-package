package main

import (
	"fmt"
	"hash/crc32"
)

func Crc32() {
	fmt.Println("没太明白原理，抄了一个 demo")
	crc := crc32.MakeTable(0xD5828281)
	fmt.Printf("%08x\n", crc32.Checksum([]byte("hello,world"), crc))
}

func main() {
	Crc32()
}
