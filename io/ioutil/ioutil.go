package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ReadAll() {
	fmt.Println("ioutil.ReadAll 读取 io.Reader 全部内容")
	str := "hello,world"
	bts, _ := ioutil.ReadAll(strings.NewReader(str))
	fmt.Printf("ioutil.ReadAll: %s\n", string(bts))
}

func ReadFile() {
	fmt.Println("ioutil.ReadFile 读取文件内容")
	fmt.Println("我读我自己\n")
	bts, _ := ioutil.ReadFile("./ioutil.go")
	fmt.Printf("ioutil.ReadFile: %s\n", string(bts))
}

func WriteFile() {
	fmt.Println("ioutil.WriteFile 写入文件内容")
	fmt.Println("我写我自己\n")
	bts, _ := ioutil.ReadFile("./ioutil.go")
	bts = append(bts, []byte("\n//到此一游\n")...)
	_ = ioutil.WriteFile("./ioutil.go", bts, 0755)
}

func ReadDir() {
	fmt.Println("ioutil.ReadDir 读取当前文件夹，返回全部 os.FileInfo")
	fmt.Printf("我之前居然不知道有这个方法，sad\n")
	files, _ := ioutil.ReadDir("../../io")
	for _, v := range files {
		fmt.Println(v.Name())
	}
}

func Discard() {
	fmt.Println("ioutil.Discard 可以用来随意写入且不会报错，写入的内容直接丢弃")
	str := "hello,world"
	n, err := ioutil.Discard.Write([]byte(str))
	fmt.Printf("ioutil.Discard.Write 返回值: %d, %+v\n", n, err)
}

func main() {
	ReadAll()
	ReadFile()
	WriteFile()
	ReadDir()
	Discard()
}

//到此一游

//到此一游

//到此一游

//到此一游

//到此一游
