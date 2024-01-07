package main

import (
	"fmt"
)

func main() {
	fmt.Println("counting")
	defer fmt.Println("done")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
