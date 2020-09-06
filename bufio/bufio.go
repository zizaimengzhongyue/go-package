package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func Reader() {
	f, err := os.OpenFile("./bufio.go", os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	buf := bufio.NewReader(f)
	fmt.Printf("初始 bufio.Reader 缓存内容大小为: %d\n", buf.Buffered())
	bts := make([]byte, 32)
	_, err = buf.Read(bts)
	if err != nil {
		panic(err)
	}
	fmt.Printf("bufio.Reader.Read: %s\n", string(bts))
	fmt.Printf("读取之后 bufio.Reader 缓存内容大小为: %d\n", buf.Buffered())
	bts = make([]byte, 8192)
	_, err = buf.Read(bts)
	if err != nil {
		panic(err)
	}
	fmt.Printf("读取 8192 字节之后 bufio.Reader 缓存内容大小为: %d\n", buf.Buffered())
}

func Writer() {
	f, err := os.OpenFile("./tmp", os.O_RDONLY|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := bufio.NewWriter(f)
	fmt.Printf("初始 bufio.Writer 的大小为: %d\n", buf.Size())
	fmt.Printf("初始 bufio.Writer 可用空间大小为: %d\n", buf.Available())
	fmt.Printf("初始 bufio.Writer 的缓存大小为: %d\n", buf.Buffered())
	_, err = buf.Write([]byte("hello,world"))
	fmt.Printf("写入 `hello,world` 之后 bufio.Writer 的可用空间大小为: %d\n", buf.Available())
	fmt.Printf("写入 `hello,world` 之后 bufio.Writer 的缓存大小为: %d\n", buf.Buffered())
	bts, err := ioutil.ReadFile("./tmp")
	if err != nil {
		panic(err)
	}
	fmt.Printf("此时文件中的内容为: %s\n", string(bts))
	err = buf.Flush()
	if err != nil {
		panic(err)
	}
	bts, err = ioutil.ReadFile("./tmp")
	fmt.Printf("bufio.Flush 之后文件中的内容为: %s\n", string(bts))
	err = buf.Flush()
	if err != nil {
		panic(err)
	}
	fmt.Printf("bufio.Flush bufio.Writer 的大小为: %d\n", buf.Size())
	fmt.Printf("bufio.Flush bufio.Writer 可用空间大小为: %d\n", buf.Available())
	fmt.Printf("bufio.Flush bufio.Writer 的缓存大小为: %d\n", buf.Buffered())
	buf.WriteString("hello,world")
	tmp01, err := os.OpenFile("./tmp01", os.O_RDONLY|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	buf.Reset(tmp01)
	if err != nil {
		panic(err)
	}
	fmt.Printf("bufio.Reset bufio.Writer 的大小为: %d\n", buf.Size())
	fmt.Printf("bufio.Reset bufio.Writer 可用空间大小为: %d\n", buf.Available())
	fmt.Printf("bufio.Reset bufio.Writer 的缓存大小为: %d\n", buf.Buffered())
}

func main() {
	Reader()
	fmt.Println("\n")
	Writer()
}

//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
//这段注释是为了让文件长度超过 bufio 默认的 buffer 大小
