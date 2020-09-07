package main

import (
	"bytes"
	"compress/flate"
	"fmt"
	"io"
	"os"
)

func Writer() {
	buf := bytes.NewBuffer(nil)

	wr, err := flate.NewWriter(buf, flate.BestSpeed)
	if err != nil {
		panic(err)
	}
	defer wr.Close()

	_, err = wr.Write([]byte("hello,world"))
	if err != nil {
		panic(err)
	}

	err = wr.Flush()
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile("./tmp01", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = io.Copy(f, buf)
	if err != nil {
		panic(err)
	}

	fmt.Println("finished")
}

func Reader() {
	f, err := os.Open("./tmp01")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := flate.NewReader(f)
	bts := make([]byte, 4096)
	_, err = rd.Read(bts)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bts))
}

func main() {
	Writer()
	Reader()
}
