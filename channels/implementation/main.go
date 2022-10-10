package main

import("fmt")

// create the control flow channels which will take the go off.
func main() { 

    c := make(chan int)

    go foo(c)
    go bar(c) 
    
    fmt.Println("About to exit")
}

// send  only an int
func foo(c chan<- int) { 
    // we're only going to be abel to put a value 
    c <- 42
}

// receive only an int
func bar(c <- chan int) { 
    fmt.Println(<-c)
}
