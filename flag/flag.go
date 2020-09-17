package main

import (
	"flag"
	"fmt"
)

var (
	Arg     int
	Args    string
	BoolVar bool
)

func Flag() {
	flag.Parse()
	fmt.Printf("flag.Arg 返回指定的命令行参数: %s\n", flag.Arg(0))
	fmt.Printf("flag.Args 返回没有被 flag 捕获的命令行参数: %#v\n", flag.Args())
	fmt.Printf("flag.NArg 返回没有被 flag 不说的命令行参数个数: %d\n", flag.NArg())
	fmt.Printf("flag.NFlag 返回设置过的参数个数: %d\n", flag.NFlag())
	fmt.Printf("flag.BoolVar: %t\n", BoolVar)
	err := flag.Set("args", "hello,world")
	if err != nil {
		panic(err)
	}
	fmt.Printf("flag.Set 可以设置 flag 注册变量的值: %s\n", Args)
	fmt.Printf("flag.Parsed 返回一个 bool 值表示 flag 是否被 porse 过: %t\n", flag.Parsed())
	flag.Visit(func(f *flag.Flag) {
		fmt.Println(f.Name)
	})
	flag.Visit(func(f *flag.Flag) {
		fmt.Println(f.Name)
	})
}

func main() {
	Flag()
}

func init() {
	flag.IntVar(&Arg, "arg", 0, "Arg 方法第 i 个命令行参数")
	flag.StringVar(&Args, "args", "", "Args 方法返回没有被 flag 捕获的命令行参数")
	flag.BoolVar(&BoolVar, "boolVar", false, "BoolVar 方法注册一个 bool 类型的 flag 变量")
}
