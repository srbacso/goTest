package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {
}

func main() {
	fmt.Println("Starting")

	resp, err := http.Get("http://google.com/")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	//bs := make([]byte, 1024*1024)
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))

	lw := logWriter{}

	io.Copy(lw, resp.Body)

	fmt.Println("Ending")
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Printf("just wrote: %v bytes", len(bs))
	return len(bs), nil
}
