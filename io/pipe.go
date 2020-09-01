package main

import (
	"fmt"
	"io"
)

var pipeStruct = `type pipe struct {
    wrMu sync.Mutex // Serializes Write operations
    wrCh chan []byte
    rdCh chan int

    once sync.Once // Protects closing done
    done chan struct{}
    rerr onceError
    werr onceError
}`

func IPC() {
	fmt.Println(pipeStruct)
	fmt.Printf("io.pipe 可以用于进程间通信\n")
	fmt.Printf("io.Pipe() 方法返回一个 io.PipeReader 和一个 io.PipeWriter，两者共享同一个 *pipe\n")
	fmt.Printf("在 *pipe 结构中包含了 wrCh、rdCh 两个 channel；PipeReader.Close 关闭 rdCh，PipeWriter.Close 关闭 wrch\n")
	fmt.Printf("pipe.done 主要用于关闭 pipe.done，```p.once.Do(func() { close(p.done) })```\n")
	reader, writer := io.Pipe()
	defer func() {
		reader.Close()
		writer.Close()
	}()
	go func() {
		writer.Write([]byte("hello,world"))
	}()

	buffer := make([]byte, 100)
	reader.Read(buffer)
	fmt.Println(string(buffer))
}

func main() {
	IPC()
}
