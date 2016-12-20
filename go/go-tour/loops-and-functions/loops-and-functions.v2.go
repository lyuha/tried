package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	const machine_epsilon = 1.11e-15
	z := float64(1)
	next_z := z - ((z*z)-x)/(2*z)
	for math.Abs(z-next_z) >= machine_epsilon {
		z = next_z
		next_z = z - ((z*z)-x)/(2*z)
	}
	return z
}

func main() {
	fmt.Println(math.Sqrt(2))
	fmt.Println(Sqrt(2))
}
