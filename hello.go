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
	p := &v
	p.X = 45
	fmt.Println(v)
}
