package main

import "fmt"

func main() {
	var names []string
	names = append(names, "Adam", "Job", "Queen")

	for i, name := range names {
		fmt.Printf("%d => %v\n", i, name)
	}
}
