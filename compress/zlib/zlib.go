package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"
)

var buf *bytes.Buffer = bytes.NewBuffer(nil)

func Writer() {
	wr := zlib.NewWriter(buf)
	wr.Write([]byte("hello,world"))
	wr.Close()
	fmt.Println(buf.String())
}

func Reader() {
	rd, err := zlib.NewReader(buf)
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, rd)
}

func main() {
	Writer()
	Reader()
}
