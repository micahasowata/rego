package main

import "fmt"

func main() {
	var msg [2]string
	msg[0] = "Hello,"
	msg[1] = "World!"

	fmt.Println(msg[0], msg[1])
	fmt.Println(msg)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
