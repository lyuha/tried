package main

import (
	"fmt"
	"math"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNegativeSqrt(f)
	}

	const machine_epsilon = 1.11e-15
	z := float64(1)
	next_z := z - ((z*z)-f)/(2*z)
	for math.Abs(z-next_z) >= machine_epsilon {
		z = next_z
		next_z = z - ((z*z)-f)/(2*z)
	}

	return z, nil
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
