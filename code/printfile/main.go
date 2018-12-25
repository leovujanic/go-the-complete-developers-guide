package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Please provide file path")
		os.Exit(1)
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	_, _ = io.Copy(os.Stdout, f)
}
