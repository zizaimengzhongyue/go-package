package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
)

type GK interface {
	GetKey() string
}

type Node struct {
	Key   string
	Value string
}

func (this Node) GetKey() string {
	return this.Key
}

func Encode() {
	n := &Node{Key: "key", Value: "value"}

	var b bytes.Buffer
	e := gob.NewEncoder(&b)
	err := e.Encode(n)
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile("./test", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	_, err = f.Write(b.Bytes())
	if err != nil {
		panic(err)
	}
}

func Decode() {
	bts, err := ioutil.ReadFile("./test")
	if err != nil {
		panic(err)
	}

	b := bytes.NewBuffer(bts)
	d := gob.NewDecoder(b)
	var i Node
	err = d.Decode(&i)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", i)
}

func main() {
	Encode()
	Decode()
}
