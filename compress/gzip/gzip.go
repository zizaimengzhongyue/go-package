package main

import (
	"compress/gzip"
	"fmt"
	"os"
)

func Writer() {
	f, err := os.OpenFile("./tmp01", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	wr := gzip.NewWriter(f)
	defer wr.Close()

	_, err = wr.Write([]byte("hello,world"))
	if err != nil {
		panic(err)
	}
	wr.Flush()

	fmt.Println("finished")
}

func Reader() {
	f, err := os.OpenFile("./tmp01", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	rd, err := gzip.NewReader(f)
	if err != nil {
		panic(err)
	}

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
