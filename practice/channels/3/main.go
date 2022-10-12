package main

import(
    "fmt"
)

func main() {
    c := gen()
    receive(c)

    fmt.Println("About to exit")
}

// receivin channel which will receive some value
func receive(c <-chan int ) { 
    // which will close this exit
    for v := range c {
        fmt.Println(v)
    }
}

// received channel here
func gen() <-chan int {
    c := make(chan int)
    
    // create the func literal 
    go func() { 
        for i := 0 ; i< 100; i++ { 
            c<- i
        }
        close(c) // close the c channel
    }()
    for i := 0; i < 100; i++ {
        c <- i
    }

    return c
}
