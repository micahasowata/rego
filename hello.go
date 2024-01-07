package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2701

	p := &i

	fmt.Println(*p)

	*p = (j - 1) / 15

	fmt.Println(*p)
}
