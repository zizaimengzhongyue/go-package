package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

func OS() {
	p, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	err = os.Chdir("/Users")
	if err != nil {
		panic(err)
	}

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(path)

	err = os.Chdir(p)
	if err != nil {
		panic(err)
	}

	err = os.Chmod("test", 0777)
	if err != nil {
		panic(err)
	}

	path, err = os.Executable()
	if err != nil {
		panic(err)
	}
	fmt.Println(path)

	err = os.Setenv("hello", "world")
	if err != nil {
		panic(err)
	}
	fmt.Println(os.Getenv("hello"))
	err = os.Unsetenv("hello")
	if err != nil {
		panic(err)
	}
	fmt.Println(os.Getenv("hello"))

	fmt.Println(os.Environ())

	fmt.Println(os.Expand("$hello", func(s string) string { return "world" }))

	fmt.Println(os.ExpandEnv("this is path: $PATH"))

	fmt.Println(os.Getegid())

	fmt.Println(os.Geteuid())

	fmt.Println(os.Getgid())

	fmt.Println(os.Getpagesize())

	fmt.Println(os.Getpid())

	fmt.Println(os.Getppid())

	fmt.Println(os.Getuid())

	host, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Println(host)

	env, _ := os.LookupEnv("PATH")
	fmt.Println(env)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	go func() {
		defer wg.Done()
		bts := make([]byte, 1024)
		_, _ = r.Read(bts)
		fmt.Println(string(bts))
	}()
	w.Write([]byte("hello, world"))
	wg.Wait()
}

func File() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	f, err := os.OpenFile(path+"/test", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Println(f.Fd())
	fmt.Println(f.Name())

	fi, err := f.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Println(fi.Name(), fi.Size(), fi.Mode(), fi.ModTime(), fi.IsDir(), fi.Sys())

	_, err = f.WriteString("hello,world")
	if err != nil {
		panic(err)
	}
}

func Process() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	bls := ""
	if filepath.Base("date") == "date" {
		bls, err = exec.LookPath("date")
		if err != nil {
			panic(err)
		}
	}

	ls, err := os.StartProcess(bls, []string{path}, &os.ProcAttr{Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", ls)
	ps, err := ls.Wait()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", ps)
}

func Signal() {
	fmt.Println(os.Interrupt.String())
	fmt.Println(os.Kill.String())
}

func main() {
	OS()
	File()
	Process()
	Signal()
}
