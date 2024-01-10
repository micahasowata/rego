package main

import "fmt"

func main() {
	a := make([]int, 5)
	fmt.Println(a, len(a), cap(a))

	b := make([]int, 0, 5)
	fmt.Println(b, len(b), cap(b))
}
