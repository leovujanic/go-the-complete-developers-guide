package main

import "fmt"

func main() {

	i := 1

	for i < 3 {
		fmt.Println("I is:", i)
		i += 1
	}

	for j := 5; j < 10; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("While loop")
		break
	}

	k := 0

	for k < 10 {
		fmt.Println("In while")
		k += 1
	}

}
