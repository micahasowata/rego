package main

import "github.com/spobly/rego/organiser"

func main() {
	o := organiser.New("Documents/test", true)
	o.Run()
}
