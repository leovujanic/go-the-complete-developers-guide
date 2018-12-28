package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
)

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64) // f = 0; ok = false

	fmt.Println(f, ok)
	// If i does not hold a T, the statement will trigger a panic.
	// f = i.(float64) // panic
	// fmt.Println(f)

	fmt.Println("got int", expectInt(2))
	fmt.Println("got int", expectInt("test"))
}

func expectInt(i interface{}) int {
	val, ok := i.(int)

	if ok == false {
		log.Fatal("unprocessable type")
	}

	return val
}
