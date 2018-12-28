package main

import "fmt"

func main() {
	x, y := split(33)

	fmt.Printf("%d %d\n", x, y)

	a, b := something(true)
	fmt.Printf("%s %s\n", a, b)

	a, b = something(false)
	fmt.Printf("%s %s\n", a, b)
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func something(exitFirst bool) (x, y string) {
	y = "Y"
	x = "X"

	if exitFirst {
		return
	}

	x = "X2"
	y = "Y2"

	return
}
