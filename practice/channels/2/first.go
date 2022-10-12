package main

import (
	"fmt"
)

func main() {
	cr := make(chan int)

    // the process order with using anonymous go routine 
	go func() {
		cr <- 42
	}()
    
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)

}
