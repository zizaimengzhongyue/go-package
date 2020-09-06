package main

import (
	"bufio"
	"fmt"
	"os"
)

func Scan() {
	f, err := os.Open("./scan.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	if scanner.Scan() == false {
		fmt.Println(scanner.Err())
	}
	for scanner.Scan() {
		str := scanner.Text()
		fmt.Printf("Scanner.String 返回值: %s\n", str)
	}
}

func main() {
	Scan()
}
