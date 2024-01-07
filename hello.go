package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{5, 7}
	fmt.Println(v.X, v.Y)
}
