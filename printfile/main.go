package main

import (
	"io"
	"os"
)

func main() {
	buff, e := os.Open(os.Args[1])

	if e != nil {
		panic(e)
	}

	io.Copy(os.Stdout, buff)
	buff.Close()
}
