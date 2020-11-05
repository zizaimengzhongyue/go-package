package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func Filepath() {
	p, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	s, err := filepath.Abs(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)

	fmt.Println(filepath.Clean(p))
	fmt.Println(filepath.Dir(p))

	s, err = filepath.EvalSymlinks(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)

	fmt.Println(filepath.Ext(p))
	fmt.Println(filepath.FromSlash(p))

	ms, err := filepath.Glob(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(ms)

	fmt.Println(filepath.HasPrefix(p, "/home"))
	fmt.Println(filepath.IsAbs(p))
	fmt.Println(filepath.Join(p, p))

	dir, f := filepath.Split(p)
	fmt.Println(dir, f)

	fmt.Println(filepath.ToSlash(p))

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		fmt.Printf("%+v\n", info)
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func main() {
	Filepath()
}
