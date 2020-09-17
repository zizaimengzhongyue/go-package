package main

import (
	"bufio"
	"image"
	"image/png"
	"os"
)

func Image() {
	bg := image.NewRGBA(image.Rect(0, 0, 500, 500))
	f, err := os.Create("test.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := bufio.NewWriter(f)
	err = png.Encode(buf, bg)
	if err != nil {
		panic(err)
	}
	buf.Flush()
}

func main() {
	Image()
}
