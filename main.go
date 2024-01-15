package main

import "github.com/spobly/rego/organiser"

func main() {
	o := organiser.New("Downloads", true)
	o.Run()
}
