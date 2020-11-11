package main

import (
	"fmt"
	"net/url"
)

func URL() {
	u := "a=b//c=d"

	e := url.PathEscape(u)
	fmt.Println(e)

	du, err := url.PathUnescape(e)
	if err != nil {
		panic(err)
	}
	fmt.Println(url.PathUnescape(du))

	q := "你好 天气 hello"

	eq := url.QueryEscape(q)
	fmt.Println(eq)

	dq, err := url.QueryUnescape(eq)
	if err != nil {
		panic(err)
	}
	fmt.Println(dq)

	r, err := url.Parse("http://www.baidu.com?hello=world")
	if err != nil {
		panic(err)
	}
	fmt.Println(r.EscapedPath())
	fmt.Println(r.Hostname())
	fmt.Println(r.IsAbs())
	bts, err := r.MarshalBinary()
	if err != nil {
		panic(err)
	}
	fmt.Println(bts)
	err = r.UnmarshalBinary(bts)
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Port())
	fmt.Println(r.Query())
	fmt.Println(r.RequestURI())
	fmt.Println(r.String())
}

func Userinfo() {
	u := url.UserPassword("张三", "123456")
	p, b := u.Password()
	fmt.Println(p, b)
	fmt.Println(u.String())
	fmt.Println(u.Username())
}

func Values() {
	v, err := url.ParseQuery("hello=world")
	if err != nil {
		panic(err)
	}
	fmt.Println(v)
	fmt.Println(v.Get("hello"))
	v.Set("hello", "world01")
	v.Add("hello01", "world02")
	fmt.Println(v.Encode())
}

func main() {
	URL()
	Userinfo()
	Values()
}
