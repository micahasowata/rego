package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 0.0
	for z < x {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(4))
}
