package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func HTTP() {
	r, err := http.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	bts, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bts))

	fmt.Println(http.DetectContentType(bts))

	fmt.Println(http.CanonicalHeaderKey("content-type"))
}

func Client() {
	c := &http.Client{}
	req, err := http.NewRequest("Get", "https://www.baidu.com", nil)
	if err != nil {
		panic(err)
	}
	res, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	bts, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bts))
}

func Serve() {
	m := http.NewServeMux()
	m.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong")
	})
	m.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "world")
	})
	http.ListenAndServe(":8080", m)
}

func main() {
	HTTP()
	Client()
	Serve()
}
