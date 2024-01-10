package main

import "fmt"

func printSlice(s []string) {
	fmt.Printf("len=%d, cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	q := []string{"Charles", "Olsen", "Rose", "Sutterland", "Bo", "Tori", "Gabriella"}
	printSlice(q)

	q = q[1:4]
	printSlice(q)

	q = q[:2]
	printSlice(q)

	q = q[1:]
	printSlice(q)
}
