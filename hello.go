package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	vOne   = Vertex{1, 2}
	vTwo   = Vertex{X: 1}
	vThree = Vertex{}
	p      = &Vertex{1, 2}
)

func main() {
	fmt.Println(vOne, p, vTwo, vThree)
}
