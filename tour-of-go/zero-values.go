package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	fmt.Printf("%T %v\n", i, i) // int 0
	fmt.Printf("%T %v\n", f, f) // float64 0
}
