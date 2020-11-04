package main

import (
	"fmt"
	"os"
	"os/user"
)

func User() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", u)

	ids, err := u.GroupIds()
	if err != nil {
		panic(err)
	}
	fmt.Println(ids)

	g, err := user.LookupGroup(u.Gid)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	fmt.Printf("%+v\n", g)
}

func main() {
	User()
}
