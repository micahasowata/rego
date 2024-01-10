package main

import "fmt"

func main() {
	var names []string
	names = append(names, "Adam", "Job", "Queen")

	for _, name := range names {
		fmt.Println(name)
	}
}
