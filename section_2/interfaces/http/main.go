package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type LogWriter struct{}

func main() {
	res, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// body := make([]byte, 99999)
	// res.Body.Read(body)
	// fmt.Println(string(body))

	io.Copy(LogWriter{}, res.Body)
}

func (LogWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
