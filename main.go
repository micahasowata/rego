package main

import "github.com/spobly/rego/organiser"

func main() {
	o, _ := organiser.New("Documents/test", true)
	o.Run()
}
