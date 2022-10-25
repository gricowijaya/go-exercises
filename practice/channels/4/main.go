package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

// create the function that receive two channels which of type receive channels
func receive (c, q <-chan int) { 
    for {
        select {
        case  v := <-c:
            fmt.Println(v)
        case  <-q:
            return ;
        }
    }
}

// receive only channel 
func gen(q <-chan int) <-chan int {
	c := make(chan int)

    go func() {
        for i := 0; i < 100; i++ {
            c <- i
        }
    }()

	return c
}
