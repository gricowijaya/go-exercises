package main

import (
    "fmt"
)

// deadlock program channel is just for one buffer
func main() { 
    // create for 1 buffer
    c := make(chan int, 1)

    c <- 43
    
    fmt.Println(<-c)


    // create the anonymous func
    go func() {
        c <- 42 // assign 42 into c 
    }()

    fmt.Println(<-c) // output: 42
}
