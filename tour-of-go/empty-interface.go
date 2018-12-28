package main

import (
	"fmt"
	"reflect"
)

func main() {
	// empty interface and type detection

	i := 1
	f := 3.14
	s := "hello"

	describe(i) // (1, int)
	describe(f) // (3.14, float64)
	describe(s) // (hello, string)

	fmt.Println(getType(i))

	typeOf(i)
	typeOf(f)
	typeOf(s)

	myType := reflect.TypeOf(s)

	fmt.Println(reflect.TypeOf(myType).String()) // *reflect.rtype
	fmt.Println(reflect.TypeOf(s).String())      // string

}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func getType(i interface{}) string {
	return reflect.TypeOf(i).String()
}

func typeOf(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("It's a int")
	case float64:
		fmt.Println("It's a float64")
	case string:
		fmt.Println("It's a string")
	default:
		fmt.Println("Unknown type")
	}
}
