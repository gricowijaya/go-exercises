package main

import (
	"fmt"
)

func main() {
	n, err := fmt.Println("Hello")

	// nil == nothing
	// which there's an error
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n)
}
