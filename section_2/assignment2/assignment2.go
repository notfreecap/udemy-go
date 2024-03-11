package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := os.Args[1]

	fmt.Println("filename", fileName)

	if fileName == "" {
		fmt.Println("Error: Please privide a valid file name, received:", fileName)
		os.Exit(1)
	}

	file, error := os.Open(fileName)
	if error != nil {
		fmt.Println("Error:", error)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)

}
