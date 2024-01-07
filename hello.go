package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 2.55
	z -= (z*z - x) / (2 * z)

	return z
}

func main() {
	i := 0.0

	for i < 10.0 {
		fmt.Printf("Square root of %.2f is => %.2f", i, Sqrt(i))
		i += 1.0
	}
}
