package main

import (
	"fmt"
	"net"
)

func JoinHostPort() {
	host := "127.0.0.1"
	port := "8088"
	fmt.Println(net.JoinHostPort(host, port))
}

func LookupAddr() {
	str := "127.0.0.1"
	names, err := net.LookupAddr(str)
	if err != nil {
		panic(err)
	}
	fmt.Println(names)
}

func LookupCname() {
	str := "www.baidu.com"
	cname, err := net.LookupCNAME(str)
	if err != nil {
		panic(err)
	}
	fmt.Println(cname)
}

func LookupHost() {
	host := "www.baidu.com"
	addrs, err := net.LookupHost(host)
	if err != nil {
		panic(err)
	}
	fmt.Println(addrs)
}

func SplitHostPort() {
	str := "127.0.0.1:8080"
	host, port, err := net.SplitHostPort(str)
	if err != nil {
		panic(err)
	}
	fmt.Println(host, port)
}

func main() {
	JoinHostPort()
	LookupAddr()
	LookupCname()
	LookupHost()
	SplitHostPort()
}
