package main

import "fmt"

func main() {
	o := NewOrganiser(".")
	fmt.Println(o.Path)
	o.Run()
}
