package main

import "fmt"

func main() {
	q := []string{"Charles", "Olsen", "Rose", "Sutterland", "Bo", "Tori", "Gabriella"}
	fmt.Println(q)

	q = q[1:4]
	fmt.Println(q)

	q = q[:2]
	fmt.Println(q)

	q = q[1:]
	fmt.Println(q)
}
