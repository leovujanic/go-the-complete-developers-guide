package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// NOTE: statement can precede conditionals
	if num := getRandom(-100, 100); num < 0 {
		fmt.Println("Num is negative")
	} else if num == 0 {
		fmt.Println("Num is zero")
	} else {
		fmt.Println("num is positive")
	}

}

func getRandom(min, max int) int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	return r.Intn(max-min) + min
}
