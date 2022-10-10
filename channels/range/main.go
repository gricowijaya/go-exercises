package main

import("fmt")

// create the control flow channels which will take the go off.
func main() { 

    c := make(chan int)

    // send 
    go func() { 
        for i:=0; i < 1000; i++ { 
            c <- i  
        }
        close(c) // close that channels
    }()
    // will loop through the channels
    for v := range c { 
        fmt.Println(v)
    }
    
    fmt.Println("About to exit")
}

