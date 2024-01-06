package main

import (
	"fmt"
	"math"
)

func main() {
	x, y := 3, 4
	f := math.Sqrt(float64(x*x + y*y))
	z := uint(f)

	fmt.Println(x, y, z)
}
