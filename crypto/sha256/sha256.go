package main

import (
	"crypto/sha256"
	"fmt"
)

func Sha256() {
	sha := sha256.New()
	bts := sha.Sum([]byte("hello,world"))
	fmt.Printf("%x\n", bts)
}

func main() {
	Sha256()
}
