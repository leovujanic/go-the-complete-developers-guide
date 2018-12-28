package main

import "fmt"

type MyReader struct{}

func (r MyReader) Read(b []byte) (n int, err error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	fmt.Println("Nothing here")
}
