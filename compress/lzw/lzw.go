package main

import (
	"bytes"
	"compress/lzw"
	"fmt"
	"io"
	"os"
)

var buf *bytes.Buffer = bytes.NewBuffer(nil)

func Writer() {
	wr := lzw.NewWriter(buf, lzw.LSB, 8)
	_, err := wr.Write([]byte("hello,world"))
	if err != nil {
		panic(err)
	}
	wr.Close()
	fmt.Println(buf.String())
}

func Reader() {
	rd := lzw.NewReader(buf, lzw.LSB, 8)
	defer rd.Close()
	io.Copy(os.Stdout, rd) //这种输出方式太帅了
}

func main() {
	Writer()
	Reader()
}
