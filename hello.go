package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go is running on: ")
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Linux")
	case "darwin":
		fmt.Println("OS X")
	default:
		fmt.Println(os)

	}
}
