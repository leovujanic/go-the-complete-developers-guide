package main

import (
	"fmt"
	"math"
)

const DEBUG = false

func Sqrt(x float64) float64 {

	if x <= 0 {
		return 0
	}

	z := 1.0
	delta := z
	const precision float64 = 0.00000000001

	for delta >= precision {

		if DEBUG {
			fmt.Println("Z:", z)
		}

		oz := z
		z -= (z*z - x) / (2 * z)

		delta = oz - z
		if delta < 0 {
			delta = delta * -1
		}
	}

	return z
}

func main() {

	for i := 0; i < 10; i += 1 {
		fmt.Println(Sqrt(float64(i)))
		fmt.Println(math.Sqrt(float64(i)))
	}
}
