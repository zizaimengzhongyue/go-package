package main

import (
	"expvar"
	"fmt"
	"net/http"
)

var visits = expvar.NewInt("visits")
var dict = expvar.NewMap("dict")

func handler(w http.ResponseWriter, r *http.Request) {
	visits.Add(1)
	fmt.Println(visits.Value())
	fmt.Println(visits.String())
	fmt.Fprintf(w, "Handler: Hello,world")
}

func handler1(w http.ResponseWriter, r *http.Request) {
	dict.Add("handler1", 1)
	fmt.Println(dict.String())
	fmt.Fprintf(w, "Handler1: hello,world")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	expvar.Do(func(kv expvar.KeyValue) {
		fmt.Printf("%+v\n", kv)
	})
	fmt.Fprintf(w, "Handler2: hello,world")
}

func Expvar() {
	http.HandleFunc("/handler", handler)
	http.HandleFunc("/handler1", handler1)
	http.HandleFunc("/handler2", handler2)
	http.ListenAndServe(":8080", nil)
}

func main() {
	Expvar()
}

func init() {
	dict.Set("handler1", &expvar.Int{})
}
