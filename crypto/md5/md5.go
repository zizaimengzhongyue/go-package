package main

import (
	"crypto/md5"
	"fmt"
)

func Md5() {
	fmt.Println("第一次知道 md5 的结果是十六进制数字形式")
	m := md5.New()
	bts := m.Sum([]byte("hello,world"))
	fmt.Printf("%+x\n", bts)
}

func main() {
	Md5()
}
