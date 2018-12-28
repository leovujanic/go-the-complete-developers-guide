package main

import (
	"fmt"
	"math"
	"strings"
)

const DEBUG = false

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {

	if x <= 0 {
		return 0, ErrNegativeSqrt(float64(x))
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

	return z, nil
}

func main() {

	separator := strings.Repeat("-", 20)

	for i := -3; i < 10; i += 1 {
		if sqrt, err := Sqrt(float64(i)); err == nil {
			fmt.Println("Main:", sqrt)
		} else {
			fmt.Println(err)
		}
		fmt.Println("Math:", math.Sqrt(float64(i)))
		fmt.Println(separator)
	}
}
