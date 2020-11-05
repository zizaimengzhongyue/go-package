package main

type Internal struct {
}

func (i Internal) Ping() string {
    return "PONG"
}

var I = Internal{}
