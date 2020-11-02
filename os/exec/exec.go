package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
)

func Lsla() {
	cmd := exec.Command("ls", "-la")
	fmt.Println(cmd.String())

	bts, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bts))
}

func Date() {
	cmd := exec.Command("date")
	fmt.Println(cmd.String())

	out, err := cmd.StdoutPipe()
	defer out.Close()
	if err != nil {
		panic(err)
	}

	err = cmd.Start()
	if err != nil {
		panic(err)
	}

	bts, err := ioutil.ReadAll(out)
	if err != nil {
		panic(err)
	}

	err = cmd.Wait()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bts))
}

func Echo() {
	cmd := exec.Command("echo", "hello,world")
	fmt.Println(cmd.String())

	buf := bytes.Buffer{}
	cmd.Stdout = &buf

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf.Bytes()))
}

func Multi() {
	cmds := `ps -ef | sort -rnk 3 | head -10`
	cmd := exec.Command("bash", "-c", cmds)
	fmt.Println(cmd.String())

	bts, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bts))
}

func main() {
	Lsla()
	Date()
	Echo()
	Multi()
}
