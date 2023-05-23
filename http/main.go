package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// sb := make([]byte, 999999)
	// resp.Body.Read(sb)
	// fmt.Println(string(sb))

	lw := logWriter{}

	//	io.Copy(os.Stdout, resp.Body)
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(sb []byte) (int, error) {
	fmt.Println(string(sb))
	fmt.Println("Bytes: ", len(sb))
	return len(sb), nil
}
