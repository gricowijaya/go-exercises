package main

import (
	"fmt"
)

func main() {
    // create the bidirectional channel
	cr := make(chan int)

	go func() {
		cr <- 42 // sending the chanel
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}
